package facebook

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGet_returns_data(t *testing.T) {
	// Not a great test, need to know the test user token in advance... Problems with dealing with FB?
	test_access_token := os.Getenv("ACCESS_TOKEN")
	fbUser := GetUser(test_access_token)

	assert.Equal(t, fbUser.Id, "10154466883655716")
	assert.Equal(t, fbUser.Name, "Ben Hall")
}
