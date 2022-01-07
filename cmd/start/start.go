package main

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sundaytycoon/buttons-api/ent"
	entcar "github.com/sundaytycoon/buttons-api/ent/car"
	entgroup "github.com/sundaytycoon/buttons-api/ent/group"
	entuser "github.com/sundaytycoon/buttons-api/ent/user"
)

func Println(title string, i ...interface{}) {
	blue := color.New(color.FgBlue).SprintfFunc()
	fmt.Println(blue(title), i)
}

func main() {

	logCtx := func(ctx context.Context, i ...interface{}) {
		Println("MYSQL[LOG]", i)
	}
	//client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	log.Fatalf("failed opening connection to sqlite: %v", err)
	//}
	drv, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true",
		"buttons", "p@ssword", "0.0.0.0", "33307", "buttons",
	))
	dbgDrv := dialect.DebugWithContext(drv, logCtx)
	client := ent.NewClient(ent.Driver(dbgDrv))
	if err != nil {
		log.Fatal().Msgf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.

	fmt.Println()
	fmt.Println()
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}
	log.Info().Msgf("\n")

	fmt.Println()
	fmt.Println()
	u, err := CreateUser(context.Background(), client)
	if err != nil {
		log.Fatal().Str("operation", "CreateUser").Err(err).Send()
	}
	log.Info().Str("operation", "CreateUser").Interface("user", u).Send()

	fmt.Println()
	fmt.Println()
	cars, err := CreateCars(context.Background(), client)
	if err != nil {
		log.Fatal().Str("operation", "CreateCars").Err(err).Send()
	}
	log.Info().Str("operation", "QueryUser").Interface("cars", cars).Send()

	fmt.Println()
	fmt.Println()
	u, err = client.User.Query().Where(entuser.ID(2)).Only(context.Background())
	if err != nil {
		log.Fatal().Str("operation", "get user a7m").Err(err).Send()
	}
	fmt.Println()
	fmt.Println()
	err = QueryCars(context.Background(), u)
	if err != nil {
		log.Fatal().Str("operation", "QueryCars").Err(err).Send()
	}

	fmt.Println()
	fmt.Println()
	err = QueryCarUsers(context.Background(), u)
	if err != nil {
		log.Fatal().Str("operation", "QueryCarUsers").Err(err).Send()
	}

	fmt.Println()
	fmt.Println()
	err = CreateGraph(context.Background(), client)
	if err != nil {
		log.Fatal().Str("operation", "CreateGraph").Err(err).Send()
	}

	fmt.Println()
	fmt.Println()
	cars, err = QueryGithub(context.Background(), client)
	if err != nil {
		log.Fatal().Str("operation", "CreateGraph").Err(err).Send()
	}
	log.Debug().Interface("cars", cars).Msg("cars returned")
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) ([]*ent.Car, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Info().Str("operation", "CreateCars.createCar").Interface("car", tesla).Send()

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Info().Str("operation", "CreateCars.createCar").Interface("car", ford).Send()

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a7m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Info().Str("operation", "CreateCars.createUser").Interface("user", a8m).Send()
	return []*ent.Car{tesla, ford}, nil
}

func QueryCars(ctx context.Context, a7m *ent.User) error {
	cars, err := a7m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Info().Str("operation", "QueryCars.a8m.QueryCars.All").Interface("user", a7m).Interface("cars", cars).Send()

	// What about filtering specific cars.
	ford, err := a7m.QueryCars().
		Where(entcar.Model("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Info().Str("operation", "QueryCars.a8m.QueryCars.Ford").Interface("user", a7m).Interface("ford", ford).Send()
	return nil
}

func QueryCarUsers(ctx context.Context, a7m *ent.User) error {
	cars, err := a7m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	// Query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", ca.Model, err)
		}
		log.Debug().Msgf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users in the creation.
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Exec(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Info().Msgf("Traversal Graph")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) ([]*ent.Car, error) {
	cars, err := client.Group.
		Query().
		Where(entgroup.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                   // (User(Name=Ariel, Age=30),)
		QueryCars().                    // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting cars: %w", err)
	}
	return cars, nil
}
