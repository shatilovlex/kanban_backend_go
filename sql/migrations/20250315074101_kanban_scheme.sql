-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists kanban.list
(
    id uuid primary key,
    project_id uuid,
    name varchar(255),
    sort int,
    constraint fk_list_projects foreign key (project_id) references kanban.project(id)
);
create index idx_list_project_id ON kanban.list USING btree (project_id);
create unique index idx_list_name_project_id on kanban.list (project_id, name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table kanban.list;
-- +goose StatementEnd
