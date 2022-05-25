package servebase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	assert := require.New(t)
	secret := "xixihahaheihei"
	id := int64(666)
	expire := int64(1)
	token, err := GenerateToken(secret, expire, id)
	assert.NoError(err)

	parsedId, err := ParseToken(token, secret)
	assert.NoError(err)
	assert.Equal(id, parsedId)

	time.Sleep(time.Second)
	_, err = ParseToken(token, secret)
	assert.Error(err)
}
