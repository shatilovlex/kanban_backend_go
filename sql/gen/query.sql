-- name: ProjectCreate :exec
INSERT INTO pg_storage.kanban.project (id, name, description)
VALUES ($1, $2, $3);

-- name: ProjectArchive :exec
UPDATE pg_storage.kanban.project SET archived=$2 WHERE id=$1;