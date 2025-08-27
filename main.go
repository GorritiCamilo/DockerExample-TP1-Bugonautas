package main

import (
    "fmt"
    "net/http"
    "time"
    "os"
    "sync"
    "github.com/splitio/go-client/v6/splitio/client"
	"github.com/splitio/go-client/v6/splitio/conf"
)

var (
    visitCounter int
    mu           sync.Mutex
)

func main() {   
    apiKey := os.Getenv("SPLIT_SDK_API_KEY")

    cfg := conf.Default()
	factory, err := client.NewSplitFactory(apiKey, cfg)
	if err != nil {
		fmt.Printf("SDK init error: %s\n", err)
		return
	}

	splitClient := factory.Client()
	err = splitClient.BlockUntilReady(10)
	if err != nil {
		fmt.Printf("SDK timeout: %s\n", err)
		return
	}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        visitCounter++
        userID := visitCounter
        mu.Unlock()
        backgroundColor := "#f0f4f8"

        featureFlagName := os.Getenv("SPLIT_FEATURE_FLAG_NAME")
        treatment := splitClient.Treatment(userID, featureFlagName, nil)

        if treatment == "on" {
            backgroundColor = "#08ac23ff"
        } else if treatment == "off" {
        } else {
        }

        fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html lang="es">
            <head>
                <meta charset="UTF-8">
                <title>¡Hola Mundo desde Go!</title>
                <style>
                    body {
                        font-family: 'Segoe UI', Arial, sans-serif;
                        background-color: %s;
                        display: flex;
                        flex-direction: column;
                        align-items: center;
                        justify-content: center;
                        min-height: 100vh;
                        margin: 0;
                    }
                    .container {
                        background: #fff;
                        padding: 2rem 3rem;
                        border-radius: 16px;
                        box-shadow: 0 4px 24px rgba(0,0,0,0.08);
                        text-align: center;
                    }
                    h1 {
                        color: #2e7d32;
                        margin-bottom: 0.5em;
                    }
                    p {
                        color: #333;
                        margin: 0.5em 0;
                    }
                    .split {
                        margin-top: 1.5em;
                        font-weight: bold;
                        color: #1565c0;
                    }
                    button {
                        padding:0.5em 1.5em;
                        font-size:1em;
                        border-radius:8px;
                        background:#1565c0;
                        color:#fff;
                        border:none;
                        cursor:pointer;
                        margin-top:1em;
                    }
                </style>
            </head>
            <body>
                <div class="container">
                    <h1>¡Hola Mundo desde Go! 🐹</h1>
                    <p>Servidor web en Go con Docker</p>
                    <p>Fecha: %s</p>
                    <p>Número de visita: %d</p>
                    <p class="split">Tratamiento Split: %s</p>
                    <form action="/reset" method="post">
                        <button type="submit">Reiniciar contador</button>
                    </form>
                </div>
            </body>
            </html>
            `, backgroundColor, time.Now().Format("2006-01-02 15:04:05"), userID, treatment)
    })

    http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        visitCounter = 0
        mu.Unlock()
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })

    fmt.Println("🚀 Servidor Go iniciado en puerto 8080")
    http.ListenAndServe(":8080", nil)
}