package buttonsapi

import (
	"context"

	"github.com/sundaytycoon/buttons-api/pkg/testdockercontainer"
)

var (
	MySQLDocker *testdockercontainer.DockerContainer
)

func TestInit() {
	if MySQLDocker == nil {
		err := RunContainers()
		if err != nil {
			panic(err)
		}
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
