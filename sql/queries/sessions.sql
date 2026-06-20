-- name: GetSessions :many
select * from sessions;

-- name: GetSessionsByID :many
select *, timediff(sessions.end_datetime, sessions.start_datetime) as 'time_elapsed' from sessions
where sessions.tracker_id = ?;

-- name: GetOngoingSessions :many
select * from sessions
where sessions.end_datetime is null;

-- name: CreateSession :exec
insert into sessions (tracker_id, start_datetime)
values
(?, datetime('now', 'localtime'));