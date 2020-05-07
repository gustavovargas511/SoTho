package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavovargas511/sotho/db"
	"github.com/gustavovargas511/sotho/models"
)

//Record es la funcion para crear en la database un nuevo user
func Record(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data error"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is MANDATORY", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password cannot be less than 6 characters", 400)
	}

	_, encontrado, _ := db.UserExists(t.Email)

	if encontrado == true {
		http.Error(w, "Email already exists. Do a double check, please.", 400)
		return
	}

	_, status, err := db.InsertUser(t)

	if err != nil {
		http.Error(w, "Error tryin' to insert user, please doublecheck"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed registration, please doublecheck", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
