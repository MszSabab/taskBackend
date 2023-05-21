package repository

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/taskBackend/config"
	"time"
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

func (pr *PgRepository) CreatePost(likes int, Title, Content, MediaURL string) error {
	_, err := pr.db.Exec(
		`INSERT INTO post (likes, title, content, media_url) VALUES ($1, $2, $3, $4)`,
		likes,
		Title,
		Content,
		MediaURL,
	)
	return err
}

type Comment struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	UserName string `db:"user_name"`
}

type getAllPosts struct {
	Id        string    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	MediaURL  string    `db:"media_url"`
	Likes     int       `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (pr *PgRepository) GetAllPosts(limit, offset int) (*[]getAllPosts, error) {
	getAll := make([]getAllPosts, 0)
	err := pr.db.Select(
		&getAll,
		`
		SELECT *
		FROM post
		LIMIT $1 OFFSET $2
		`,
		limit,
		offset,
	)
	return &getAll, err
}

func (pr *PgRepository) LikePost(id int) error {
	_, err := pr.db.Exec(
		`
		UPDATE post
    	SET likes = likes + 1
		WHERE id = $1
		`,
		id,
	)
	return err
}

func (pr *PgRepository) UnLikePost(id int) error {
	_, err := pr.db.Exec(
		`
		UPDATE post
    	SET likes = likes - 1
		WHERE id = $1
		`,
		id,
	)
	return err
}

func (pr *PgRepository) AddComment(content, username string, postID int) error {
	_, err := pr.db.Exec(
		`
		INSERT INTO comments
		(content, user_name, post_id) VALUES ($1, $2, $3)
		`,
		content,
		username,
		postID,
	)
	return err
}
