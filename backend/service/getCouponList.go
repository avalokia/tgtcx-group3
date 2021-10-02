package service

import (
	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func GetCouponList() ([]dictionary.Coupons, error) {

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
	`
	// actual query process
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dictionary.Coupons
	for rows.Next() {
		var data dictionary.Coupons

		rows.Scan(
			&data.ID,
			&data.Name,
			&data.Category,
			&data.StartDate,
			&data.EndDate,
			&data.Status,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}
