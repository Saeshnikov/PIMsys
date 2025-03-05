package storage

import (
	"context"
	"database/sql"
	"fmt"

	proto "pim-sys/gen/go/shop"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB DB
}

type DB interface {
	Prepare(query string) (*sql.Stmt, error)
	Close() error
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

func (s *Storage) CreateShop(
	ctx context.Context,
	userID int,
	name string,
	description string,
	url string,
) error {
	// Добавление нового шопа и связи с юзером
	stmt, err := s.DB.Prepare(
		"with rows as (INSERT INTO shop (name, avatar_url, description) VALUES($1, $2, $3) RETURNING id) INSERT INTO users_shop (users_id,shop_id) VALUES ($4, (SELECT id FROM rows))",
	)

	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, name, url, description, userID).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) AlterShop(
	ctx context.Context,
	shopId int32,
	name string,
	description string,
	url string,
) error {
	stmt, err := s.DB.Prepare("UPDATE shop SET name=$1, avatar_url=$2, description=$3 WHERE id=$4")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, name, url, description, shopId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) DeleteShop(
	ctx context.Context,
	shopId int32,
) error {
	stmt, err := s.DB.Prepare("DELETE FROM shop WHERE id=$1") // Нужна валидация на то, что такой ид существует
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, shopId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) ListShops(
	ctx context.Context,
	userId int32,
) (
	[]*proto.ShopInfo,
	error,
) {

	var res []*proto.ShopInfo

	stmt, err := s.DB.Prepare("SELECT shop.id,name,description,avatar_url FROM shop JOIN users_shop ON shop.id=users_shop.shop_id WHERE users_shop.users_id=$1") // Добавить поиск с джоином по юзеру
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		shopInfo := proto.ShopInfo{}
		err := rows.Scan(&shopInfo.ShopId, &shopInfo.Name, &shopInfo.Description, &shopInfo.Url)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		fmt.Println(&shopInfo)
		res = append(res, &shopInfo)
	}

	return res, nil
}
