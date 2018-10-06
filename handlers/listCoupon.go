package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dmps/PerkBoxTest/storage"
)

func List(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitParam := r.URL.Query().Get("limit")
		var limit uint64 = 0
		if limitParam != "" {
			var err error
			limit, err = strconv.ParseUint(limitParam, 10, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("bad limit passed to API: %s", err), http.StatusInternalServerError)
				return
			}
		}
		coupons, err := db.List()
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		var result []byte
		for _, coupon := range coupons {
			if limit < 10 {
				fmt.Println(limit, coupon)
				result = append(result, coupon...)
				limit++
			} else {
				break
			}

		}
		fmt.Println(result, limit)
		w.Write(result)
	})
}
