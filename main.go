package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>¡Hola Mundo desde Go! 🐹</h1>")
        fmt.Fprintf(w, "<p>Servidor web en Go con Docker</p>")
        fmt.Fprintf(w, "<p>Fecha: %s</p>", time.Now().Format("2006-01-02 15:04:05"))
    })

    http.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
        nombre := r.URL.Query().Get("nombre")
        if nombre == "" {
            nombre = "Amigo"
        }
        fmt.Fprintf(w, "¡Hola %s! 👋", nombre)
    })

    fmt.Println("🚀 Servidor Go iniciado en puerto 8080")
    http.ListenAndServe(":8080", nil)
}