package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avalokia/tgtcx/backend/dictionary"
	"github.com/avalokia/tgtcx/backend/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func UploadCoupon(w http.ResponseWriter, r *http.Request) {

	var c dictionary.Coupons

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	result, err := service.UploadCoupon(c)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, "success adding", result.Name)
}

func GetCouponList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	couponList, err := service.GetCouponList()
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	parsedList, err := json.Marshal(couponList)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(parsedList)
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	var c dictionary.Coupons

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	result, err := service.UpdateCoupon(c)

	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, "success updating", result.ID, ": ", result.Name)
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
