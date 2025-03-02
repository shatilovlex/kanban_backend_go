-- +goose Up
-- +goose StatementBegin
create schema kanban;
create table if not exists kanban.project
(
    id       uuid primary key,
    name     varchar(255),
    description varchar(255),
    archived bool  default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists kanban.project;
drop schema if exists kanban;
-- +goose StatementEnd
