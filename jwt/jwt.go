package jwt

import (
	"github.com/gorilla/context"
	"log"
	"net/http"
	"strings"
	"time"
	"twiggg/packages/jwt"
)

// ------------------------------- JSON WEB TOKENS ------------------------------------------
// ------------------------------------------------------------------------------------------
/*
// ---- require a jwttoken in the header to be sent to access the page
func jwtRequiredMiddleWare(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		required := context.Get(r, "requiredformethod")
		if required == true {
			token := r.Header.Get("jwttoken")
			if len(token) == 0 {
				http.Error(w, "A jwttoken was expected in the header.", http.StatusExpectationFailed)
				return
			}
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}*/

func Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("jwttoken")
		if len(token) == 0 {
			http.Error(w, "JWT token required in header.", http.StatusBadRequest)
			log.Println("JWT verification: no token in the header.")
			return
		}
		claims, err := jwt.Validate(token, keys)
		if err != nil {
			http.Error(w, "JWT token is not valid.", http.StatusUnauthorized)
			log.Println("JWT verification:", token, ",", err)
			return
		}
		if claims.Exp < time.Now().Unix() {
			http.Error(w, "JWT token has expired", http.StatusUnauthorized)
			log.Println("JWT verification: expired token.", token)
			return
		}
		context.Set(r, "jwtclaims", claims)
		context.Set(r, "userid", claims.Sub.Username)
		context.Set(r, "role", claims.Sub.Role)
		next.ServeHTTP(w, r)
	})
}
