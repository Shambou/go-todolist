package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Shambou/todolist/internal/models"
	log "github.com/sirupsen/logrus"
)

// GetCompletedItems returns all completed items
func (h *Handler) GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get completed TodoItems")

	completedTodoItems, err := h.DB.GetTodoItems(context.Background(), true)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error", err)
	}

	dataMap := make(map[string]interface{})
	dataMap["completed_items"] = completedTodoItems

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}

// GetActiveItems returns all active items
func (h *Handler) GetActiveItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("Get active TodoItems")

	activeTodoItems, err := h.DB.GetTodoItems(context.Background(), false)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error", err)
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
	}

	log.WithFields(log.Fields{"description": r.FormValue("description")}).Info("Add new TodoItem. Saving to database.")

	item := models.TodoItem{Description: r.FormValue("description"), Completed: false}

	item, err := h.DB.InsertItem(context.Background(), item)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Internal server error, please contact admin", err)
	}

	dataMap := make(map[string]interface{})
	dataMap["item"] = item

	jsonResponse(w, http.StatusOK, "OK", dataMap)
}

func jsonResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	if err := json.NewEncoder(w).Encode(Response{
		Status:  status,
		Message: message,
		Data:    data,
	}); err != nil {
		panic(err)
	}
}
