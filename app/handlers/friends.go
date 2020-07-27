package handlers

func (h *Handler) ApiAddFriendHandler() error {
	h.writer.Write([]byte("<h1>Hello from ApiAddFriendHandler!</h1>"))

	return nil
}
