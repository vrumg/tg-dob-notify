package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func (r *Repo) InsertUserDateWithChannel(groupTgID int64, login string, date string) error {

	// start db transaction
	// we know, that each rows response contain exact 1 row
	// dont forget to close rows before next query due to lib/pq realization
	tx, err := r.conn.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// insert or update user
	user := &user{Birthdate: date, Login: login}
	rows, err := tx.NamedQuery(insertOrUpdateUser, user)
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return err
	}
	// check insert/update response
	rows.Next()
	err = rows.StructScan(user)
	_ = rows.Close()
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return err
	}
	if user.UserID == 0 {
		_ = tx.Rollback()
		return errors.New("")
	}

	// insert group
	group := &group{TelegramID: groupTgID}
	rows, err = tx.NamedQuery(insertGroup, group)
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return err
	}
	rows.Next()
	err = rows.StructScan(group)
	_ = rows.Close()
	if err != nil {
		log.Println(err)
		return errors.New("")
	}
	if group.GroupID == 0 {
		_ = tx.Rollback()
		return errors.New("")
	}

	// insert user-group relation
	userToGroup := &userToGroup{GroupID: group.GroupID, UserID: user.UserID}
	rows, err = tx.NamedQuery(insertUserGroupRelation, userToGroup)
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return err
	}
	rows.Next()
	err = rows.StructScan(userToGroup)
	_ = rows.Close()
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return errors.New("")
	}
	if userToGroup.UserID == 0 {
		_ = tx.Rollback()
		return errors.New("")
	}

	_ = tx.Commit()

	return nil
}

const insertOrUpdateUser = `
insert into public."user" as u (login, birthdate)
values (:login, :birthdate)
on conflict
on constraint user_pk
do update
set birthdate = :birthdate
returning user_id;`

const insertGroup = `
with e as(
    insert into public."group" (telegram_id)
        values (:telegram_id)
        on conflict
        on constraint group_pk
        do nothing
        returning group_id
)
select * from e
union
select group_id from public."group"
where telegram_id = :telegram_id;`

const insertUserGroupRelation = `
with e as(
    insert into public."user_group" (user_id, group_id)
        values (:user_id, :group_id)
        on conflict
        on constraint user_group_pk
        do nothing
        returning user_group_id
)
select * from e
union
select user_group_id from public."user_group"
where user_id = :user_id and group_id = :group_id;`
