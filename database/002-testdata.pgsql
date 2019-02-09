insert into categories values ('books', 'books and things', 'all the books');
insert into categories values ('merch', 'merchandise', 'tshirts and mugs and stuff like that');

insert into items values ( 'abc', 'books','oos', 'the abc item');
insert into items values ( 'def', 'books','oos', 'the def item');
insert into items values ( 'ghi', 'books','oos', 'the ghi item');
insert into items values ( 'jkl', 'books','oos', 'the jkl item');
insert into items values ( 'mno', 'merch','oos', 'the mno item');
insert into items values ( 'prq', 'merch','oos', 'the prq item');
insert into items values ( 'stu', 'merch','oos', 'the stu item');
insert into items values ( 'vwy', 'merch','oos', 'the vwy item');
insert into items values ( 'zzz', 'merch','oos', 'the zzz item');

insert into items values ( 'first', 'books','stuff', 'the first item');
insert into items values ( 'second', 'books','stuff', 'the second item');
insert into items values ( 'third', 'books','stuff', 'the third item');
insert into items values ( 'fourth', 'books','stuff', 'the fourth item');
insert into items values ( 'fifth', 'merch','stuff', 'the fifth item');
insert into items values ( 'sixth', 'merch','stuff', 'the sixth item');
insert into items values ( 'seventh', 'merch','stuff', 'the seventh item');
insert into items values ( 'eighth', 'merch','stuff', 'the eighth item');
insert into items values ( 'ninth', 'merch','stuff', 'the ninth item');

insert into brands values ('oos', 'The Old Owls Scarf', 'Childrens book','');
insert into brands values ('stuff', 'Stuff the magic dragon', 'Childrens book','');

insert into brandStockStatus values ('oos', now(), 'Few left in stock, next edition ready early April 2019');
insert into brandStockStatus values ('stuff', now(), 'Sold out, next batch due at the start of March 2019');