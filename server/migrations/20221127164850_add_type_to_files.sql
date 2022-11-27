-- +goose Up
-- +goose StatementBegin
alter table files add type text
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table files drop column type
-- +goose StatementEnd
