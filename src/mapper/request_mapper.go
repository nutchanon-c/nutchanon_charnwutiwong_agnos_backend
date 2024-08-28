package mapper

import "agnos/backend/src/usecase/password"

func MapToPasswordRequest(body map[string]interface{}) password.PasswordRequest {
	return password.PasswordRequest{
		InitPassword: body["init_password"].(string),
	}
}
