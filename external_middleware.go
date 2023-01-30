package external_middleware

import (
	"log"
	"net/http"
	"time"
)

func ExternalAuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			log.Println(`r.Cookie("token") Exception: ` + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		email := c.Value

		if email != "abc123@example.com" {
			log.Println(`email != "abc123@example.com" Exception: ` + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    email,
			Expires:  time.Now().Add(time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteDefaultMode,
			Secure:   true,
		})

		next.ServeHTTP(w, r)
	})
}
