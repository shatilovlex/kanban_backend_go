-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists kanban.tasks
(
    id uuid primary key,
    list_id uuid not null,
    title varchar(255),
    description varchar(255),
    sort int,
    archived bool default false not null,
    constraint fk_tasks_list foreign key (list_id) references kanban.list(id)
    );
create index idx_task_list_id ON kanban.tasks USING btree (list_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table kanban.tasks;
-- +goose StatementEnd
