package db

import (
	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/pkg/errors"
)

var Db *sql.DB

type User struct {
	Id   string
	Name string
}

func SearchUserById(id string) (User, error) {
	var u User
	row := Db.QueryRow("select id ,name from user where id = ?", id)
	err := row.Scan(&u.Id, &u.Name)
	if err != nil {
		return u, errors.Wrap(err, "search user error")
	}
	return u, nil
}

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}
