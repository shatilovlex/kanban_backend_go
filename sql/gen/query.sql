-- name: ProjectCreate :exec
INSERT INTO pg_storage.kanban.project (id, name, description)
VALUES ($1, $2, $3);
