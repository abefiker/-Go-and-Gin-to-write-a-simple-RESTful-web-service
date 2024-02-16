package main

import (
    "example/web-service-gin/router"
)

func main() {
    // Initialize the router
    r := router.SetupRouter()

    // Start the server
    r.Run("localhost:4000")
}
