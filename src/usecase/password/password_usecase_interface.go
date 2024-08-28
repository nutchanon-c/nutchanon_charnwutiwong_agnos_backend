package password

type Usecase interface {
	CalculateNumOfSteps(request PasswordRequest) (PasswordResponse, error)
}

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}
