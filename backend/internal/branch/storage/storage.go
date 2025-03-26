package storage

import (
	"context"
	"database/sql"
	"fmt"

	proto "pim-sys/gen/go/branch"

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

func (s *Storage) CreateBranch(
	ctx context.Context,
	name string,
	shopID int32,
	description string,
	address string,
	site string,
	branch_type string,
) error {
	// Добавление нового шопа и связи с юзером
	stmt, err := s.DB.Prepare(
		"INSERT INTO branch (name, description, address, site, type, shop_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
	)
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, name, description, address, site, branch_type, shopID).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) AlterBranch(
	ctx context.Context,
	name string,
	branchId int32,
	description string,
	address string,
	site string,
) error {
	stmt, err := s.DB.Prepare("UPDATE branch SET name=$1, description=$2, address=$3, site=$4 WHERE id=$5")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, name, description, address, site, branchId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) DeleteBranch(
	ctx context.Context,
	branchId int32,
) error {
	stmt, err := s.DB.Prepare("DELETE FROM branch WHERE id=$1") // Нужна валидация на то, что такой ид существует
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, branchId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) ListBranches(
	ctx context.Context,
	shopId int32,
) (
	[]*proto.BranchInfo,
	error,
) {

	var res []*proto.BranchInfo

	stmt, err := s.DB.Prepare("SELECT branch.id, name, description, address, site, type FROM branch WHERE shop_id=$1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, shopId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		branchInfo := proto.BranchInfo{}
		err := rows.Scan(&branchInfo.BranchId, &branchInfo.Name, &branchInfo.Description, &branchInfo.Address, &branchInfo.Site, &branchInfo.BranchType)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, &branchInfo)
	}

	return res, nil
}

func (s *Storage) ListShops(
	ctx context.Context,
	userId int32,
) (
	[]int32,
	error,
) {

	var res []int32

	stmt, err := s.DB.Prepare("SELECT shop.id FROM shop JOIN users_shop ON shop.id=users_shop.shop_id WHERE users_shop.users_id=$1") // Добавить поиск с джоином по юзеру
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}
	var shopid int32
	for rows.Next() {

		err := rows.Scan(&shopid)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, shopid)
	}

	return res, nil
}

func (s *Storage) GetShopId(
	ctx context.Context,
	branchId int32,
) (
	int32,
	error,
) {
	stmt, err := s.DB.Prepare("SELECT shop_id FROM branch WHERE id=$1")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	var shopid int32

	err = stmt.QueryRowContext(ctx, branchId).Scan(&shopid)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return shopid, nil
}

// // User returns user by email.
// func (s *Storage) User(ctx context.Context, email string) (User, error) {
// 	const op = "storage.postgres.User"

// 	stmt, err := s.db.Prepare("SELECT id, email, password, isAdmin FROM users WHERE email = $1")
// 	if err != nil {
// 		return User{}, fmt.Errorf("%s: %w", op, err)
// 	}
// 	defer stmt.Close()

// 	row := stmt.QueryRowContext(ctx, email)

// 	var user User
// 	err = row.Scan(&user.ID, &user.Email, &user.PassHash, &user.IsAdmin)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return User{}, fmt.Errorf("%s: %w", op, auth_errors.ErrUserNotFound)
// 		}
// 		return User{}, fmt.Errorf("%s: %w", op, err)
// 	}

// 	return user, nil
// }

// // IsAdmin checks if the user is an admin.
// func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
// 	const op = "storage.postgres.IsAdmin"

// 	stmt, err := s.db.Prepare("SELECT isAdmin FROM users WHERE id = $1")
// 	if err != nil {
// 		return false, fmt.Errorf("%s: %w", op, err)
// 	}
// 	defer stmt.Close()

// 	row := stmt.QueryRowContext(ctx, userID)

// 	var isAdmin bool
// 	err = row.Scan(&isAdmin)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return false, fmt.Errorf("%s: %w", op, auth_errors.ErrUserNotFound)
// 		}
// 		return false, fmt.Errorf("%s: %w", op, err)
// 	}

// 	return isAdmin, nil
// }
