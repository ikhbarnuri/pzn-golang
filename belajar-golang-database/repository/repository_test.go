package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "eko@mail.com",
		Comment: "Ini komentar",
	}

	insert, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 23)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentRepositoryImpl_FindByAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(comments)
}
