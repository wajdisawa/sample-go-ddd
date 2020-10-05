package infrastructure

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sample-go-ddd/domain/entity"
	"sample-go-ddd/domain/repository"
)

type userRepository struct {
	conn *sql.DB
}

func (u *userRepository) Get(ctx context.Context, id int) (*entity.User, error) {
	row, err := u.fetchOne(ctx, "select id, name, email from users where id=?", id)
	if err != nil {
		return nil, err
	}
	usr := &entity.User{}
	row.Scan()
	err = row.Scan(&usr.Id, &usr.Name, &usr.Email)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	rows, err := u.fetch(ctx, "select id, name from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	usrs := make([]*entity.User, 0)
	for rows.Next() {
		usr := &entity.User{}
		err = rows.Scan(&usr.Id, &usr.Name, &usr.Email)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}
	return usrs, nil
}

func (u userRepository) Save(ctx context.Context, user *entity.User) error {
	stmt, err := u.conn.Prepare("insert into users (name, email) values (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Email)
	return err
}
func (u *userRepository) fetchOne(ctx context.Context, s string, args ...interface{}) (*sql.Row, error) {
	query, err := u.conn.Prepare(s)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	return query.QueryRowContext(ctx, args...), nil
}

func (u *userRepository) fetch(ctx context.Context, s string, args ...interface{}) (*sql.Rows, error) {
	query, err := u.conn.Prepare(s)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	return query.QueryContext(ctx, args...)
}

func NewUserRepository(connection *sql.DB) repository.UserRepository {
	return &userRepository{conn: connection}
}
