-- +goose Up
-- +goose StatementBegin
CREATE TABLE notes (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL,
    Content TEXT NOT NULL,

    CONSTRAINT notes_category_id
        FOREIGN KEY(category_id) 
        REFERENCES categories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE notes;
-- +goose StatementEnd
