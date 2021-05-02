-- name: CreateInvitee :one
INSERT INTO invitees (
    full_name,
    inviter,
    stream_id,
    email,
    mobile_number
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ListInvitees :many
SELECT *
FROM invitees
WHERE inviter = $1
  AND stream_id = $2
ORDER BY id
LIMIT $3 OFFSET $4;