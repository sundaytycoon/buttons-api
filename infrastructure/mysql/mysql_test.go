package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	profilemeserver "github.com/sundaytycoon/profile.me-server"
)

func init() {
	profilemeserver.TestInit()
}

func Test_New(t *testing.T) {
	a, err := MockNew(profilemeserver.MySQLDocker)
	assert.Empty(t, err)
	rows, err := a.DB.Query("SELECT 1+1")
	assert.Empty(t, err)
	var r int64
	for rows.Next() {
		rows.Scan(&r)
	}
	err = rows.Close()
	assert.Empty(t, err)

	assert.EqualValues(t, 2, r)
	err = a.Close()
	assert.Empty(t, err)
}
