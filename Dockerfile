# Usar la imagen oficial de Go
FROM golang:1.20-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Ejecutar la aplicación Go
CMD ["go", "run", "main.go"]
