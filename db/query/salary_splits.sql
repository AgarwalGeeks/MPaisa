-- name: AddSalarySplit :one
INSERT INTO "Finance"."Salary_splits" (
    user_id,
    total_salary,
    month,
    notes,
    is_fully_transferred
) VALUES (
  $1, $2, $3, $4, $5
) 
RETURNING *;

-- name: GetSalarySplitsByUserId :many
SELECT * FROM "Finance"."Salary_splits"
WHERE user_id = $1
ORDER BY month DESC;

-- name: GetSalarySplitById :one
SELECT * FROM "Finance"."Salary_splits"
WHERE id = $1;

-- name: MarkSalarySplitAsFullyTransferredById :exec
UPDATE "Finance"."Salary_splits"
SET is_fully_transferred = TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteSalarySplitById :exec
DELETE FROM "Finance"."Salary_splits"
WHERE id = $1;

-- name: UpDateSalarySplitTotalById :exec
UPDATE "Finance"."Salary_splits"
SET total_salary = $1, updated_at = CURRENT_TIMESTAMP
WHERE id = $2;