package auth

import "fmt"

// Dummy function for user authentication
func AuthenticateUser(username, password string) bool {
    // Replace with real authentication logic
    fmt.Println("Authenticating user:", username)
    return username == "admin" && password == "password"
}
