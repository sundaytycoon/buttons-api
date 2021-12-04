package profilemeserver

import (
	"context"
	"sync"

	"github.com/sundaytycoon/profile.me-server/pkg/testdockercontainer"
)

var (
	wg          = sync.WaitGroup{}
	MySQLDocker *testdockercontainer.DockerContainer
)

func TestInit() {
	if MySQLDocker == nil {
		wg.Add(1)
		err := RunContainers()
		if err != nil {
			panic(err)
		}
		wg.Done()
	}
}

func RunContainers() error {
	ctx := context.Background()
	mysqlDocker, err := testdockercontainer.RunMySQL(
		ctx,
		"mysql",
		"test-db",
		"test-user",
		"test-password",
	)
	if err != nil {
		return nil
	}
	MySQLDocker = mysqlDocker

	return nil
}

func ItNeedDockerWait() {
	wg.Wait()
}
