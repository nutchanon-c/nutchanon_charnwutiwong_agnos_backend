package password_test

import (
	"agnos/backend/src/usecase/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNumOfSteps(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test case 1",
			input:    "a",
			expected: 5,
		},
		{
			name:     "Test case 2",
			input:    "aA",
			expected: 4,
		},
		{
			name:     "Test case 3",
			input:    "aA1",
			expected: 3,
		},
		{
			name:     "Test case 4",
			input:    "1445D1cd",
			expected: 0,
		},
		{
			name:     "Test case 5",
			input:    "aaaBc1",
			expected: 1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := password.NewPasswordUsecase().CalculateNumOfSteps(password.PasswordRequest{InitPassword: test.input})
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result.NumOfSteps)
		})
	}
}
