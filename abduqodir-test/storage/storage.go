package storage

import (
	"app/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
}

type BookRepoI interface {
	Create(*models.CreateBook) (string, error)
	GetByID(*models.BookPrimaryKey) (*models.Book, error)
	Update(*models.UpdateBook) (*models.Book, error)
	GetList(*models.GetListBookRequest) (*models.GetListBookResponse, error)
}