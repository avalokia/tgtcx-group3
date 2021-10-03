INSERT INTO users (user_name,user_tier , isjawa , top_category) VALUES
('a', '1', True, 'Kosmetik'),
('b', '1', False, 'Makanan'),
('c', '1', FALSE, 'None'),
('c2', '1', FALSE, 'None'),
('d', '1', FALSE, 'Tagihan'),
('e', '2', TRUE, 'Kosmetik'),
('f', '2', TRUE, 'Makanan'),
('g', '2', FALSE, 'Tagihan'),
('h', '2', FALSE, 'Tagihan'),
('i', '3', TRUE, 'Kosmetik'),
('j', '3', TRUE, 'Makanan'),
('k', '3', FALSE, 'Tagihan'),
('l', '3', FALSE, 'Makanan'),
('m', '4', TRUE, 'Kosmetik'),
('n', '4', TRUE, 'Makanan'),
('o', '4', FALSE, 'Tagihan'),
('p', '4', FALSE, 'Makanan')


INSERT INTO coupons (coupon_name,coupon_category,coupon_potongan,start_date,end_date) VALUES
('coupon 50', 'ONGKIR',50, '2021-10-03', '2021-10-10'),
('coupon 70', 'ONGKIR',70, '2021-10-03',  '2021-10-10'),
('coupon 100', 'ONGKIR',100, '2021-10-03',  '2021-10-10'),
('coupon special', 'Special',100,'2021-10-03',  '2021-10-10'),
('coupon 50', 'POTONGAN',50, '2021-10-03',  '2021-10-10'),
('coupon 70', 'POTONGAN',70, '2021-10-03',  '2021-10-10'),
('coupon 100', 'POTONGAN',100, '2021-10-03',  '2021-10-10'),
('coupon 50', 'ONGKIR',50, '2021-10-10', '2021-10-17'),
('coupon 70', 'ONGKIR',70, '2021-10-10',  '2021-10-17'),
('coupon 100', 'ONGKIR',100, '2021-10-10',  '2021-10-17'),
('coupon special', 'Special',100, '2021-10-10',  '2021-10-17'),
('coupon 50', 'POTONGAN',50, '2021-10-10',  '2021-10-17'),
('coupon 70', 'POTONGAN',70, '2021-10-10',  '2021-10-17'),
('coupon 100', 'POTONGAN',100, '2021-10-10',  '2021-10-17'),
('coupon', 'POTONGAN dan ongkir',100, '2021-10-10',  '2021-10-17')




insert into user_coupons (user_id,coupon_id)
select user_id,case
when user_tier = '1' and isjawa = false and top_category !='None' then 3
when user_tier = '1' and isjawa = TRUE and top_category !='None' then 7
when user_tier = '2' and isjawa = false and top_category !='None' then 1
when user_tier = '2' and isjawa = TRUE and top_category !='None' then 5
when user_tier = '3' and isjawa = TRUE and top_category !='None' then 2
when user_tier = '3' and isjawa = TRUE and top_category !='None' then 6
when user_tier = '4' and isjawa = false and top_category !='None' then 3
when user_tier = '4' and isjawa = TRUE and top_category !='None' then 7
when top_category ='None' then 15
end as coupond_id
from users

insert into user_coupons (user_id,coupon_id)
select user_id, case
when user_tier = '3' or user_tier = '4'   then 11
end as coupon_id
from users

