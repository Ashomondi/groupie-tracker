package main

import (
    "log"
)

func main() {
    data, err := loadAllData()
    if err != nil {
        log.Fatal("Failed to load data:", err)
    }
    log.Printf("Loaded %d artists\n", len(data.Artists))

    Routes(data)

    port := ":8080"
    log.Printf("Server starting on http://localhost%s", port)
    log.Fatal(StartServer(port))
}