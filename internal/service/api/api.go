package api

import (
	"errors"
	"math/rand"
	"strings"
	"test-ozon/internal/storage"
)

// type ApiService struct {
// 	Storage *storage.DB
// }

// func NewApiService(storage *storage.DB) *ApiService {
// 	return &ApiService{
// 		Storage: storage,
// 	}
// }

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

func (s *ServiceApi) SaveUrl(urlToSave, alias string) error {
	if alias == "" {
		alias = RandAlias()

		ok := s.Storage.IsDublicate(alias)
		if ok {
			return errors.New("try to explicitly specify the alias")
		}
	} else {
		ok := s.Storage.IsDublicate(alias)
		if ok {
			return errors.New("alias is already exist")
		}
	}

	return s.Storage.SaveUrl(urlToSave, alias)
}

func RandAlias() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	s := make([]string, 10)
	for i := range s {
		s[i] = string(charset[rand.Intn(len(charset))])
	}

	res := strings.Join(s, "")
	return res
}
