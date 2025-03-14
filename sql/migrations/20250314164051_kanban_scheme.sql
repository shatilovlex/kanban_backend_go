-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO kanban.project (id, name, description, archived)
VALUES (
        '7142c1a1-30d4-452c-af3e-47fb821e4646',
        'My first board',
        'Simple board',
        false
       )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE kanban.project;
SELECT 'down SQL query';
-- +goose StatementEnd
