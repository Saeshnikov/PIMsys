package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "opening database connection: ", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

type Logs struct {
	ShopId    int32
	BranchId  int32
	ProductId int32
	Info      string
}

func (s *Storage) GetLogs(
	ctx context.Context,
	userId int32,
) (
	[]*Logs,
	error,
) {
	var res []*Logs

	stmt, err := s.db.Prepare("SELECT shop_id,branch_id,product_id,info FROM log_journal JOIN users_shop ON log_journal.shop_id=users_shop.shop_id WHERE users_shop.users_id=$1") // Добавить поиск с джоином по юзеру
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		logs := Logs{}
		err := rows.Scan(&logs.ShopId, &logs.BranchId, &logs.ProductId, &logs.Info)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, &logs)
	}

	return res, nil
}
