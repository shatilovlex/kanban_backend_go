-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO kanban.tasks (id, list_id, title, description, sort, archived)
VALUES (
        'b45a1852-8cb6-4333-9262-8c0d159e1e74',
        'b45a1852-8cb6-4337-9262-8c0d159e1e74',
        'New Task',
        'New Task',
        0,
        DEFAULT
       )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
TRUNCATE kanban.list;
-- +goose StatementEnd
