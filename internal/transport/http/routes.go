package http

// MapRoutes maps the routes to the handlers
func (h *Handler) MapRoutes() {
	h.Router.HandleFunc("/healthz", h.Healthz).Methods("GET")
	h.Router.HandleFunc("/items/completed", h.GetCompletedItems).Methods("GET")
	h.Router.HandleFunc("/items/active", h.GetActiveItems).Methods("GET")
	h.Router.HandleFunc("/items", h.CreateItem).Methods("POST")
	h.Router.HandleFunc("/items/{id}", h.UpdateItem).Methods("POST")
}
