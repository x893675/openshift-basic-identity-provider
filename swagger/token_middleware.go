package swagger

import (
	"fmt"
	"net/http"
	"time"

	"openshift-basic-identity-provider/db"
	"openshift-basic-identity-provider/helper"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	User string `json:"user"`
	Role int    `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(user *db.User) (string, error) {
	fmt.Println(user.Username)
	fmt.Println(user.Role)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			helper.ResponseWithJson(w, http.StatusUnauthorized,
				helper.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
		} else {
			token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					helper.ResponseWithJson(w, http.StatusUnauthorized,
						helper.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
					return nil, fmt.Errorf("not authorization")
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				helper.ResponseWithJson(w, http.StatusUnauthorized,
					helper.Response{Code: http.StatusUnauthorized, Msg: "not authorized"})
			} else {
				claims, _ := token.Claims.(*MyCustomClaims)
				user := claims.User
				role := claims.Role
				fmt.Println(user)
				fmt.Println(role)
				if role == 1 {
					r.Header.Add("role", "admin")
				} else {
					r.Header.Add("role", "common")
				}
				fmt.Println(r.Header)

				r.Header.Add("username", user)
				next.ServeHTTP(w, r)
			}
		}
	})
}
