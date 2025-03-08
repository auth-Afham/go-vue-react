package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    _ "github.com/lib/pq" // Blank import to use pq for PostgreSQL
    "log"
    "database/sql"
    "fmt"
)

var db *sql.DB

// Initialize the database connection
func initDB() {
    var err error
    db, err = sql.Open("postgres", "user=postgres password=admin dbname=mydatabase sslmode=disable")
    if err != nil {
        log.Fatal("Error opening database: ", err)
    }

    // Check if the database connection is valid
    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging database: ", err)
    }
    fmt.Println("Connected to the database successfully!")
}

// Close the database connection
func closeDB() {
    if err := db.Close(); err != nil {
        log.Fatal("Error closing database: ", err)
    }
}

// Handler for the API
func getGreeting(c *gin.Context) {
    var message string
    err := db.QueryRow("SELECT message FROM greetings LIMIT 1").Scan(&message)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(404, gin.H{"error": "No greeting found"})
        } else {
            c.JSON(500, gin.H{"error": "Database error"})
        }
        return
    }
    c.JSON(200, gin.H{"message": message})
}

func main() {
    // Initialize the database
    initDB()
    defer closeDB()

    // Set up the Gin router
    router := gin.Default()

    // Enable CORS
    // router.Use(cors.Default())  // This will allow all origins, customize it if needed
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:3000",  // React frontend
            "http://localhost:8081",  // Another allowed origin (if needed)
            "https://yourdomain.com", // Example for a production environment
        },        
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
    }))    

    // Define the API route
    router.GET("/api", getGreeting)

    // Start the Gin server
    router.Run(":8080")
}
