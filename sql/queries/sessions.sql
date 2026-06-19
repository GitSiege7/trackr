-- name: GetSessions :many
select * from sessions;

-- name: GetSessionsByID :many
select * from sessions
where sessions.tracker_id = ?;