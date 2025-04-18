package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"mystore/internal/models"

	"go.uber.org/zap"
)

type productRepo struct {
	db  *sql.DB
	Log *zap.Logger
}

type ProductRepo interface {
	GetAllProduct() ([]*models.Product, error)
	GetById(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

func NewProductRepo(db *sql.DB, logger *zap.Logger) ProductRepo {
	return &productRepo{db: db,
		Log: logger}
}

func (r *productRepo) GetAllProduct() ([]*models.Product, error) {
	query := "SELECT id, name, description, price, quantity, created_at from products"
	rows, err := r.db.Query(query)
	if err != nil {
		r.Log.Error("Failed to get all products", zap.Error(err))
		return nil, errors.New("failed to get all products")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			r.Log.Error("Failed to close rows", zap.Error(err))
		}
	}(rows)

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt); err != nil {
			r.Log.Error("Failed to get all products", zap.Error(err))
			return nil, errors.New("failed to get all products")
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		r.Log.Error("Failed to get all products", zap.Error(err))
		return nil, errors.New("failed to get all products")
	}
	return products, nil
}

func (r *productRepo) GetById(id int) (*models.Product, error) {
	product := &models.Product{}
	query := "SELECT id,  name, description, price, quantity, created_at FROM products WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)

	if err != nil {
		r.Log.Error("Failed to get product by ID", zap.Error(err))
		return nil, errors.New("failed to get product by id")

	}

	return product, nil
}

func (r *productRepo) Create(product *models.Product) error {
	err := r.db.QueryRow(
		`INSERT INTO products (name, description, price, quantity)
   VALUES ($1,$2,$3,$4) RETURNING id, created_at`,
		product.Name, product.Description, product.Price, product.Quantity,
	).Scan(&product.ID, &product.CreatedAt)

	if err != nil {
		r.Log.Error("Failed to insert product", zap.Error(err))
		return fmt.Errorf("productRepo.Create: %w", err)
	}
	return nil
}

func (r *productRepo) Update(p *models.Product) error {
	query := `
        UPDATE products
        SET name=$1, description=$2, price=$3, quantity=$4
        WHERE id=$5
        RETURNING created_at`
	if err := r.db.QueryRow(query,
		p.Name, p.Description, p.Price, p.Quantity, p.ID,
	).Scan(&p.CreatedAt); err != nil {
		r.Log.Error("productRepo.Update QueryRow failed", zap.Error(err))
		return fmt.Errorf("productRepo.Update: %w", err)
	}
	return nil
}

func (r *productRepo) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		r.Log.Error("Failed to delete product", zap.Error(err))
		return fmt.Errorf("failed to delete product %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		r.Log.Error("Failed to delete product, no rows affected", zap.Error(err))
		return errors.New("failed to delete product")
	}
	return nil
}
