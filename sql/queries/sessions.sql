-- name: GetSessions :many
select * from sessions;

-- name: GetSessionsByTracker :many
select *, timediff(sessions.end_datetime, sessions.start_datetime) as 'time_elapsed' from sessions
where sessions.tracker_id = ?;

-- name: GetOngoingSessions :many
select * from sessions
where sessions.end_datetime is null;

-- name: CreateSession :one
insert into sessions (tracker_id, start_datetime)
values
(?, datetime('now', 'localtime'))
returning *;

-- name: GetSession :one
select *, timediff(sessions.end_datetime, sessions.start_datetime) as 'time_elapsed' from sessions
where sessions.id = ?;

-- name: DeleteSession :exec
delete from sessions
where sessions.id = ?;

-- name: EndSession :one
update sessions
set end_datetime = datetime('now', 'localtime')
where sessions.id = ?
returning *;

-- name: UpdateNote :one
update sessions
set note = ?
where sessions.id = ?
returning *;