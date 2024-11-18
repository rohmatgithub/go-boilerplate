package usecase

import (
	"boilerplate/internal/common"
	"boilerplate/internal/constanta"
	"boilerplate/internal/domain"
	"boilerplate/internal/dto"
	"boilerplate/internal/repository"
	"database/sql"
	"fmt"
	"time"
)

type BookUseCase interface {
	GetBookByID(id int64) (dto.Payload, common.ErrorModel)
	InsertBook(book *dto.BookRequest, ctxModel *common.ContextModel) (result dto.Payload, errMdl common.ErrorModel)
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
	ex, err := u.bookRepository.GetDetailByID(id)
	if err != nil {
		errMdl = common.GenerateDatabaseError(err)
		return
	}

	if ex.ID == 0 {
		errMdl = common.GenerateDataNotFoundError()
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

func (u *bookUsecase) InsertBook(book *dto.BookRequest, ctxModel *common.ContextModel) (result dto.Payload, errMdl common.ErrorModel) {
	t, err := time.Parse("2006-01-02", book.PublishedDate)
	if err != nil {
		errMdl = common.GenerateInvalidFormatError(constanta.PublishedDate)
		return
	}
	err = u.bookRepository.Insert(&domain.Book{
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: sql.NullTime{Time: t, Valid: true},
		CategoryID:    book.CategoryID,
		Price:         book.Price,
		Stock:         book.Stock,
	})

	if err != nil {
		errMdl = common.GenerateDatabaseError(err)
		return
	}
	return
}
