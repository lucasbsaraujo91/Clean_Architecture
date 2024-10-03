package database

import (
	"Clean_Architecture/internal/entity"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) FindAll(page, limit int, sort string) ([]entity.Order, error) {
	var orders []entity.Order
	var err error

	// Validação do sort
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	// Construindo a query base
	query := "SELECT id, price, tax, final_price FROM orders ORDER BY id " + sort

	// Adicionando paginação, se necessário
	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		query += " LIMIT ? OFFSET ?"
		rows, err := r.Db.Query(query, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var order entity.Order
			if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
				return nil, err
			}
			orders = append(orders, order)
		}
	} else {
		// Caso não tenha paginação
		rows, err := r.Db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var order entity.Order
			if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
				return nil, err
			}
			orders = append(orders, order)
		}
	}

	return orders, err
}
