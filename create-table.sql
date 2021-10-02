CREATE TABLE IF NOT EXISTS 
users (
	user_id SERIAL PRIMARY KEY, 
	user_name TEXT, 
	user_tier TEXT, 
	user_location TEXT, 
	top_category TEXT
);

CREATE TABLE IF NOT EXISTS 
coupons (
	coupon_id SERIAL PRIMARY KEY, 
	coupon_name TEXT,
	coupon_category TEXT,
	start_date TIMESTAMP,
	end_date TIMESTAMP,
	status TEXT
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