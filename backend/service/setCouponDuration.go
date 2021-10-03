package service

import (
	"errors"

	"github.com/avalokia/tgtcx/backend/database"
	"github.com/avalokia/tgtcx/backend/dictionary"
)

func SetCouponDuration(data dictionary.Coupons) (dictionary.Coupons, error) {

	// you can connect and
	// get current database connection
	db := database.GetDB()

	// construct query
	query := `
	UPDATE
		coupons
	SET
		start_date = $2,
		end_date = $3
	WHERE
		coupon_id = $1
	`

	// if data.StartDate < now && data.EndDate > now { status = active, can be added? }

	// actual query process
	result, err := db.Exec(query,
		data.ID,
		data.StartDate,
		data.EndDate,
		// data.Status,
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
