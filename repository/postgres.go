package repository

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/taskBackend/config"
	"log"
)

type PgRepository struct {
	db *sqlx.DB
}

func NewPgRepository(cfg *config.Config) (*PgRepository, error) {
	db, err := sqlx.Connect("pgx", cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}
	return &PgRepository{
		db: db,
	}, nil
}

func (pr *PgRepository) Create(Title, Content, MediaURL string, ID int) {

	err := pr.db.QueryRow("INSERT INTO posts (title, content, media_url) VALUES ($1, $2, $3) RETURNING id", Title, Content, MediaURL).Scan(&ID)
	if err != nil {
		log.Println(err)
		//w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
