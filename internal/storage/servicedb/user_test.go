package servicedb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	profilemeserver "github.com/sundaytycoon/profile.me-server"
	"github.com/sundaytycoon/profile.me-server/infrastructure/mysql"
)

func init() {
	profilemeserver.TestInit()
}

func TestGetUser(t *testing.T) {
	a, err := mysql.MockNew(profilemeserver.MySQLDocker)
	assert.Empty(t, err)

	serviceStorage := New()
	assert.Empty(t, err)

	ctx := context.Background()
	conn, err := a.DB.Conn(ctx)
	assert.Empty(t, err)
	tx, err := conn.BeginTx(ctx, nil)
	assert.Empty(t, err)
	expectedId := "1"
	u, err := serviceStorage.GetUser(ctx, tx, expectedId)
	assert.Empty(t, err)
	err = tx.Commit()
	assert.Empty(t, err)
	err = conn.Close()
	assert.Empty(t, err)
	assert.EqualValues(t, expectedId, u.ID)
}
