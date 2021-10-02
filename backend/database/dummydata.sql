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
('coupon 100', 'POTONGAN', '2021-10-10',  '2021-10-17')



select *,case
when user_tier = 'Silver' and user_location = 'luarjawa' and top_category !='None' then 'Coupon 100 Ongkir'
when user_tier = 'Silver' and user_location = 'jawa' and top_category !='None' then 'Coupon 100 Potongan'
when user_tier = 'Gold' and user_location = 'luarjawa' and top_category !='None' then 'Coupon 50 Ongkir'
when user_tier = 'Gold' and user_location = 'jawa' and top_category !='None' then 'Coupon 50 Potongan'
when user_tier = 'Platinum' and user_location = 'luarjawa' and top_category !='None' then 'Coupon 70 Ongkir + Special'
when user_tier = 'Platinum' and user_location = 'jawa' and top_category !='None' then 'Coupon 70 Potongan + Special'
when user_tier = 'Diamond' and user_location = 'luarjawa' and top_category !='None' then 'Coupon 100 Ongkir + Special'
when user_tier = 'Diamond' and user_location = 'jawa' and top_category !='None' then 'Coupon 100 Potongan + Special'
when top_category ='None' then 'Coupon'
end as kupon
from users
