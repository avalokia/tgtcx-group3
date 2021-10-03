package service

import (
	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func GetCouponByUser(paramID int) (dictionary.Coupons, error) {

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
		c.end_date,
		c.status
	FROM public.user_coupons AS uc
	INNER JOIN public.coupons c
		on uc.coupon_id = c.coupon_id
	INNER JOIN public.users u
		on uc.user_id = u.user_id
	`
	// actual query process
	row := db.QueryRow(query, paramID)

	var data dictionary.Coupons
	err := row.Scan(
		&data.ID,
		&data.Name,
		&data.Category,
		&data.StartDate,
		&data.EndDate,
		&data.Status,
	)
	if err != nil {
		return data, err
	}

	return data, nil
}
