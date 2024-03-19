package data

import (
	"context"
	"strconv"

	"gorm.io/gorm"
	"kratos-admin/api"
	v1 "kratos-admin/api/product/v1"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/data/po"
)

type ProductCategoryRepo struct {
	data *Data
}

func NewProductCategoryRepo(data *Data) biz.ProductCategoryRepo {
	return &ProductCategoryRepo{data: data}
}

func (p *ProductCategoryRepo) FindProductCategoryList(ctx context.Context, param *v1.FindProductCategoryListRequest) (reply *v1.FindProductCategoryListReply, err error) {
	var (
		count int64
		list  []*po.PmsProductCategory
	)
	tx := p.data.DB.WithContext(ctx).
		Model(&po.PmsProductCategory{}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if param.Name != "" {
				db = db.Where("name like ?", "%"+param.Name+"%")
			}
			if param.Level != 0 {
				db = db.Where("level = ?", param.Level)
			}
			if param.ParentId != 0 {
				db = db.Where("parent_id = ?", param.ParentId)
			}
			if param.ShowStatus != v1.DisplayStatus_DISPLAY_STATUS_UNKNOWN {
				db = db.Where("show_status = ?", param.ShowStatus)
			}
			if param.NavStatus != v1.DisplayStatus_DISPLAY_STATUS_UNKNOWN {
				db = db.Where("nav_status = ?", param.NavStatus)
			}
			return db
		})
	err = tx.Count(&count).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}
	err = tx.Offset(int((param.GetPageNumber() - 1) * param.PageSize)).Limit(int(param.PageSize)).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}

	reply.ProductCategories = make([]*v1.ProductCategory, 0, len(list))
	for _, v := range list {
		reply.ProductCategories = append(reply.ProductCategories, p.productCategoryPoToDto(v))
	}
	reply.Total = int32(count)
	return
}

func (p *ProductCategoryRepo) productCategoryPoToDto(v *po.PmsProductCategory) *v1.ProductCategory {
	return &v1.ProductCategory{
		Id:           strconv.FormatInt(v.Id, 10),
		Name:         v.Name,
		Level:        v.Level,
		ParentId:     strconv.FormatInt(v.ParentId, 10),
		Sort:         v.Sort,
		Icon:         v.Icon,
		Description:  v.Description,
		Keywords:     v.Keywords,
		ProductCount: v.ProductCount,
		ProductUnit:  v.ProductUnit,
		NavStatus:    v1.DisplayStatus(v.NavStatus),
		ShowStatus:   v1.DisplayStatus(v.ShowStatus),
	}
}

func (p *ProductCategoryRepo) productCategoryDtoToPo(category *v1.ProductCategory) *po.PmsProductCategory {
	parentId, _ := strconv.ParseInt(category.ParentId, 10, 64)
	var id int64
	if category.Id == "" {
		id = 0
	} else {
		id, _ = strconv.ParseInt(category.Id, 10, 64)
	}
	return &po.PmsProductCategory{
		Id:           id,
		Name:         category.Name,
		Level:        category.Level,
		ParentId:     parentId,
		Sort:         category.Sort,
		Icon:         category.Icon,
		Description:  category.Description,
		Keywords:     category.Keywords,
		ProductCount: category.ProductCount,
		ProductUnit:  category.ProductUnit,
		NavStatus:    int32(category.NavStatus),
		ShowStatus:   int32(category.ShowStatus),
	}
}

func (p *ProductCategoryRepo) CreateProductCategory(ctx context.Context, param *v1.CreateProductCategoryRequest) (int64, error) {
	product := p.productCategoryDtoToPo(param.ProductCategory)
	if err := p.data.DB.WithContext(ctx).Create(product).Error; err != nil {
		return 0, api.ErrorDbError("Failed to create product category").WithCause(err)
	}
	return product.Id, nil
}

func (p *ProductCategoryRepo) UpdateProductCategory(ctx context.Context, param *v1.UpdateProductCategoryRequest) error {
	product := p.productCategoryDtoToPo(param.ProductCategory)
	err := p.data.DB.WithContext(ctx).Model(&po.PmsProductCategory{}).Updates(product).Error
	if err != nil {
		return api.ErrorDbError("Failed to update product category %d", param.ProductCategory.Id).WithCause(err)
	}
	return nil
}

func (p *ProductCategoryRepo) DeleteProductCategory(ctx context.Context, id int64) error {
	err := p.data.DB.WithContext(ctx).Delete(&po.PmsProductCategory{}, id).Error
	if err != nil {
		return api.ErrorDbError("Failed to delete product category %d", id).WithCause(err)
	}
	return nil
}

func (p *ProductCategoryRepo) FindProductCategoryListByParentId(ctx context.Context, parentId int64) ([]*v1.ProductCategory, error) {
	var list []*po.PmsProductCategory
	err := p.data.DB.WithContext(ctx).Where("parent_id = ?", parentId).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}
	categories := make([]*v1.ProductCategory, len(list))
	for i, v := range list {
		categories[i] = p.productCategoryPoToDto(v)
	}
	return categories, nil
}

func (p *ProductCategoryRepo) FindProductCategoryTree(ctx context.Context) ([]*v1.ProductCategoryTreeNode, error) {
	var list []*po.PmsProductCategory
	err := p.data.DB.WithContext(ctx).Model(&po.PmsProductCategory{}).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}

	data := make(map[int64]*v1.ProductCategoryTreeNode)
	for _, v := range list {
		data[v.Id] = &v1.ProductCategoryTreeNode{
			ProductCategory: p.productCategoryPoToDto(v),
			Children:        nil,
		}
	}
	nodes := p.listToTree(data, 0)
	return nodes, nil
}

func (p *ProductCategoryRepo) listToTree(data map[int64]*v1.ProductCategoryTreeNode, parentId int64) []*v1.ProductCategoryTreeNode {
	nodes := make([]*v1.ProductCategoryTreeNode, 0)
	for id, node := range data {
		if node.ProductCategory.ParentId == strconv.FormatInt(parentId, 10) {
			node.Children = p.listToTree(data, id)
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (p *ProductCategoryRepo) FindProductCategoryById(ctx context.Context, id int64) (*v1.ProductCategory, error) {
	var productCategory po.PmsProductCategory
	if err := p.data.DB.WithContext(ctx).First(&productCategory, id).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product category %d", id).WithCause(err)
	}
	return p.productCategoryPoToDto(&productCategory), nil
}
