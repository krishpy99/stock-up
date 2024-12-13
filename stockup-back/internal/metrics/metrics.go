package metrics

import "fmt"

// Dummy function to log API metrics
func LogRequest(endpoint string) {
    fmt.Printf("Request logged for endpoint: %s\n", endpoint)
}
