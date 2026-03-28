package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/JasperRosales/todo-api/internal/utils"
)

func TestPasswordHashing(t *testing.T) {
	pwd := "Jrosales26"

	hasher := utils.NewPasswordHasher()

	hash, err := hasher.Hash(pwd)
	assert.NoError(t, err, "Error hashing password")
	assert.NotEmpty(t, hash, "Hash should not be empty")

	ok, err := hasher.Check(hash, pwd)
	assert.NoError(t, err, "Error checking password")
	assert.True(t, ok, "Expected password match to be true")
}
