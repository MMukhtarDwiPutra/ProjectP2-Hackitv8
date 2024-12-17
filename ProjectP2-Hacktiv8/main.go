package main

import(
	"P2-Hacktiv8/config"
	"P2-Hacktiv8/routes"
	// _ "P2-Hacktiv8/docs" // Import the generated Swagger docs
	"os"
)

func main(){
	db := config.InitDatabase()

	router := routes.NewRouter(db)

	// Get the port from the environment variable, default to 8080 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"  // Default to 8080 if no PORT is provided
    }

	router.Logger.Fatal(router.Start(":"+port))
}