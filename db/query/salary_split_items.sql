-- name: AddSalarySplitItem :one
INSERT INTO "Finance"."Salary_split_items" (
    split_id,
    category_name,
    amount,
    move_to,
    is_transferred
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSalarySplitItemsBySplitId :many
SELECT * FROM "Finance"."Salary_split_items"
WHERE split_id = $1;    

-- name: MarkSalarySplitItemAsTransferredById :exec
UPDATE "Finance"."Salary_split_items"
SET is_transferred = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;   

-- name: DeleteSalarySplitItemsBySplitId :exec
DELETE FROM "Finance"."Salary_split_items"
WHERE split_id = $1;

-- name: GetAllSalarySplitItemsByUserId :many
SELECT ssi.* FROM "Finance"."Salary_split_items" ssi
JOIN "Finance"."Salary_splits" ss ON ssi.split_id = ss.id
WHERE ss.user_id = $1
ORDER BY ss.month DESC, ssi.category_name;

-- name: UpdateSalarySplitItemAmountById :exec
UPDATE "Finance"."Salary_split_items"
SET amount = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;
