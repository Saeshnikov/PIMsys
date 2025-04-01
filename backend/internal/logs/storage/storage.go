package storage

import (
	"context"
	"database/sql"
	"fmt"
	proto "pim-sys/gen/go/logs"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(connectionString string) (*Storage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "opening database connection: ", err)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) Stop() error {
	return s.DB.Close()
}

type Sales struct {
	Price    int32
	Quantity int32
}

func (s *Storage) GetLogs(
	ctx context.Context,
	userId int32,
) (
	[]*proto.Log,
	error,
) {
	var res []*proto.Log

	stmt, err := s.DB.Prepare("SELECT log_journal.shop_id, log_journal.branch_id, log_journal.product_id, log_journal.info FROM log_journal JOIN users_shop ON log_journal.shop_id=users_shop.shop_id WHERE users_shop.users_id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		logs := &proto.Log{}
		err := rows.Scan(&logs.ShopId, &logs.BranchId, &logs.ProductId, &logs.Info)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, logs)
	}

	return res, nil
}

func (s *Storage) GetSales(
	ctx context.Context,
	TimeFrom int64,
	TimeTo int64,
	userId int32,
) (
	[]*proto.Graph,
	error,
) {
	var res []*proto.Graph

	stmt, err := s.DB.Prepare(`SELECT sales.date, sales.price, sales.quantity FROM sales
								JOIN branch ON sales.branch_id=branch.id
								JOIN users_shop ON users_shop.shop_id=branch.shop_id
								WHERE sales.date >= $1 and sales.date <= $2 and users_shop.users_id= $3`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, TimeFrom, TimeTo, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		sales := proto.Graph{}
		err := rows.Scan(&sales.Date, &sales.TotalSales, &sales.TotalQuantity)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, &sales)
	}

	return res, nil
}

func (s *Storage) GetMinDate(
	ctx context.Context,
	dateFrom int64,
) error {
	stmt, err := s.DB.Prepare("SELECT MIN(date) AS earliest_date FROM sales")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	var minDate int64

	err = stmt.QueryRowContext(ctx).Scan(&minDate)
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	if dateFrom < minDate {
		return fmt.Errorf("%s ", "DateFrom can't be less than minimal date of sales")
	}

	return nil
}
