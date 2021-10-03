CREATE TABLE IF NOT EXISTS 
users (
	user_id SERIAL PRIMARY KEY, 
	user_name varchar(20), 
	user_tier char(1), 
	user_location boolean, 
	top_category text
);

CREATE TABLE IF NOT EXISTS 
coupons (
	coupon_id SERIAL PRIMARY KEY, 
	coupon_name TEXT,
	coupon_category C,
	start_date TIMESTAMP,
	end_date TIMESTAMP,
	);

CREATE TABLE IF NOT EXISTS 
user_coupons (
	relation_id SERIAL PRIMARY KEY, 
	user_id INT,
	coupon_id INT,
		FOREIGN KEY(user_id)
			REFERENCES users(user_id),
		FOREIGN KEY(coupon_id)
			REFERENCES coupons(coupon_id)
	);
