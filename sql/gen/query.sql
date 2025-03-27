-- name: ProjectCreate :exec
INSERT INTO kanban.project (id, name, description)
VALUES ($1, $2, $3);

-- name: ProjectArchive :exec
UPDATE kanban.project SET archived=$2 WHERE id=$1;

-- name: ProjectList :many
SELECT id, name, description FROM kanban.project WHERE archived IS FALSE;

-- name: Board :many
SELECT id, name, project_id, sort FROM kanban.list WHERE project_id=$1;

-- name: ListAdd :exec
INSERT INTO kanban.list (id, project_id, name, sort) VALUES ($1, $2, $3, $4);

-- name: ListRemove :exec
DELETE FROM kanban.list WHERE id=$1;

-- name: SaveListOrder :exec
UPDATE kanban.list SET sort=$3 WHERE id=$1 AND project_id=$2;

-- name: RenameList :exec
UPDATE kanban.list SET name=$2 WHERE id=$1;

-- name: TaskCreate :exec
INSERT INTO kanban.tasks (id, list_id, title, description, sort) VALUES ($1, $2, $3, $4, $5);

-- name: TaskUpdate :exec
UPDATE kanban.tasks SET list_id=$2, title=$3, description=$4, sort=$5 WHERE id=$1;