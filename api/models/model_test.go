package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionModels(t *testing.T) {
	t.Run("justTest", func(t *testing.T) {
		result := fmt.Sprintf("sawerya-%s", randStringRune(16))
		assert.Equal(t, nil, result, "result must be nil")

	})
}
