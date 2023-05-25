package acceptance

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pressly/goose"
	"log"
	"os"
	"time"
)

const (
	DbName     = "marketplace"
	DbUser     = "acceptancetest"
	DbPassword = "pwd"
	DbPort     = "5431"
)

var (
	Db       *sql.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource

	Srv          server.Server
	Dependencies *configuration.AdConfiguration
)

func InitAcceptanceTest() {
	setEnvVariables()

	err := initDocker()
	if err != nil {
		stopDocker()
		log.Fatalf("something goes wrong :( %s", err)
	}

	err = setUpDatabase()
	if err != nil {
		stopDocker()
		log.Fatalf("something goes wrong :( %s", err)
	}
	Dependencies, _ = configuration.BuildAdConfiguration(Db, 10*time.Second)
	_, Srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, Dependencies)

	if err != nil {
		stopDocker()
		log.Fatalf("something goes wrong :( %s", err)
	}
}

func setEnvVariables() {
	os.Setenv("MARKETPLACE_DBUSER", DbUser)
	os.Setenv("MARKETPLACE_DBPORT", DbPort)
	os.Setenv("MARKETPLACE_DBPASS", DbPassword)
	os.Setenv("MARKETPLACE_DBNAME", DbName)
}

func StopAcceptanceTest() {
	stopDocker()
}

func initDocker() (err error) {
	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Failed to create Docker pool: %v", err)
	}
	resource, err = pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "14.7",
		ExposedPorts: []string{"5432/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432/tcp": {{HostIP: "", HostPort: DbPort}},
		},
		Env: []string{
			"POSTGRES_USER=" + DbUser,
			"POSTGRES_PASSWORD=" + DbPassword,
			"POSTGRES_DB=" + DbName,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Failed to run and purge Docker container: %v", err)
		}
	}

	if err = pool.Retry(func() error {
		Db, err = initDatabase()
		if err != nil {
			log.Println("Database not ready yet (it is booting up, wait for a few tries)...")
			return err
		}

		return err
	}); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	return err
}

func stopDocker() {
	err := pool.Purge(resource)
	if err != nil {
		log.Fatalf("No se pudo detener y eliminar los contenedores: %s", err)
	}
}

func initDatabase() (*sql.DB, error) {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		DbUser,
		DbPassword,
		"localhost",
		DbPort,
		DbName)

	postgresDb, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	// Ping DB to check if the connection was established successfully
	if err = postgresDb.Ping(); err != nil {
		return nil, err
	}
	return postgresDb, err
}

func setUpDatabase() error {
	err := goose.Up(Db, "../../../db/migrations")
	if err != nil {
		return err
	}
	return goose.Up(Db, "../db")
}
