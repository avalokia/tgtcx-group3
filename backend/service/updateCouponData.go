package service

import (
	"errors"

	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func UpdateCoupon(data dictionary.Coupons) (dictionary.Coupons, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE
		coupons
	SET
		coupon_name = $2,
		coupon_category = $3,
		start_date = $4,
		end_date = $5
	WHERE
		coupon_id = $1
	`
	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.Name,
		data.Category,
		data.StartDate,
		data.EndDate,
	)
	if err != nil {
		return data, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return data, err
	}

	if affected == 0 {
		return data, errors.New("no row updated")
	}

	return data, nil
}
