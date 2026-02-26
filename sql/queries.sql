-- name: CreateClient :one
INSERT INTO clients (name, phone)
VALUES ($1, $2)
RETURNING *;


-- name: GetClientByID :one
SELECT * FROM clients
WHERE id = $1;


-- name: ListClients :many
SELECT * FROM clients
ORDER BY created_at DESC;




-- name: CreateLoan :one
INSERT INTO loans (
    client_id,
    principal,
    monthly_interest_rate,
    months,
    total_amount,
    disbursed_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)

RETURNING *;

-- name: GetLoanByID :one
SELECT * FROM loans
WHERE id = $1;

-- name: ListLoansByClient :many
SELECT * FROM loans
WHERE client_id = $1
ORDER BY disbursed_at DESC;

-- name: ListAllLoans :many
SELECT * FROM loans
ORDER BY created_at DESC;




-- name: CreatePayment :one
INSERT INTO payments (
    loan_id,
    amount
) VALUES (
    $1, $2
)
RETURNING *;


-- name: ListPaymentsByLoan :many
SELECT * FROM payments
WHERE loan_id = $1
ORDER BY paid_at ASC;



-- name: GetLoanBalance :one
SELECT
    l.id,
    l.total_amount,
    COALESCE(SUM(p.amount), 0) AS total_paid,
    l.total_amount - COALESCE(SUM(p.amount), 0) AS remaining_balance
FROM loans l
LEFT JOIN payments p ON l.id = p.loan_id
WHERE l.id = $1
GROUP BY l.id;

-- name: GetTotalProfit :one
SELECT
    COALESCE(SUM(p.amount), 0) - COALESCE(SUM(l.principal), 0) AS total_profit
FROM loans l
LEFT JOIN payments p ON l.id = p.loan_id;