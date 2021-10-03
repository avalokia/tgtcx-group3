package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "CouponListGetter",
			Description: "All query related to getting coupon data",
			Fields: graphql.Fields{
				"GetCouponByID": &graphql.Field{
					Type:        CouponType,
					Description: "Get Coupon List by Coupon ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetCouponByID(),
				},
				"GetCouponByUser": &graphql.Field{
					Type:        UserCouponType,
					Description: "Get Coupon List for Certain User",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.productResolver.GetCouponByUser(),
				},
				"GetCouponList": &graphql.Field{
					Type:        graphql.NewList(CouponType),
					Description: "Get products",
					Resolve:     s.productResolver.GetCouponList(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "CouponSetter",
			Description: "All query related to modify coupon data",
			Fields: graphql.Fields{
				"UploadCoupon": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Create coupon",
					Args: graphql.FieldConfigArgument{
						"coupon_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"coupon_category": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"potongan": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"end_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.productResolver.UploadCoupon(),
				},
				// "UpdateProduct": &graphql.Field{
				// 	Type:        graphql.Boolean,
				// 	Description: "Update product by ID",
				// 	Args: graphql.FieldConfigArgument{
				// 		"product_id": &graphql.ArgumentConfig{
				// 			Type: graphql.NewNonNull(graphql.Int),
				// 		},
				// 		"product_name": &graphql.ArgumentConfig{
				// 			Type: graphql.NewNonNull(graphql.String),
				// 		},
				// 		"product_shop_name": &graphql.ArgumentConfig{
				// 			Type: graphql.String,
				// 		},
				// 		"product_price": &graphql.ArgumentConfig{
				// 			Type: graphql.Float,
				// 		},
				// 		"product_image": &graphql.ArgumentConfig{
				// 			Type: graphql.String,
				// 		},
				// 	},
				// 	Resolve: s.productResolver.UpdateProducts(),
				// },
			},
		}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
