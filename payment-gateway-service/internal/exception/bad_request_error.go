package exception

type BadRequest struct {
	Message string
}

func (badRequest BadRequest) Error() string {
	return badRequest.Message
}
