package handlers

func (h *Handler) ApiAddFriendHandler() {
	h.writer.Write([]byte("<h1>Hello from ApiAddFriendHandler!</h1>"))
}
