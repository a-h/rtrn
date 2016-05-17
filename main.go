package main

import (
	"flag"
	"github.com/a-h/rtrn/dataaccess/mongo"
	"github.com/a-h/rtrn/httpservice"
)

var connectionString = flag.String("connectionString", "mongodb://localhost:27017",
	"The MongoDB connection string used to store data.")

// The main partition; instantiate all of the concrete implementations and wire together
func main() {
	flag.Parse()
	da := mongo.NewMongoDataAccess(*connectionString)
	httpservice.Serve(da)
}
