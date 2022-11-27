-- +goose Up
-- +goose StatementBegin
create table cards
(
    id          integer generated always as identity,
    title       text not null,
    description text not null,
    status      text not null,
    deadline_at timestamp not null default current_timestamp,
    created_at  timestamp not null default current_timestamp
);
create index cards_status_idx on cards (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table cards;
-- +goose StatementEnd
