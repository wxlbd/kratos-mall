package data

import (
	"context"
	"gorm.io/gorm"
	"kratos-admin/api"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/data/po"
)

type ProductCategoryRepo struct {
	data *Data
}

func (p *ProductCategoryRepo) FindProductCategoryList(ctx context.Context, param *biz.ListProductCategoryParam) (total int64, list []*biz.ProductCategory, err error) {
	err = p.data.DB.Model(&po.PmsProductCategory{}).Count(&total).Error
	if err != nil {
		return 0, nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}

	err = p.data.DB.Model(&po.PmsProductCategory{}).Offset((param.Page.Number - 1) * param.Page.Size).Scopes(func(db *gorm.DB) *gorm.DB {
		if param.Name != "" {
			db = db.Where("name like ?", "%"+param.Name+"%")
		}
		if param.Level != 0 {
			db = db.Where("level = ?", param.Level)
		}
		if param.ParentId != 0 {
			db = db.Where("parent_id = ?", param.ParentId)
		}
		if param.ShowStatus != biz.DisplayStatus_UNKNOWN {
			db = db.Where("show_status = ?", param.ShowStatus)
		}
		if param.NavStatus != biz.DisplayStatus_UNKNOWN {
			db = db.Where("nav_status = ?", param.NavStatus)
		}
		return db
	}).Limit(param.Page.Size).Find(&list).Error
	if err != nil {
		return 0, nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}
	return
}

func (p *ProductCategoryRepo) CreateProductCategory(ctx context.Context, param *biz.CreateProductCategoryParam) (int64, error) {
	p.data.DB.WithContext(ctx).Model(&po.PmsProductCategory{}).Updates(&param.ProductCategory)
	return param.ProductCategory.Id, nil
}

func (p *ProductCategoryRepo) UpdateProductCategory(ctx context.Context, param *biz.UpdateProductCategoryParam) error {
	err := p.data.DB.WithContext(ctx).Model(&po.PmsProductCategory{}).Updates(&param.ProductCategory).Error
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

func (p *ProductCategoryRepo) FindProductCategoryListByParentId(ctx context.Context, parentId int64) ([]*biz.ProductCategory, error) {
	var list []*biz.ProductCategory
	err := p.data.DB.WithContext(ctx).Where("parent_id = ?", parentId).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}
	return list, nil
}

func (p *ProductCategoryRepo) FindProductCategoryTree(ctx context.Context) ([]*biz.TreeNode, error) {
	var list []*po.PmsProductCategory
	err := p.data.DB.WithContext(ctx).Model(&po.PmsProductCategory{}).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product category list").WithCause(err)
	}

	data := make(map[int64]*biz.TreeNode)
	for _, v := range list {
		data[v.Id] = &biz.TreeNode{
			ProductCategory: p.productCategoryPoToDo(v),
		}
	}
	nodes := p.listToTree(data, 0)
	return nodes, nil
}

func (p *ProductCategoryRepo) listToTree(data map[int64]*biz.TreeNode, parentId int64) []*biz.TreeNode {
	nodes := make([]*biz.TreeNode, 0)
	for id, node := range data {
		if node.ParentId == parentId {
			node.Children = p.listToTree(data, id)
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (p *ProductCategoryRepo) FindProductCategoryById(ctx context.Context, id int64) (*biz.ProductCategory, error) {
	var productCategory po.PmsProductCategory
	if err := p.data.DB.First(&productCategory, id).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product category %d", id).WithCause(err)
	}
	return p.productCategoryPoToDo(&productCategory), nil
}

func (p *ProductCategoryRepo) productCategoryPoToDo(productCategory *po.PmsProductCategory) *biz.ProductCategory {
	return &biz.ProductCategory{
		Id:          productCategory.Id,
		Name:        productCategory.Name,
		Level:       productCategory.Level,
		ParentId:    productCategory.ParentId,
		Icon:        productCategory.Icon,
		Keywords:    productCategory.Keywords,
		Description: productCategory.Description,
		ShowStatus:  biz.DisplayStatus(productCategory.ShowStatus),
		Sort:        string(productCategory.Sort),
	}
}
