-- +goose Up
create table trackers (
    id integer primary key,
    name text not null,
    unique(name)
);

-- +goose Down
drop table trackers;