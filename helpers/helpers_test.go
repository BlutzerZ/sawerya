package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwtHelpers(t *testing.T) {
	var token string

	t.Run("generateAccessToken", func(t *testing.T) {
		var err error
		ID := 1
		username := "jamal"
		token, err = GenerateAccessToken(uint(ID), username)
		fmt.Println("Token =", token)
		assert.Equal(t, nil, err, "result must be nil")
	})

	t.Run("ValidateToken", func(t *testing.T) {
		claims, err := ValidateToken(token)
		fmt.Println("Claims ID", claims.ID)
		fmt.Println("Claims Username", claims.Username)
		assert.Equal(t, nil, err, "result must be nil")
	})
	t.Run("updateToken", func(t *testing.T) {
		newToken, err := UpdateToken(token)
		fmt.Println(newToken)
		assert.Equal(t, nil, err, "result must be nil")
	})
}
