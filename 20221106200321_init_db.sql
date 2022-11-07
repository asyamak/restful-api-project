-- +goose Up

    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY, 
        data VARCHAR
    );

-- +goose Down
	DROP TABLE users;

