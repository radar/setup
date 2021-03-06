package ruby

import (
	"errors"

	"github.com/radar/setup/output"
	"github.com/radar/setup/runner"
)

func checkForMongoid() error {
	return runner.OptionalAction("bundle show mongoid", checkForMongo, forgetAboutMongo)
}


func checkForMongo() error {
	output.Found("Found mongoid in the bundle.", 6)
	output.Info("Checking Mongo server is running...", 6)

	return runner.OptionalAction("nc -z localhost 27017", mongoRunning, mongoNotRunning)
}

func mongoRunning() error {
	output.Success("MongoDB is up and running at localhost:27017.", 8)
	return nil
}


func mongoNotRunning() error {
	output.Fail("MongoDB is not running at localhost:27017.", 8)
	output.Info("You may be able to start it with:", 10)
	output.Info("$ brew services start mongodb@3.4", 10)

	return errors.New("MongoDB is not running.")
}


func forgetAboutMongo() error {
	// Mongoid not found, so let's not worry
	return nil
}
