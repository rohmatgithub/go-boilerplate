package usecase

import (
	"boilerplate/internal/common"
	"boilerplate/internal/dto"
	"boilerplate/internal/repository"
	"fmt"
)

type BookUseCase interface {
	GetBookByID(id int64) (dto.Payload, common.ErrorModel)
}

type bookUsecase struct {
	bookRepository repository.BookRepository
}

func NewBookUseCase(bookRepository repository.BookRepository) BookUseCase {
	return &bookUsecase{
		bookRepository: bookRepository,
	}
}

func (u *bookUsecase) GetBookByID(id int64) (result dto.Payload, errMdl common.ErrorModel) {
	ex, err := u.bookRepository.GetExampleByID(id)
	if err != nil {
		errMdl = common.GenerateInternalDBServerError(err)
		return
	}

	result = dto.Payload{
		Data: dto.BookDetail{
			ID:            ex.ID,
			Title:         ex.Title,
			Author:        ex.Author,
			PublishedDate: ex.PublishedDate.Time.Format("2006-01-02"),
			CategoryName:  "Category",
			Price:         fmt.Sprintf("%.2f", ex.Price),
			Stock:         fmt.Sprintf("%d", ex.Stock),
		},
	}
	return
}
