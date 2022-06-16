package api

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	claims, ok := ValidateToken(r)
	fmt.Println(claims, ok)

}
