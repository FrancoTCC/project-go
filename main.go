package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    // Ruta para la página de inicio (portafolio)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Servir el archivo index.html
        http.ServeFile(w, r, "index.html")
    })

    // Ruta para la página de contacto
    http.HandleFunc("/contacto", func(w http.ResponseWriter, r *http.Request) {
        // Servir el archivo contacto.html
        http.ServeFile(w, r, "contacto.html")
    })

    // Ruta para la página "Acerca de"
    http.HandleFunc("/acerca-de", func(w http.ResponseWriter, r *http.Request) {
        // Servir el archivo acerca-de.html
        http.ServeFile(w, r, "acerca-de.html")
    })

    // Configurar el puerto en el que se ejecutará el servidor
    fmt.Println("Servidor corriendo en http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

