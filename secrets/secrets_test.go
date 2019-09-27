package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecret(t *testing.T) {

	tests := []struct {
		requestSecretName string
		expectResult      string
		err               error
	}{
		{
			requestSecretName: "humn/test/getSecretTest",
			expectResult:      "abc123",
			err:               nil,
		},
	}

	for _, test := range tests {
		result, errGetSecret := GetSecret(test.requestSecretName)
		assert.EqualValues(t, test.expectResult, result)
		assert.IsType(t, test.err, errGetSecret)
	}
}
