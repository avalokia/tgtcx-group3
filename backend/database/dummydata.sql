INSERT INTO users (user_name,user_tier , user_location , top_category) VALUES
('a', 'Silver', 'jawa', 'Kosmetik'),
('b', 'Silver', 'jawa', 'Makanan'),
('c', 'Silver', 'luarjawa', 'None'),
('c2', 'Silver', 'luarjawa', 'None'),
('d', 'Silver', 'luarjawa', 'Tagihan'),
('e', 'Gold', 'jawa', 'Kosmetik'),
('f', 'Gold', 'jawa', 'Makanan'),
('g', 'Gold', 'luarjawa', 'Tagihan'),
('h', 'Gold', 'luarjawa', 'Tagihan'),
('i', 'Platinum', 'jawa', 'Kosmetik'),
('j', 'Platinum', 'jawa', 'Makanan'),
('k', 'Platinum', 'luarjawa', 'Tagihan'),
('l', 'Platinum', 'luarjawa', 'Makanan'),
('m', 'Diamond', 'jawa', 'Kosmetik'),
('n', 'Diamond', 'jawa', 'Makanan'),
('o', 'Diamond', 'luarjawa', 'Tagihan'),
('p', 'Diamond', 'luarjawa', 'Makanan')


INSERT INTO coupons (coupon_name,coupon_category,start_date,end_date) VALUES
('coupon 50', 'ONGKIR', '2021-10-03', '2021-10-10'),
('coupon 70', 'ONGKIR', '2021-10-03',  '2021-10-10'),
('coupon 100', 'ONGKIR', '2021-10-03',  '2021-10-10'),
('coupon spesial', 'Special', '2021-10-03',  '2021-10-10'),
('coupon 50', 'POTONGAN', '2021-10-03',  '2021-10-10'),
('coupon 70', 'POTONGAN', '2021-10-03',  '2021-10-10'),
('coupon 100', 'POTONGAN', '2021-10-03',  '2021-10-10'),
('coupon 50', 'ONGKIR', '2021-10-10', '2021-10-17'),
('coupon 70', 'ONGKIR', '2021-10-10',  '2021-10-17'),
('coupon 100', 'ONGKIR', '2021-10-10',  '2021-10-17'),
('coupon spesial', 'Special', '2021-10-10',  '2021-10-17'),
('coupon 50', 'POTONGAN', '2021-10-10',  '2021-10-17'),
('coupon 70', 'POTONGAN', '2021-10-10',  '2021-10-17'),
('coupon 100', 'POTONGAN', '2021-10-10',  '2021-10-17'),
('coupon', 'POTONGAN dan ongkir', '2021-10-10',  '2021-10-17')




insert into user_coupons (user_id,coupon_id)
select user_id,case
when user_tier = 'Silver' and user_location = 'luarjawa' and top_category !='None' then 3
when user_tier = 'Silver' and user_location = 'jawa' and top_category !='None' then 7
when user_tier = 'Gold' and user_location = 'luarjawa' and top_category !='None' then 1
when user_tier = 'Gold' and user_location = 'jawa' and top_category !='None' then 5
when user_tier = 'Platinum' and user_location = 'luarjawa' and top_category !='None' then 2
when user_tier = 'Platinum' and user_location = 'jawa' and top_category !='None' then 6
when user_tier = 'Diamond' and user_location = 'luarjawa' and top_category !='None' then 3
when user_tier = 'Diamond' and user_location = 'jawa' and top_category !='None' then 7
when top_category ='None' then 15
end as coupond_id
from users


