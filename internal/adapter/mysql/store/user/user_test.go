package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	profilemeserver "github.com/sundaytycoon/profile.me-server"
	adaptermysql "github.com/sundaytycoon/profile.me-server/internal/adapter/mysql"
)

func init() {
	profilemeserver.TestInit()
}

func TestGetUser(t *testing.T) {
	a, err := adaptermysql.MockNew(profilemeserver.MySQLDocker)
	assert.Empty(t, err)

	userStore, err := New()
	assert.Empty(t, err)

	ctx := context.Background()
	conn, err := a.Conn(ctx)
	assert.Empty(t, err)
	tx, err := conn.BeginTx(ctx, nil)
	assert.Empty(t, err)
	expectedId := "1"
	u, err := userStore.GetUser(ctx, tx, expectedId)
	assert.Empty(t, err)
	err = tx.Commit()
	assert.Empty(t, err)
	err = conn.Close()
	assert.Empty(t, err)
	assert.EqualValues(t, expectedId, u.ID)
}
