package handlers

func (h *Handler) ViewIndexHandler() error {
	h.writer.Write([]byte("<h1>Hello from service users!</h1>"))

	return nil
}
