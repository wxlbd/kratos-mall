package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewProductCategoryUseCase, NewProductUseCase, NewProductAttributeUseCase, NewProductAttributeValueUseCase, NewAuthUseCase)
