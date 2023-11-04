package postgres

import (
	"fmt"
	"test-ozon/internal/storage/config"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	postgres *sqlx.DB
}

func NewPostgres(cfg *config.Config_Postgres) (*Postgres, error) {
	fmt.Println(cfg)
	link := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Name, cfg.Postgres.Password, cfg.Postgres.SSLMode)
	db, err := sqlx.Open("postgres", link)

	if err != nil {
		return &Postgres{}, err
	}

	err = db.Ping()
	if err != nil {
		return &Postgres{}, err
	}

	smtp, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS url" +
		"( " +
			"id SERIAL, " +
			"url VARCHAR(255) NOT NULL, " +
			"alias VARCHAR(10) NOT NULL UNIQUE" +
		")")

	if err != nil{
		return &Postgres{}, err
	}

	_, err = smtp.Exec()
	if err != nil{
		return &Postgres{}, err
	}

	return &Postgres{
		postgres: db,
	}, nil
}

func (p *Postgres) SaveUrl(urlToSave, aliasUrl string) error {
	smtp, err := p.postgres.Prepare("INSERT INTO url (url, alias) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	_, err = smtp.Exec(urlToSave, aliasUrl)
	return err
}

func (p *Postgres) GetUrl(aliasUrl string) (string, error) {
	smtp, err := p.postgres.Prepare("SELECT url FROM url WHERE alias = $1")
	if err != nil {
		return "", err
	}

	var resUrl string
	err = smtp.QueryRow(aliasUrl).Scan(&resUrl)

	return resUrl, err
}


func (p *Postgres) IsDublicate(alias string) bool{
	smtp, err := p.postgres.Prepare("SELECT EXISTS (SELECT alias FROM url WHERE alias = $1)")
	if err != nil {
		return true
	}

	var exists bool
	err = smtp.QueryRow(alias).Scan(&exists)
	if err != nil {
		return true
	}
	return exists
}