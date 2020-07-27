package handlers

func (h *Handler) ViewFlowHandler() error {
	h.writer.Write([]byte("<h1>Hello from Flow!</h1>"))

	return nil
}
