package usecase

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/repository"
)

type ProductUseCase interface {
	GetProductByID(companyID, productCode string) (dto.Payload, common.ErrorModel)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) ProductUseCase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (u *productUsecase) GetProductByID(companyID, productCode string) (result dto.Payload, errMdl common.ErrorModel) {
	product, err := u.productRepository.GetProductByID(companyID, productCode)
	if err != nil {
		errMdl = common.GenerateInternalDBServerError(err)
		return
	}

	result = dto.Payload{
		Data: dto.ProductDetail{
			ProductCode: product.ProductCode,
			ProductName: product.ProductName.String,
		},
	}
	return
}
