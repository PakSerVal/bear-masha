-- +goose Up
-- +goose StatementBegin
create table card_files
(
    id      integer generated always as identity,
    file_id integer not null,
    card_id integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table card_files;
-- +goose StatementEnd
