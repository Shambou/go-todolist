package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Shambou/todolist/internal/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// GetCompletedItems returns all completed items
func (h *Handler) GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get completed TodoItems")

	completedTodoItems, err := h.DB.GetItems(context.Background(), true)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error", err)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["completed_items"] = completedTodoItems

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}

// GetActiveItems returns all active items
func (h *Handler) GetActiveItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get active TodoItems")

	activeTodoItems, err := h.DB.GetItems(context.Background(), false)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error", err)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["active_items"] = activeTodoItems

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		log.Error(err)
		jsonResponse(w, http.StatusUnprocessableEntity, "Error parsing the form", "")
		return
	}

	log.WithFields(log.Fields{"description": r.FormValue("description")}).Info("Add new TodoItem. Saving to database.")

	item := models.TodoItem{Description: r.FormValue("description"), Completed: false}

	item, err := h.DB.InsertItem(context.Background(), item)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Internal server error, please contact admin", err)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["item"] = item

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := r.ParseForm(); err != nil {
		log.Error(err)
		jsonResponse(w, http.StatusUnprocessableEntity, "Error parsing the form", "")
		return
	}

	item, err := h.DB.GetItemById(context.Background(), id)
	if err != nil {
		jsonResponse(w, http.StatusNotFound, "Not found", "")
		return
	}

	data := make(map[string]interface{})
	data["completed"], _ = strconv.ParseBool(r.FormValue("completed"))

	updatedItem, err := h.DB.UpdateItem(context.Background(), item, data)
	if err != nil {
		log.Error(err)
		jsonResponse(w, http.StatusInternalServerError, "Internal server error, please contact admin", err)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["item"] = updatedItem

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}
