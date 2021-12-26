package main

import (
	"error/db"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	user, err := db.SearchUserById("123456")
	if err != nil {
		fmt.Printf("db err : %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: %+v", err)
		return
	}
	fmt.Println("user: ", user)
	db.Db.Close()
}
