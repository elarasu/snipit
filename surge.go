package main

import (
        "fmt"

        "github.com/surgemq/surgemq/service"
)

func main() {
        // Create a new server
        svr := &service.Server{
                KeepAlive:        300,           // seconds
                ConnectTimeout:   2,             // seconds
                SessionsProvider: "mem",         // keeps sessions in memory
                Authenticator:    "mockSuccess", // always succeed
                TopicsProvider:   "mem",         // keeps topic subscriptions in memory
        }

        // Listen and serve connections at localhost:1883
        err := svr.ListenAndServe("tcp://:1883")
        fmt.Printf("%v", err)
}
