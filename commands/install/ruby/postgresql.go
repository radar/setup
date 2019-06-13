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
	output.Found("Found pg in the bundle.", 6)
	output.Info("Checking PostgreSQL server is running...", 8)

	return runner.OptionalAction("nc -z localhost 5432", pgRunning, pgNotRunning)
}

func pgRunning() error {
	output.Success("PostgreSQL is up and running at localhost:5432.", 10)
	return nil
}


func pgNotRunning() error {
	output.Fail("PostgreSQL is not running at localhost:5432.", 10)
	output.Info("You may be able to start it with:", 10)
	output.Info("$ brew services start postgresql", 10)

	return errors.New("PostgreSQL is not running.")
}


func forgetAboutPg() error {
	// pg not found in bundle, so let's not worry
	return nil
}
