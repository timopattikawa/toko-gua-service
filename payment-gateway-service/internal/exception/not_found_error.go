package exception

type NotFoundException struct {
	Message string `json:"message"`
}

func (n NotFoundException) Error() string {
	return n.Message
}
