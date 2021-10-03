package gqlserver

import (
	"github.com/avalokia/tgtcx/backend/dictionary"
	"github.com/avalokia/tgtcx/backend/service"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetCouponByID() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)

		// update to use Usecase from previous session
		return service.GetCouponByID(id)
	}
}

func (r *Resolver) GetCouponByUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)

		// update to use Usecase from previous session
		return service.GetCouponByUser(id)
	}
}

func (r *Resolver) GetCouponList() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {

		// update to use Usecase from previous session
		return service.GetCouponList()
	}
}

func (r *Resolver) UploadCoupon() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["coupon_name"].(string)
		category, _ := p.Args["coupon_category"].(string)
		potongan, _ := p.Args["potongan"].(int64)
		startDate, _ := p.Args["start_date"].(string)
		endDate, _ := p.Args["end_date"].(string)

		req := dictionary.Coupons{
			Name:      name,
			Category:  category,
			Potongan:  potongan,
			StartDate: startDate,
			EndDate:   endDate,
		}

		_, err := service.UploadCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}
