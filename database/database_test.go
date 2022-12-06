package database_test

import (
	"testing"

	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/database"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabase(t *testing.T) {
	db := database.GetDatabase()

	assert.NotNil(t, db, "db instance should not be nil")
}
