package middlew

import (
	"net/http"

	"github.com/gustavovargas511/sotho/db"
)

//DBCheck ...middleware que me permite conocer el estado de la base de datos
func DBCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ConnCheck() == 0 {
			http.Error(w, "Connection lost... :(", 500)
		}

		next.ServeHTTP(w, r)
	}

}
