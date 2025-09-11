-- name: AddCreditCard :one
INSERT INTO "Finance"."Credit_Card" (
    bank_name,
    card_name,
    card_number,
    cvv,
    pin,
    usage,
    user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetCreditCardsByUserId :many
SELECT * FROM "Finance"."Credit_Card"
WHERE user_id = $1;

-- name: DeleteCreditCardByCardNumber :exec
DELETE FROM "Finance"."Credit_Card"
WHERE card_number = $1 AND user_id = $2;

-- name: UpdateCreditCardUsageByCardNumber :exec
UPDATE "Finance"."Credit_Card"
SET usage = $1
WHERE card_number = $2 AND user_id = $3;

-- name: GetCreditCardByCardNumber :one
SELECT * FROM "Finance"."Credit_Card"
WHERE card_number = $1 AND user_id = $2;

-- name: GetAllCreditCards :many
SELECT * FROM "Finance"."Credit_Card"
WHERE user_id = $1 ORDER BY bank_name, card_name;    

-- name: GetCreditCardByUsage :one
SELECT * FROM "Finance"."Credit_Card"
WHERE usage LIKE $1 AND user_id = $2;