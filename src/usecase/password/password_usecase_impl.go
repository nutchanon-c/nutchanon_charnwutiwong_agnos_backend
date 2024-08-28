package password

import (
	"unicode"
)

type PasswordUsecase struct {
}

func NewPasswordUsecase() Usecase {
	return &PasswordUsecase{}
}

func (p *PasswordUsecase) CalculateNumOfSteps(request PasswordRequest) (PasswordResponse, error) {
	input := request.InitPassword

	result := getPasswordSteps(input)
	return PasswordResponse{NumOfSteps: result}, nil
}

func getPasswordSteps(input string) int {
	consecutive, lowerCase, upperCase, digit := getInputData(input)

	moves := 0

	moves += consecutive

	if len(input) < 6 {
		moves += 6 - len(input)
	} else if len(input) > 20 {
		moves += len(input) - 20
	}

	if lowerCase == 0 && moves == 0 {
		moves++
	}

	if upperCase == 0 && moves == 0 {
		moves++
	}

	if digit == 0 && moves == 0 {
		moves++
	}

	return moves
}

func getInputData(input string) (consecutive int, lower int, upper int, digit int) {
	cnt := 0
	totalCnt := 0
	lowerCnt := 0
	upperCnt := 0
	digitCnt := 0
	currentChar := string(input[0])
	for _, char := range input {
		if unicode.IsLower(char) {
			lowerCnt++
		} else if unicode.IsUpper(char) {
			upperCnt++
		} else if unicode.IsDigit(char) {
			digitCnt++
		}
		if string(char) == currentChar {
			cnt++
		} else {
			currentChar = string(char)
		}
		if cnt == 3 {
			totalCnt++
			cnt = 0
		}
	}
	return totalCnt, lowerCnt, upperCnt, digitCnt
}
