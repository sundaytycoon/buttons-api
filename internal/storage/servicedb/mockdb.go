package servicedb

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/rs/zerolog/log"

	edgemysql "github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/internal/utils/retry"
	"github.com/sundaytycoon/buttons-api/internal/utils/unittestdocker"
)

type MockServiceDB struct {
	Adapter  *adapter
	Fixtures *testfixtures.Loader
}

func MockNew() *MockServiceDB {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	dockerMysql, err := unittestdocker.RunMySQL(ctx, "mysql")
	if err != nil {
		log.Panic().Err(err).Send()
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true",
		dockerMysql.Get("user"),
		dockerMysql.Get("password"),
		dockerMysql.ExternalHost,
		dockerMysql.ExternalPort,
		dockerMysql.Get("database"),
	)
	idleConnsStr := "10"
	openConnsStr := "10"

	idleConns, err := strconv.ParseInt(idleConnsStr, 10, 64)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	openConns, err := strconv.ParseInt(openConnsStr, 10, 64)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	db := edgemysql.MustNew(dsn, int(idleConns), int(openConns))
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if err = db.PingContext(ctx); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Panic().Err(err).Send()
			}
			continue
		}
		break
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("../testdata/fixtures"),
	)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	a := New(db, false)

	if err = a.Migrate(); err != nil {
		log.Panic().Err(err).Send()
	}

	return &MockServiceDB{
		Adapter:  a,
		Fixtures: fixtures,
	}
}

func (m *MockServiceDB) PrepareTestDatabase() {
	_, err := retry.Retry(5, time.Second*2, func() (interface{}, error) {
		subError := m.Fixtures.Load()
		if subError != nil {
			log.Debug().Err(subError).Msg("at fixture loading")
		}
		return nil, subError
	})
	if err != nil {
		log.Panic().Err(err).Send()
	}
}

func (m *MockServiceDB) Close() error {
	return m.Adapter.db.Close()
}
