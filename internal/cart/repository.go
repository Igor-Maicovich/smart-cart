package cart

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddItem(ctx context.Context, item Item) error {

	query := `
	INSERT INTO cart_items (product_id, name, price, quantity)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		item.Name,
		item.Price,
		item.Quantity,
	)

	return err
}

func (r *Repository) GetAll() ([]Item, error) {
	rows, err := r.db.Query("SELECT id, product_id, name, price, quantity FROM cart_items")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
		}
	}()

	var items []Item
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Price, &it.Quantity); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

func (r *Repository) Update(id int, input Item) (Item, error) {
	_, err := r.db.Exec("UPDATE cart_items SET name=$1, price=$2, quantity=$3 WHERE id=$4",
		input.Name, input.Price, input.Quantity, id)
	if err != nil {
		return Item{}, err
	}
	input.ID = id
	return input, nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM cart_items WHERE id=$1", id)
	return err
}
