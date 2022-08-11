-- name: GetAccount :one
SELECT * FROM accounts
WHERE username= $1 
AND password= $2 LIMIT 1;

-- name: GetAccountByUsername :one 
SELECT * FROM accounts
WHERE username= $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (
    firstname, lastname, emailaddress, username, password
)
VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;
