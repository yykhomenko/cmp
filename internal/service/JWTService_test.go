package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValidateToken(t *testing.T) {
	jwtService := NewJWTService()
	tokenStr := jwtService.GenerateToken("yykhomenko", true)
	assert.NotEmpty(t, tokenStr)

	token, err := jwtService.ValidateToken(tokenStr)
	assert.Nil(t, err)
	assert.True(t, token.Valid)
}
