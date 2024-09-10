-- name: GetComment :one
SELECT * FROM comments
WHERE comment_id = $1 AND  user_id=$2 LIMIT 1;

-- name: CreateComment :exec
INSERT INTO  comments (user_id, post_id, message,comment_like) 
VALUES ($1, $2, $3, $4);

-- name: UpdateComment :one
UPDATE comments
  set 
    message= $1,
    comment_like = $2
WHERE comment_id = $3 AND user_id=$4
returning message,comment_like;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = $1;