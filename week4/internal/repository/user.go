//这一层是数据持久层，像数据库存取，缓存的处理

package repository

import (
	"context"
	"database/sql"
	"github.com/Liusiyuan-git/Go-learning/week4/internal/domain"
	"github.com/pkg/errors"
)

var NotFound = errors.New("not found")
var Unknown = errors.New("unknown")

type repository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) domain.IUserRepo {
	return &repository{client: client}
}

func (r *repository) GetUserInfo(ctx context.Context, id int) (domain.User, error) {
	var u domain.User
	row := r.client.QueryRow("select id ,name from user where id = ?", id)
	err := row.Scan(&u.Id, &u.Name)
	if errors.Is(err, NotFound) {
		return u, errors.Wrapf(NotFound, "user not found, id: %d, err: %+v", id, err)
	}

	if err != nil {
		return u, errors.Wrapf(Unknown, "db query err: %+v, id: %d,", err, id)
	}

	return u, nil
}
