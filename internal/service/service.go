package service

import "test-ozon/internal/storage"

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Api
type Api interface{
	GetUrl(alias string) (string, error)
	SaveUrl(urlToSave string) (string, error)
}

type Service struct {
	Api
}

func NewService(storage *storage.DB) *Service {
	return &Service{
		Api: NewServiceApi(storage),
	}
}