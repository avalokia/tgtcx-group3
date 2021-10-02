package gqlserver

import (
	"github.com/avalokia/tgtcx/backend/dictionary"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetCouponList() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		// id, _ := p.Args["product_id"].(int)

		// update to use Usecase from previous session
		return dictionary.Coupons{}, nil
	}
}
