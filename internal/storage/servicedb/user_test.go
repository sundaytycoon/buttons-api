package servicedb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/edge/mysql"
)

func init() {
	buttonsapi.TestInit()
}

func TestGetUser(t *testing.T) {
	db, err := mysql.MockNew(buttonsapi.MySQLDocker)
	assert.Empty(t, err)

	serviceStorage := New()
	assert.Empty(t, err)

	ctx := context.Background()
	conn, err := db.Conn(ctx)
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
