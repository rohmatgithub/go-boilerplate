package usecase

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/repository"
)

type ExampleUseCase interface {
	GetExampleByID(companyID, exampleCode string) (dto.Payload, common.ErrorModel)
}

type exampleUsecase struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleUseCase(exampleRepository repository.ExampleRepository) ExampleUseCase {
	return &exampleUsecase{
		exampleRepository: exampleRepository,
	}
}

func (u *exampleUsecase) GetExampleByID(companyID, exampleCode string) (result dto.Payload, errMdl common.ErrorModel) {
	ex, err := u.exampleRepository.GetExampleByID(companyID, exampleCode)
	if err != nil {
		errMdl = common.GenerateInternalDBServerError(err)
		return
	}

	result = dto.Payload{
		Data: dto.ExampleDetail{
			ID:          ex.ID,
			ExampleCode: ex.ExampleCode,
		},
	}
	return
}
