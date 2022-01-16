package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	buttonsapi "github.com/sundaytycoon/buttons-api"
)

func init() {
	buttonsapi.TestInit()
}

func Test_New(t *testing.T) {
	db, err := MockNew(buttonsapi.MySQLDocker)
	assert.Empty(t, err)
	rows, err := db.Query("SELECT 1+1")
	assert.Empty(t, err)
	var r int64
	for rows.Next() {
		rows.Scan(&r)
	}
	err = rows.Close()
	assert.Empty(t, err)

	assert.EqualValues(t, 2, r)
	err = db.Close()
	assert.Empty(t, err)
}
