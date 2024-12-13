package main

import (
    "github.com/gin-gonic/gin"
    "stockup-back/internal/auth"
    "stockup-back/internal/metrics"
    "stockup-back/internal/model"
    "net/http"
)

func main() {
    r := gin.Default()

    // Route: Health Check
    r.GET("/health", func(c *gin.Context) {
        metrics.LogRequest("/health")
        c.JSON(http.StatusOK, gin.H{"status": "healthy"})
    })

    // Route: Login
    r.POST("/login", func(c *gin.Context) {
        var credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        if err := c.BindJSON(&credentials); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        if auth.AuthenticateUser(credentials.Username, credentials.Password) {
            c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        }
    })

    // Route: Get User
    r.GET("/user/:id", func(c *gin.Context) {
        metrics.LogRequest("/user/:id")
        user := model.User{
            ID:       1,
            Username: "admin",
            Email:    "admin@example.com",
        }
        c.JSON(http.StatusOK, user)
    })

    r.Run(":8080") // Start server on port 8080
}
