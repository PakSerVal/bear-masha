-- +goose Up
-- +goose StatementBegin
create table files
(
    id   integer generated always as identity,
    path text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table files;
-- +goose StatementEnd
