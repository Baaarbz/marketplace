-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table ad_posted_sites
(
    ad_id       uuid
        constraint ad_id
            references ads,
    posted_site varchar,
    constraint ad_posted_sites_pk
        primary key (ad_id, posted_site)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
