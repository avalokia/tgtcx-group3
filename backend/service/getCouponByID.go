package service

import (
	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func GetCouponByID(paramID int) (dictionary.Coupons, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	SELECT 
		coupon_id, 
		coupon_name, 
		coupon_category, 
		start_date, 
		end_date,
		status
	FROM 
		coupons
	WHERE coupon_id = $1
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
