-- name: ListMeasurement :many
SELECT *
FROM measurements;

-- name: InsertCategory :one
WITH existing AS (SELECT ctg.id
                  FROM categories as ctg
                  WHERE ctg.name = $1),
     inserted AS (
INSERT
INTO categories ("name", measurement_id, category_id)
SELECT $1,
       $2,
       $3 WHERE NOT EXISTS (SELECT 1 FROM existing)
    RETURNING id
)
SELECT id
FROM inserted
UNION ALL
SELECT id
FROM existing LIMIT 1;

-- name: InsertProduct :one
WITH existing AS (SELECT p.id, p.category_id
                  FROM products as p
                  WHERE p.name = $1),
     inserted AS (
INSERT
INTO products (name, category_id)
SELECT $1,
       $2 WHERE NOT EXISTS (SELECT 1 FROM existing)
    RETURNING id
)
SELECT id
FROM inserted
UNION ALL
SELECT id
FROM existing LIMIT 1;

-- name: CreatePurchase :one
INSERT INTO purchase_orders (dealer_id)
VALUES ($1) RETURNING id;

-- name: AddPurchaseItem :one
INSERT INTO purchase_items (purchase_order_id, product_id, quantity, price, status)
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: AddExpensesPurchaseItem :exec
INSERT INTO purchase_expenses (sum, purchase_item_id, data)
VALUES ($1, $2, $3);

-- name: GetProductsBalance :many
SELECT * FROM products_balance;
