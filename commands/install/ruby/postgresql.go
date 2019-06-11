package ruby

import (
	"errors"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkForPG() error {
	return runner.OptionalAction("bundle show pg", checkForPostgresql, forgetAboutPg)
}

func checkForPostgresql() error {
	output.Found("Found pg in the bundle.")
	output.Info("Checking PostgreSQL server is running...")

	return runner.OptionalAction("nc -z localhost 5432", pgRunning, pgNotRunning)
}

func pgRunning() error {
	output.Success("PostgreSQL is up and running at localhost:5432.")
	return nil
}


func pgNotRunning() error {
	output.Fail("PostgreSQL is not running at localhost:5432.")
	output.Info("You may be able to start it with:")
	output.Info("$ brew services start postgresql")

	return errors.New("PostgreSQL is not running.")
}


func forgetAboutPg() error {
	// pg not found in bundle, so let's not worry
	return nil
}
