-- create database oos;

drop table if exists orders;
create table orders (
    ID char(36) primary key,
    email text not null,
    name text not null,
    address text not null,
    suburb text not null,
    state text not null,
    postcode text not null
);

drop table if exists orderItems;
create table orderItems (
    orderID char(36) not null,
    itemID char(36) not null,
    qty int not null,
    unitPrice int not null,
    total int not null
);
create unique index orderItemsIdx on orderItems (orderID, itemID); 

drop table if exists items;
create table items (
    ID char(36) primary key,
    name text,
    image text,
    descr text
);


