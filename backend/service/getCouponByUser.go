package service

import (
	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func GetCouponByUser(paramID int) (dictionary.UserCoupons, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT 
		uc.user_id,
		u.user_name,
		uc.coupon_id,
		c.coupon_name,
		c.coupon_category,
		c.start_date,
		c.end_date
	FROM public.user_coupons AS uc
	INNER JOIN public.coupons c
		on uc.coupon_id = c.coupon_id
	INNER JOIN public.users u
		on uc.user_id = u.user_id
	WHERE uc.user_id = $1
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	var data dictionary.UserCoupons
	err := row.Scan(
		&data.UserID,
		&data.UserName,
		&data.CouponID,
		&data.CouponName,
		&data.CouponCategory,
		&data.StartDate,
		&data.EndDate,
	)
	if err != nil {
		return data, err
	}

	return data, nil
}
