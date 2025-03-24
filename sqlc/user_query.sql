-- name: GetAllUsers :many
Select fname, lname, email, phone_number, is_admin from Users;

-- name: GetUserByClerkID :one
Select fname, lname, email, phone_number, is_admin, bible_id from Users
where clerk_id = $1 limit 1;
