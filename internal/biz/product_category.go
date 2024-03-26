package biz

import (
	"context"

	v1 "github.com/wxlbd/kratos-pms/api/product/v1"
)

type ProductCategoryUseCase struct {
	productCategoryRepo ProductCategoryRepo
}

func NewProductCategoryUseCase(productCategoryRepo ProductCategoryRepo) *ProductCategoryUseCase {
	return &ProductCategoryUseCase{productCategoryRepo: productCategoryRepo}
}

func (u *ProductCategoryUseCase) FindProductCategoryList(ctx context.Context, param *v1.FindProductCategoryListRequest) (reply *v1.FindProductCategoryListReply, err error) {
	return u.productCategoryRepo.FindProductCategoryList(ctx, param)
}

func (u *ProductCategoryUseCase) CreateProductCategory(ctx context.Context, param *v1.CreateProductCategoryRequest) error {
	_, err := u.productCategoryRepo.CreateProductCategory(ctx, param)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductCategoryUseCase) UpdateProductCategory(ctx context.Context, param *v1.UpdateProductCategoryRequest) error {
	return u.productCategoryRepo.UpdateProductCategory(ctx, param)
}

func (u *ProductCategoryUseCase) DeleteProductCategory(ctx context.Context, id int64) error {
	return u.productCategoryRepo.DeleteProductCategory(ctx, id)
}

func (u *ProductCategoryUseCase) FindProductCategoryListByParentId(ctx context.Context, parentId int64) ([]*v1.ProductCategory, error) {
	return u.productCategoryRepo.FindProductCategoryListByParentId(ctx, parentId)
}

func (u *ProductCategoryUseCase) FindProductCategoryTree(ctx context.Context) ([]*v1.ProductCategoryTreeNode, error) {
	return u.productCategoryRepo.FindProductCategoryTree(ctx)
}

func (u *ProductCategoryUseCase) FindProductCategoryById(ctx context.Context, id int64) (*v1.ProductCategory, error) {
	return u.productCategoryRepo.FindProductCategoryById(ctx, id)
}
