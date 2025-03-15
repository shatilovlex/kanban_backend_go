-- name: ProjectCreate :exec
INSERT INTO pg_storage.kanban.project (id, name, description)
VALUES ($1, $2, $3);

-- name: ProjectArchive :exec
UPDATE pg_storage.kanban.project SET archived=$2 WHERE id=$1;

-- name: ProjectList :many
SELECT id, name, description FROM pg_storage.kanban.project WHERE archived IS FALSE;

-- name: Board :many
SELECT id, name FROM pg_storage.kanban.list WHERE project_id=$1;