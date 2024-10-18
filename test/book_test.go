package test

import (
	"boilerplate/internal/dto"
	"boilerplate/pkg/configs"
	"boilerplate/pkg/util"
	"fmt"
	"log"
	"testing"
)

func TestBookDto(t *testing.T) {
	err := configs.Init()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	util.InitializeI18n()
	validate := util.NewAppValidator()

	b := dto.BookRequest{
		Title:         "The Lord of the Rings",
		Author:        "J. R. R. Tolkien",
		PublishedDate: "1954-07-29",
		CategoryID:    1,
		Price:         12.99,
		Stock:         2,
	}

	// b = dto.BookRequest{}

	maps, err := validate.ValidateRequest(b)
	fmt.Println(maps)
	if err != nil {
		t.Error(err)
	}

}
