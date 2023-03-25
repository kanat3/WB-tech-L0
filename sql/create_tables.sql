create table if not exists orders (
    order_id varchar(10) not null unique,
    data json not null
);

create or replace procedure push( order_id varchar(10), data json)
language plpgsql 
as $$
declare
begin
    insert into orders values( order_id, data);
end;
$$;