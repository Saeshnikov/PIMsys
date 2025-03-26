package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	proto "pim-sys/gen/go/products"

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

func (s *Storage) CreateProduct(
	ctx context.Context,
	content *proto.ProductInfo,
) (int32, error) {
	// Добавление нового шопа и связи с юзером
	stmt, err := s.DB.Prepare(
		"INSERT INTO product (category_id, status, branch_id, name, price, amount) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, content.GetCategoryId(), content.GetStatus(), content.GetBranchId(), content.GetName(), content.GetPrice(), content.GetAmount())
	if row.Err() != nil {
		return 0, fmt.Errorf("%s: %w", "executing query: ", err)
	}
	var productId int32
	err = row.Scan(&productId)
	if row.Err() != nil {
		return 0, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return productId, nil
}

type Attributes struct {
	AttributeId int32  `json:"attribute_id"`
	ValueText   string `json:"value_text"`
	ValueNumber int32  `json:"value_number"`
	ValueBool   bool   `json:"value_bool"`
}

func (s *Storage) AlterAttributes(
	ctx context.Context,
	productId int32,
	attr *proto.Attribute,
) error {
	// Добавление нового шопа и связи с юзером
	stmt, err := s.DB.Prepare(
		`UPDATE product_attribute_value SET value_text=$3, value_number=$4, value_boolean=$5 WHERE product_id=$1 AND attribute_id=$2;`,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, productId, attr.Id, attr.ValueText, attr.ValueNumber, attr.ValueBool).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	stmt, err = s.DB.Prepare(
		`INSERT INTO product_attribute_value (product_id, attribute_id, value_text, value_number, value_boolean) 
       		SELECT $1, $2, $3, $4, $5 
       		WHERE NOT EXISTS (SELECT * FROM product_attribute_value WHERE product_id=$1 AND attribute_id=$2);`,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, productId, attr.Id, attr.ValueText, attr.ValueNumber, attr.ValueBool).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) AlterProduct(
	ctx context.Context,
	content *proto.ProductInfoWithId,
) error {
	stmt, err := s.DB.Prepare("UPDATE product SET status=$1, branch_id=$2, name=$3, amount=$4, price=$5 WHERE id=$6")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(
		ctx,
		content.GetProduct().GetStatus(),
		content.GetProduct().GetBranchId(),
		content.GetProduct().GetName(),
		content.GetProduct().GetAmount(),
		content.GetProduct().GetPrice(),
		content.GetProductId(),
	).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) DeleteProduct(
	ctx context.Context,
	content *proto.DeleteProductRequest,
) error {
	stmt, err := s.DB.Prepare("DELETE FROM product WHERE id=$1") // Нужна валидация на то, что такой ид существует
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, content.ProductId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) ListProducts(
	ctx context.Context,
	userId int32,
) (
	[]*proto.ProductInfoWithId,
	error,
) {

	var res []*proto.ProductInfoWithId

	stmt, err := s.DB.Prepare("SELECT product.id,category_id,status,branch_id,name,price,amount FROM product JOIN users_shop ON (SELECT shop_id from branch where branch.id=product.branch_id)=users_shop.shop_id WHERE users_shop.users_id=$1") // Добавить поиск с джоином по юзеру
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		productInfo := proto.ProductInfoWithId{Product: &proto.ProductInfo{}}

		err := rows.Scan(&productInfo.ProductId, &productInfo.Product.CategoryId, &productInfo.Product.Status, &productInfo.Product.BranchId, &productInfo.Product.Name, &productInfo.Product.Price, &productInfo.Product.Amount)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}

		attr, err := s.DB.Prepare("SELECT attribute_id,value_text,value_number,value_boolean FROM product_attribute_value WHERE product_id=$1") // Добавить поиск с джоином по юзеру
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "creating query: ", err)
		}
		defer stmt.Close()

		rowsAttr, err := attr.QueryContext(ctx, productInfo.ProductId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "executing query: ", err)
		}

		for rowsAttr.Next() {
			attribute := &proto.Attribute{}

			err := rowsAttr.Scan(&attribute.Id, &attribute.ValueText, &attribute.ValueNumber, &attribute.ValueBool)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
			}

			productInfo.Product.Attributes = append(productInfo.Product.Attributes, attribute)
		}

		res = append(res, &productInfo)
	}

	return res, nil
}

func (s *Storage) SellProduct(
	ctx context.Context,
	content *proto.SellProductRequest,
) error {
	stmt, err := s.DB.Prepare("UPDATE product SET amount=amount-$1 WHERE id=$2")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(
		ctx,
		content.GetAmount(),
		content.GetProductId(),
	).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	insertSales, err := s.DB.Prepare("INSERT INTO sales (date,branch_id,product_id,price,quantity) VALUES ($1,(SELECT branch_id FROM product WHERE id=$2),$2,(SELECT price FROM product WHERE id=$2),$3)")
	if err != nil {
		return fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer insertSales.Close()

	err = insertSales.QueryRowContext(
		ctx,
		time.Now().Unix(),
		content.GetProductId(),
		content.GetAmount(),
	).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "executing query: ", err)
	}

	return nil
}

func (s *Storage) GetAccessableBranchIds(
	ctx context.Context,
	userId int32,
) (
	[]int32,
	error,
) {

	var res []int32

	stmt, err := s.DB.Prepare("SELECT branch.id FROM branch JOIN users_shop ON branch.shop_id=users_shop.shop_id WHERE users_shop.users_id=$1") // Добавить поиск с джоином по юзеру
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "creating query: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "executing query: ", err)
	}

	for rows.Next() {
		var productId int32
		err := rows.Scan(&productId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scan query result: ", err)
		}
		res = append(res, productId)
	}

	return res, nil
}
