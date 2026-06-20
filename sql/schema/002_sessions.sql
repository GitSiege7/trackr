-- +goose Up
create table sessions(
    id integer primary key,
    tracker_id integer not null,
    start_datetime text not null,
    end_datetime text,
    note text,
    foreign key (tracker_id) references trackers(id)
);

-- +goose Down
drop table sessions;