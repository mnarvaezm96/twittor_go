package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDusuario
	t.UsurioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := db.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false

	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
