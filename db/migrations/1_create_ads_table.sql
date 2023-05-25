-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table ads
(
    id          uuid    not null
        constraint ads_pk primary key,
    title       varchar not null,
    description text    not null,
    price       float   not null,
    postedAt    timestamp    not null
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
