package storage

import (
	"test-ozon/internal/storage/config"
	"test-ozon/internal/storage/memory"
	"test-ozon/internal/storage/postgres"
)


type DB struct {
    dbInterface
}

type dbInterface interface {
	GetUrl(aliasUrl string) (string, error)
    SaveUrl(urlToSave, aliasUrl string) error
    IsDublicate(alias string) bool
}

func NewDB(name string) (*DB, error) {
    var (
		db dbInterface
		err error
	)

    switch name {
    case "memory":
		db = memory.NewMemory()
    case "postgres":
		cfg := config.NewConfig_Postgres()
        db, err = postgres.NewPostgres(cfg)
    default:
        cfg := config.NewConfig_Postgres()
        db, err = postgres.NewPostgres(cfg)
    }

    return &DB{
        dbInterface: db,
    }, err
}