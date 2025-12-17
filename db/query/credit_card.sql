-- name: AddCreditCard :one
INSERT INTO "Finance"."Credit_Card" (
    bank_name,
    card_name,
    card_number,
    cvv,
    pin,
    expiary_date,
    usage,
    user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetCreditCardByCardNumber :one
SELECT * FROM "Finance"."Credit_Card"
WHERE card_number = $2 AND user_id = $1;

-- name: GetAllCreditCards :many
SELECT * FROM "Finance"."Credit_Card"
WHERE user_id = $1 ORDER BY bank_name, card_name;    

-- name: GetCreditCardByUsage :one
SELECT * FROM "Finance"."Credit_Card"
WHERE usage LIKE $2 AND user_id = $1;

-- name: UpdateCreditCardDetails :exec
UPDATE "Finance"."Credit_Card"
SET card_number = $3, cvv = $4, pin = $5, expiary_date = $6
WHERE card_name = $2 AND user_id = $1;

-- name: UpdateCreditCardUsage :exec
UPDATE "Finance"."Credit_Card"
SET usage = $3
WHERE card_number = $2 AND user_id = $1;

-- name: UpdateCreditCardPin :exec
UPDATE "Finance"."Credit_Card"
SET pin = $3
WHERE user_id = $1 AND card_number = $2;

-- name: DeleteCreditCard :exec
DELETE FROM "Finance"."Credit_Card"
WHERE card_number = $2 AND user_id = $1;

-- name: DeleteAllCreditCards :exec
DELETE FROM "Finance"."Credit_Card"
WHERE user_id = $1;