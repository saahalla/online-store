package database

import "github.com/jmoiron/sqlx"

func Migration(db *sqlx.DB) error {
	err := CreateTableUsers(db)
	if err != nil {
		return err
	}

	err = CreateTableProducts(db)
	if err != nil {
		return err
	}

	err = CreateTableCategories(db)
	if err != nil {
		return err
	}

	err = AlterTableProducts(db)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableUsers(db *sqlx.DB) error {
	users := `
		CREATE TABLE IF NOT EXISTS users (
			id int auto_increment not null primary key,
			username varchar(255) not null,
			password varchar(255) not null,
			email varchar(255) not null,
			phone varchar(255),
			user_role_id int not null default 0,
			created_at timestamp not null default current_timestamp,
			created_by varchar(255),
			modified_at timestamp not null default current_timestamp,
			modified_by varchar(255)
		  );
	`

	_, err := db.Exec(users)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableProducts(db *sqlx.DB) error {
	products := `
		CREATE TABLE IF NOT EXISTS products (
			id int AUTO_INCREMENT not null PRIMARY KEY,
			product_name varchar(255) not null,
			stock int not null default 0,
			price decimal(25,2) not null default 0,
			image varchar(255),
			created_at timestamp not null default current_timestamp,
			created_by varchar(255),
			modified_at timestamp not null default current_timestamp,
			modified_by varchar(255)
		);
	`

	_, err := db.Exec(products)
	if err != nil {
		return err
	}

	return nil
}

func AlterTableProducts(db *sqlx.DB) error {
	products := `
		ALTER TABLE products add category_id int; 
	`

	_, err := db.Exec(products)
	if err != nil {
		return err
	}

	return nil
}

func CreateTableCategories(db *sqlx.DB) error {
	categories := `
		CREATE TABLE IF NOT EXISTS categories (
			id int AUTO_INCREMENT not null PRIMARY KEY,
			category_name varchar(255) not null,
			created_at timestamp not null default current_timestamp,
			created_by varchar(255),
			modified_at timestamp not null default current_timestamp,
			modified_by varchar(255)
		);
	  `

	_, err := db.Exec(categories)
	if err != nil {
		return err
	}

	return nil
}
