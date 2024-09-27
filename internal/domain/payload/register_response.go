package payload

type RegisterResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
