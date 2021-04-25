package repo

import (
	"github.com/jmoiron/sqlx"
)

type Group struct {
	groupID    int `db:"group_id"`
	telegramID int `db:"telegram_id"`
}

type User struct {
	userID    int    `db:"user_id"`
	birthdate string `db:"birthdate"`
}

type Repo struct {
	conn *sqlx.DB
}

func InitRepo(conn *sqlx.DB) *Repo {
	return &Repo{conn: conn}
}
