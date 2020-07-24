package handlers

func (h *Handler) ViewFlowHandler() {
	h.writer.Write([]byte("<h1>Hello from Flow!</h1>"))
}
