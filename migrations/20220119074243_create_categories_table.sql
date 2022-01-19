-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(300) NOT NULL,

    CONSTRAINT categories_user_id
        FOREIGN KEY(user_id) 
        REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories;
-- +goose StatementEnd
