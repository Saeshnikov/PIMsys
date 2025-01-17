package storage

import (
	"context"
	"database/sql"
	"fmt"

	proto "pim-sys/gen/go/template"

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

func (s *Storage) CreateTemplate(
	ctx context.Context,
	branch_id int32,
	name string,
	description string,
	attributes []*proto.AttributeInfo,
) error {
	stmt, err := s.db.Prepare(
		`with rows as (INSERT INTO category (name, description, is_unique) VALUES($1, $2, true) RETURNING id)
		 INSERT INTO category_branch (branch_id, category_id) VALUES ($3, (SELECT id FROM rows)) RETURNING id`,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error in CreateTemplate (step: prepare query INSERT category)", err)
	}
	defer stmt.Close()

	var category_id int32
	err = stmt.QueryRowContext(ctx, name, description, branch_id).Scan(&category_id)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error in CreateTemplate query (step: execute query INSERT category): ", err)
	}

	stmt, err = s.db.Prepare(
		`with rows as (INSERT INTO attribute (type, is_value_required, is_unique, name, description) VALUES ($1, $2, $3, $4, $5) RETURNING id)
		 INSERT INTO category_attribute(category_id, attribute_id) VALUES ($6, (SELECT id FROM rows))`,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", "Error in CreateTemplate (step: prepare query INSERT attribute)", err)
	}
	defer stmt.Close()

	for _, attr := range attributes {
		err = stmt.QueryRowContext(ctx, attr.Type, attr.IsValueRequired, attr.IsUnique, attr.Name, attr.Description, category_id).Err()
		if err != nil {
			return fmt.Errorf("%s: %w", "Error in CreateTemplate query (step: execute query INSERT attribute): ", err)
		}
	}

	return nil
}

func (s *Storage) DeleteTemplate(
	ctx context.Context,
	templateId int32,
) error {
	stmt, err := s.db.Prepare("DELETE FROM category WHERE id=$1") // Нужна валидация на то, что такой ид существует
	if err != nil {
		return fmt.Errorf("%s: %w", "Error in DeleteTemplate (step: prepare)", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, templateId).Err()
	if err != nil {
		return fmt.Errorf("%s: %w", "Error in DeleteTemplate (step: execute)", err)
	}

	return nil
}

func (s *Storage) ListTemplates(
	ctx context.Context,
	branch_id int32,
) (
	[]*proto.TemplateInfo,
	error,
) {
	var res []*proto.TemplateInfo

	/* Get all categories on requested branch*/
	stmtCategories, err := s.db.Prepare(
		`SELECT category.id,name,description FROM category 
		JOIN category_branch ON category.id=category_branch.id WHERE category_branch.branch_id=$1`,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: preparing SELECT category query)", err)
	}
	defer stmtCategories.Close()

	categoriesRows, err := stmtCategories.QueryContext(ctx, branch_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: executing SELECT category query)", err)
	}

	/* Get all attributes on every category id */
	stmt, err := s.db.Prepare(
		`SELECT attribute_id FROM category_attribute WHERE category_attribute.category_id=$1`,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: preparing1 SELECT category query)", err)
	}
	defer stmt.Close()

	attrMt, err := s.db.Prepare(
		`SELECT type, is_value_required, is_unique, name, description FROM attribute WHERE id = $1`,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: preparing2 SELECT category query)", err)
	}
	defer attrMt.Close()

	for categoriesRows.Next() {
		var attributesArray []*proto.AttributeInfo
		categoryElem := proto.TemplateInfo{}
		// Get category id
		err := categoriesRows.Scan(&categoryElem.TemplateId, &categoryElem.Name, &categoryElem.Description)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "scanning categories rows: ", err)
		}

		// Get attribute id's on category id
		attributesIdRows, err := stmt.QueryContext(ctx, categoryElem.TemplateId)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: executing SELECT category_attribute query)", err)
		}
		for attributesIdRows.Next() {
			var currentAttributeId int32

			// Parse attribute id
			err = attributesIdRows.Scan(&currentAttributeId)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: parsing attributeRow id)", err)
			}

			// Get attribute fields on id
			attributeRows, err := attrMt.QueryContext(ctx, currentAttributeId)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: executing SELECT attribute query)", err)
			}

			// Parse attribute rows
			attrInfo := &proto.AttributeInfo{
				Id: currentAttributeId,
			}
			for attributeRows.Next() {
				err = attributeRows.Scan(
					&attrInfo.Type,
					&attrInfo.IsValueRequired,
					&attrInfo.IsUnique,
					&attrInfo.Name,
					&attrInfo.Description,
				)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", "Error in ListTemplates (step: parsing attribute info row)", err)
				}
				attributesArray = append(attributesArray, attrInfo)
			}
		}

		categoryElem.Attributes = attributesArray
		res = append(res, &categoryElem)
	}

	return res, nil
}

func (s *Storage) GetUserListBranches(
	ctx context.Context,
	user_id int32,
) (
	[]int32,
	error,
) {
	var resList []int32
	/* Get all shops on requested user*/
	stmt, err := s.db.Prepare(
		`SELECT b.id
		FROM users u
		JOIN users_shop us ON u.id = us.users_id
		JOIN shop s ON us.shop_id = s.id
		JOIN branch b ON s.id = b.shop_id
		WHERE u.id = $1`,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in GetUserListBranches (step: preparing SELECT query on user_shop)", err)
	}

	defer stmt.Close()

	branchIdRows, err := stmt.QueryContext(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", "Error in GetUserListBranches (step: executing SELECT query on user_shop)", err)
	}

	for branchIdRows.Next() {
		var currentBranchId int32
		branchIdRows.Scan(&currentBranchId)
		resList = append(resList, currentBranchId)
	}

	return resList, err
}
