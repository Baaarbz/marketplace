-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO ads (id, title, description, price, postedat) VALUES ('90bc793b-3419-4b0a-9101-7ff4e7c12664', 'TV 45"" Sony', 'This is the description to test the app', 799.50, '2023-05-24 12:11:32.000000');
INSERT INTO ads (id, title, description, price, postedat) VALUES ('ef483c32-ca95-47c9-8643-c3f23706ee4c', 'Sportiva Rock Climbing Foot', 'This is the description to test the app', 799.50, '2023-05-24 12:11:32.000000');
INSERT INTO ads (id, title, description, price, postedat) VALUES ('c5fc62e5-5eea-40dd-a532-6262e0bec55a', 'Macbook Pro 16"', 'This is the description to test the app', 799.50, '2023-05-24 12:11:32.000000');
INSERT INTO ads (id, title, description, price, postedat) VALUES ('87dd12af-051e-4567-ac87-43361df0bf81', 'Rolex limited edition', 'This is the description to test the app', 799.50, '2023-05-24 12:11:32.000000');
INSERT INTO ads (id, title, description, price, postedat) VALUES ('2502bfa6-af82-427c-a8c6-e73d84f4d7a7', 'Northface t-shirt', 'This is the description to test the app', 799.50, '2023-05-24 12:11:32.000000');
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd