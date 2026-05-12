package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	ok := CheckPassword(hash, "password")
	assert.NoError(t, err)
	assert.True(t, ok)
}
