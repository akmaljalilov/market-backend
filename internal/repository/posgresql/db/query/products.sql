-- name: ListMeasurement :many
SELECT *
FROM measurements;

-- name: InsertCategory :one
INSERT INTO categories (name, measurement_id, category_id)
VALUES ($1, $2, $3) RETURNING id;

-- name: InsertProduct :one
INSERT INTO products (name, category_id)
VALUES ($1, $2) RETURNING id;

-- name: CreatePurchase :one
INSERT INTO purchase_orders (dealer_id)
VALUES ($1) RETURNING id;

-- name: AddPurchaseItem :one
INSERT INTO purchase_items (purchase_order_id, product_id, quantity, price, status)
VALUES ($1, $2, $3, $3, $4) RETURNING id;

-- name: AddExpensesPurchaseItem :exec
UPDATE purchase_expenses SET sum=$1
WHERE purchase_item_id=$2;
