-- +goose Up
-- +goose StatementBegin
CREATE TABLE article_views (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    article_id BIGINT NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    viewed_at TIMESTAMP NOT NULL DEFAULT now(),

    UNIQUE(user_id, article_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE article_views;
-- +goose StatementEnd
