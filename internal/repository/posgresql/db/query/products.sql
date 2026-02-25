-- name: ListMeasurement :many
SELECT *
FROM measurements;

-- name: InsertCategory :one
INSERT INTO categories (name, measurement_id)
VALUES ($1, $2) RETURNING id;

-- name: InsertProduct :one
INSERT INTO products (name, category_id)
VALUES ($1, $2) RETURNING id;

-- name: CreatePurchase :one
INSERT INTO purchase_orders (dealer_id)
VALUES ($1) RETURNING id;

-- name: AddConsumptionPurchase :exec
UPDATE purchase_orders SET consumption=$1
WHERE id=$2;

-- name: AddPurchaseItem :one
INSERT INTO purchase_items (purchase_order_id, product_id, quantity, price, status)
VALUES ($1, $2, $3, $3, $4) RETURNING id;