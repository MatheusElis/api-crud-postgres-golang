package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MatheusElis/api-crud-postgres-golang/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  var todo models.Todo

  err = json.NewDecoder(r.Body).Decode(&todo)
  if err != nil {
    log.Printf("Erro ao fazer decode do json: %v", err)
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  rows, err := models.Update(int64(id), todo)
  if err != nil {
    log.Printf("Error ao atualizar registro: %v", err)
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  if rows > 1 {
    log.Printf("Error: foram atualizados %d registros", rows)
  }

  resp := map[string]any{
    "Error": false,
    "Message": "dados atualizados com sucesso!",
  }

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
