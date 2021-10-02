package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avalokia/tgtcx/backend/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func UploadCoupon(w http.ResponseWriter, r *http.Request) {

	// var p dictionary.Product

	// if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	// 	http.Error(w, "bad request", 400)
	// 	return
	// }
	// // p = product.AddProduct(context.Background(), p)
	// q, err := service.CreateProduct(p)
	// if err != nil {
	// 	// log.Fatal(err)
	// 	fmt.Fprintln(w, err.Error())
	// }
	// fmt.Fprintln(w, "success ", q.Name)
}

func GetCouponList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// p, err := product.GetProduct(context.Background(), idInt64)
	couponList, err := service.GetCouponList()
	if err != nil {
		// log.Fatal(err)
		fmt.Fprintln(w, err.Error())
	}

	parsedList, err := json.Marshal(couponList)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Fprintln(w, string(val))
	w.Write(parsedList)
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	// var p dictionary.Product

	// if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	// 	http.Error(w, "bad request", 400)
	// 	return
	// }

	// product.UpdateProduct(context.Background(), p)

	// fmt.Fprintf(w, "success")
}

func SetTargetUsers(w http.ResponseWriter, r *http.Request) {
	// var p dictionary.Product

	// if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
	// 	http.Error(w, "bad request", 400)
	// 	return
	// }

	// product.UpdateProduct(context.Background(), p)

	// fmt.Fprintf(w, "success")
}
