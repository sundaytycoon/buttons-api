package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
	profilemeserver "github.com/sundaytycoon/profile.me-server"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	"go.uber.org/dig"
)

func Test_Main(t *testing.T) {
	profilemeserver.TestInit()
	profilemeserver.ItNeedDockerWait()

	a, err := New(
		struct {
			dig.In
			ServiceDatabase *config.Database
		}{
			ServiceDatabase: &config.Database{
				Host:     profilemeserver.MySQLDocker.ExternalHost,
				Port:     profilemeserver.MySQLDocker.ExternalPort,
				User:     profilemeserver.MySQLDocker.Get("user"),
				Password: profilemeserver.MySQLDocker.Get("password"),
				Name:     profilemeserver.MySQLDocker.Get("name"),
				Dialect:  profilemeserver.MySQLDocker.Get("dialect"),
			},
		},
	)
	assert.Empty(t, err)
	rows, err := a.db.Query("SELECT 1+1")
	assert.Empty(t, err)
	var r *int64
	for rows.Next() {
		rows.Scan(r)
	}
	err = rows.Close()
	assert.Empty(t, err)

	assert.Equal(t, 2, *r)
	err = a.Close()
	assert.Empty(t, err)
}
