# Proyecto - Aplicación Go con Docker

Instrucciones para construir y ejecutar la aplicación Go usando Docker.

## Requisitos

- Docker instalado en tu máquina.

## Pasos

1. **Construir la imagen Docker:**

    ```bash
    docker build -t my-go-app .
    ```

2. **Ejecutar el contenedor:**

    ```bash
    docker run -p 8080:8080 my-go-app
    ```


¡Listo! Si todo funciona correctamente, tu aplicación Go estará en ejecución.
