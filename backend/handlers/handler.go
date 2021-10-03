package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/avalokia/tgtcx/backend/dictionary"
	"github.com/avalokia/tgtcx/backend/service"
)

var layout = "2021-10-01T00:00:00Z"

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

func GetCouponByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idstring := r.URL.Query().Get("id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// p, err := product.GetProduct(context.Background(), idInt64)
	p, err := service.GetCouponByID(int(idInt64))
	if err != nil {
		// log.Fatal(err)
		fmt.Fprintln(w, err.Error())
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(val)
}

func GetCouponByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idstring := r.URL.Query().Get("user_id")

	idInt64, err := strconv.ParseInt(idstring, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// p, err := product.GetProduct(context.Background(), idInt64)
	p, err := service.GetCouponByUser(int(idInt64))
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	val, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(val)
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

func SetCouponDuration(w http.ResponseWriter, r *http.Request) {
	var c dictionary.Coupons

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	//checking if the coupon is live or not
	data, err := service.GetCouponByID(int(c.ID))
	if err != nil {
		// log.Fatal(err)
		fmt.Fprintln(w, err.Error())
	}

	start, err := time.Parse(layout, data.StartDate)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	end, err := time.Parse(layout, data.EndDate)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	if start.Unix() < time.Now().Unix() && end.Unix() > time.Now().Unix() {
		fmt.Fprintln(w, "The coupon is live and can't be updated!")
	} else { //else, set the coupon duration
		result, err := service.SetCouponDuration(c)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		fmt.Fprintln(w, "success setting coupon duration", result.Name)
	}

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
