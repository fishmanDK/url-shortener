package service

import (
	"errors"
	"math/rand"
	"test-ozon/internal/storage"
)

const (
	sizeRand = 10
)

type ServiceApi struct {
	Storage *storage.DB
}

func NewServiceApi(storage *storage.DB) *ServiceApi {
	return &ServiceApi{
		Storage: storage,
	}
}

func (s *ServiceApi) GetUrl(alias string) (string, error) {
	return s.Storage.GetUrl(alias)
}

func (s *ServiceApi) SaveUrl(urlToSave string) (string, error) {
	alias := RandAlias(sizeRand)

	ok := s.Storage.IsDublicate(alias)
	if ok {
		return "", errors.New("try to explicitly specify the alias")
	}

	return alias, s.Storage.SaveUrl(urlToSave, alias)
}

func RandAlias(size int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	b := make([]byte, size)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
