package service

import (
	"errors"

	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func UploadCoupon(data dictionary.Coupons) (dictionary.Coupons, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	INSERT INTO coupons 
	(
		coupon_name, 
		coupon_category, 
		start_date, 
		end_date,
		status) 
		VALUES (
			$1, 
			$2, 
			$3, 
			$4,
			$5)
	`
	// actual query process
	result, err := db.Exec(query,
		data.Name,
		data.Category,
		data.StartDate,
		data.EndDate,
		data.Status,
	)
	if err != nil {
		return data, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return data, err
	}

	if affected == 0 {
		return data, errors.New("no row created")
	}

	return data, err
}
