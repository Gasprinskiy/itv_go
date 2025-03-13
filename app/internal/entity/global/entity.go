package global

type MessageResponse struct {
	Message string `json:"message"`
}

type CreatedOrUpdatedResponse struct {
	ID int `json:"id"`
}
