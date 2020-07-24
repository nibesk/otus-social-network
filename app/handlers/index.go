package handlers

func (h *Handler) ViewIndexHandler() {
	h.writer.Write([]byte("<h1>Hello from index!</h1>"))
}
