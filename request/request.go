package request

import (
	"github.com/gorilla/context"
	"net/http"
)

//these middleware restrict the application of next middleware, to specific request methods
func GET(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			context.Set(r, "requiredformethod", true)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func POST(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			context.Set(r, "requiredformethod", true)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func PUT(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			context.Set(r, "requiredformethod", true)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func DELETE(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			context.Set(r, "requiredformethod", true)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func Clear(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "requiredformethod", false)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func HEAD(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "HEAD" {
			context.Set(r, "requiredformethod", true)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func All(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "requiredformethod", true)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func Forbidden(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		required := context.Get(r, "requiredformethod")
		if required == true {
			http.Error(w, "Sorry, you have sent a forbidden request...", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
