-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO kanban.list (id, project_id, name)
VALUES (
        'b45a1852-8cb6-4337-9262-8c0d159e1e74',
           '7142c1a1-30d4-452c-af3e-47fb821e4646',
           'Todo'
       ),
       (
           '27865cb8-b510-4daa-82d1-629db7be95de',
           '7142c1a1-30d4-452c-af3e-47fb821e4646',
           'In Progress'
       ),
       (
           '7142c1a1-30d4-452c-af3e-47fb821e4644',
           '7142c1a1-30d4-452c-af3e-47fb821e4646',
           'Done'
       );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE kanban.project;
SELECT 'down SQL query';
-- +goose StatementEnd
