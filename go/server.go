package main

import (
    "log"
    "net/http"

    "github.com/checkout/sample-projects/go/src/api/application"
)

func main() {
    app := application.BuildApplication()

    fs := http.FileServer(http.Dir("resources"))

    http.Handle("/", fs)
    http.HandleFunc("/payments", app.PaymentsController.RequestPayment)

    addr := "localhost:4242"
    log.Printf("Listening on %s ...", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
