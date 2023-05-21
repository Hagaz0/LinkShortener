CREATE TABLE links (
    id serial not null unique,
    original_link varchar(255) not null unique,
    short_link varchar(255) not null unique
);