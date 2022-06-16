package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"ks/db"
	"ks/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("fjslajsaflhasughaslihaslhdasldhasliuhasjbfbvisflh")

func Basic(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./tmpl/auth_form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	fmt.Println(r.Body, r.ParseForm())
	u, p, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(401)

		return
	}
	if db.FindUser(u, p) {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(401)
	}

}

func JWT(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./tmpl/auth_form.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	var creds model.User
	// Get the JSON body and decode into credentials
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		fmt.Println(err)
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, err := db.GetPwd(creds.Name)
	userid, err := db.GetId(creds.Name)

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if err != nil || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &model.Claims{
		Id:       userid,
		Username: creds.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func ValidateToken(r *http.Request) (*model.Claims, bool) {
	claims := &model.Claims{}
	c, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		return claims, false
	}
	userToken := c.Value

	tkn, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, false
		}

		return claims, false
	}
	return claims, tkn.Valid
}
