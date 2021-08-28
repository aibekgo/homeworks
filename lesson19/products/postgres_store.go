package products

import (
"database/sql"
"errors"
"fmt"
"github.com/google/uuid"
_ "github.com/lib/pq"
)

type productsPostgreStore struct {
	db *sql.DB
}

func NewProductPostgresStore(host, database, user, password, port string) (ProductStore, error) {
	createTableQuery := `
		create table if not exists products(
			id text,
			name text,
			price int,
			primary key(id)
		);
	`
	params := "sslmode=disable"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", user, password, host, port, database, params)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}
	return &productsPostgreStore{db: db}, nil
}

func (u *productsPostgreStore) Create(product *Product) (*Product, error) {
	product.Id = uuid.New().String()
	queryInsert := "insert into users (id, name, price) values ($1, $2, $3)"
	result, err := u.db.Exec(queryInsert, product.Id, product.Name,  product.Price)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("error during creatring of product")
	}
	return product, nil
}

func (u *productsPostgreStore) List() ([]Product, error) {
	product := []Product{}
	selectQuery := "select id, name, price from product"
	rows, err := u.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := Product{}
		err = rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		product = append(product, item)
	}
	return product, nil
}

func (u *productsPostgreStore) Delete(id string) error {
	result, err := u.db.Exec("delete from users where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("error in deleting")
	}
	return nil
}
