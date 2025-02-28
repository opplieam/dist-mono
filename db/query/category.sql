-- name: GetCategoryByID :one
SELECT id, name
FROM categorys
WHERE user_id = $1;