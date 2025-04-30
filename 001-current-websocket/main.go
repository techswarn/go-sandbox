package main

import (
    "log"
    "math/rand"
    "strconv"
    "time"

    "github.com/gorilla/websocket"
)

const (
    baseWebSocketURL = "wss://uvicron-cldru.ondigitalocean.app/ws/"
    numConnections   = 30             // Number of WebSocket connections to open
    messageInterval  = 1 * time.Second // Interval between messages
)

func main() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

	// Create a channel to monitor goroutines
    done := make(chan bool)

    // Create multiple WebSocket connections
    for i := 0; i < numConnections; i++ {
        go func(id int) {
            openWebSocketConnection(id)
            done <- true // Notify when a goroutine finishes
        }(i)
    }

	// Wait for all goroutines to finish (this prevents deadlock)
    for i := 0; i < numConnections; i++ {
        <-done
    }

    // Keep the main function running
    log.Println("All WebSocket connections closed. Exiting program.")

}

func openWebSocketConnection(id int) {
    // Generate a random user ID
    userID := generateRandomUserID()

    // Construct the WebSocket URL with the user ID
    websocketURL := baseWebSocketURL + userID

    // Connect to the WebSocket server
    conn, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
    if err != nil {
        log.Fatalf("Connection %d: Failed to connect to WebSocket server: %v", id, err)
    }
    defer conn.Close()

    log.Printf("Connection %d: Connected to WebSocket server with user ID: %s", id, userID)

    // Send a message every 2 seconds
    for {
        message := "Hello from user " + userID
        err := conn.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Printf("Connection %d: Failed to send message: %v", id, err)
            return
        }
        log.Printf("Connection %d: Sent message: %s", id, message)
        time.Sleep(messageInterval)
    }
}

func generateRandomUserID() string {
	// Generate a random user ID (e.g., "user12345")
    return "user" + strconv.Itoa(rand.Intn(100000))
}