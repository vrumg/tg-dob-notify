package repo

import (
	"github.com/jmoiron/sqlx"
)

type group struct {
	GroupID    int64 `db:"group_id"`
	TelegramID int64 `db:"telegram_id"`
}

type user struct {
	UserID    int64  `db:"user_id"`
	Login     string `db:"login"`
	Birthdate string `db:"birthdate"`
}

type userToGroup struct {
	UserToGroupID int64 `db:"user_group_id"`
	UserID        int64 `db:"user_id"`
	GroupID       int64 `db:"group_id"`
}

type Repo struct {
	conn *sqlx.DB
}

func InitRepo(conn *sqlx.DB) *Repo {
	return &Repo{conn: conn}
}
