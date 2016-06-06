package auth

import (
	"net/http"
	"strings"
	"time"
)

// ---- require BasicAuth with specific credentials to access the page
func Basic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Not Authorized to access this page without credentials.", http.StatusUnauthorized)
			return
		}

		type expectedAuth struct {
			user string
			pass string
		}

		valid := expectedAuth{}
		y, m, d := time.Now().Date()
		h, min, _ := time.Now().Clock()
		portion := y + int(m) + d + h + int(min/15)

		valid.user = "twigggLaahs"
		valid.pass = strings.Join([]string{"1K9l8985", strconv.Itoa(portion)}, "")

		if !(username == valid.user && password == valid.pass) {
			http.Error(w, "Wrong credentials provided... Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// ---- require an API key in the header to be sent to access the page, and the key to be valid
func APIkeyRequ(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		required := context.Get(r, "requiredformethod")
		if required == true {
			key := r.Header.Get("APIkey")
			if len(key) == 0 {
				http.Error(w, "An API key was expected in the header.", http.StatusExpectationFailed)
				return
			}
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
