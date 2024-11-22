package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// Function to return a list of student names
func studentsNames() []string {
    return []string{"Reem", "Kinzy", "Yassin", "Toty Boty", "Koky"}
}

// Dashboard handler to display the student names
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    students := studentsNames()

    // Set the response header to HTML
    w.Header().Set("Content-Type", "text/html")
    
    // Write the HTML response, including the student names in a table
    fmt.Fprintf(w, `
    <html>
        <head>
            <title>Go Dashboard</title>
        </head>
        <body>
            <h1>Welcome to Go Dashboard</h1>
            <p>This is a simple student names dashboard!</p>
            <table border="1">
                <tr>
                    <th>Student Name</th>
                </tr>
    `)

    // Loop through the student names and add them as rows in the table
    for _, student := range students {
        fmt.Fprintf(w, "<tr><td>%s</td></tr>", student)
    }

    // Close the HTML table and body
    fmt.Fprintf(w, `
            </table>
        </body>
    </html>
    `)
}


// Function to test port connectivity with an HTML response
func testConnectivity(w http.ResponseWriter, r *http.Request) {
    // Define the host and port to check connectivity
    host := "localhost:8081" // Replace with the host and port you want to test

    // Set a timeout for the connection attempt
    timeout := 2 * time.Second
    conn, err := net.DialTimeout("tcp", host, timeout)

    // Set the response header to HTML
    w.Header().Set("Content-Type", "text/html")

    if err != nil {
        // Connection failed
        fmt.Fprintf(w, `
        <html>
            <head>
                <title>Port Status</title>
            </head>
            <body>
                <h1>Port Status</h1>
                <p>Port %s is not reachable: %v</p>
            </body>
        </html>
        `, host, err)
    } else {
        // Connection succeeded
        fmt.Fprintf(w, `
        <html>
            <head>
                <title>Port Status</title>
            </head>
            <body>
                <h1>Port Status</h1>
                <p>Port %s is reachable!</p>
            </body>
        </html>
        `, host)
        conn.Close() // Close the connection after testing
    }
}


func main() {
    // Route for the dashboard
    http.HandleFunc("/", dashboardHandler)

    // Route to test port connectivity
    http.HandleFunc("/test", testConnectivity)

    // Start the server on port 8081
    fmt.Println("Server is running on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
