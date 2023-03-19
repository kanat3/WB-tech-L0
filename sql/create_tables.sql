create table if not exists orders (
    id varchar(21) not null unique,
    body json not null
);