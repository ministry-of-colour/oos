-- create database oos;

drop table if exists brands;
create table brands (
    ID char(12) primary key,
    name text not null,
    descr text not null,
    icon text not null
);

drop table if exists brandStockStatus;
create table brandStockStatus (
    brandID char(12) not null,
    date timestamp,
    descr text
);
create index brandStockIdx on brandStockStatus (brandID, date);

drop table if exists orders;
create table orders (
    ID char(36) primary key,
    brand char(12) not null,
    email text not null,
    name text not null,
    address text not null,
    suburb text not null,
    state text not null,
    postcode text not null
);
create unique index brandOrdersIdx on orders (brand, ID);

drop table if exists categories;
create table categories (
    ID char(36) not null primary key,
    name text not null,
    descr text not null
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
    categoryID char(36) not null,
    brand char(6) not null,
    name text not null,
    image text,
    descr text
);
create unique index brandItemsIdx on items (brand, ID);



