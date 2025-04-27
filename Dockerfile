# Etapa de compilación
FROM golang:1.23.4 AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar el código fuente primero
COPY . .

# Descargar dependencias
RUN go mod download

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go

# Etapa de ejecución
FROM alpine:latest

# Instalar dependencias necesarias
RUN apk --no-cache add ca-certificates tzdata

# Establecer variables de entorno
ENV TZ=UTC \
    APP_USER=appuser \
    APP_HOME=/app

# Crear un usuario no root
RUN addgroup -S $APP_USER && adduser -S $APP_USER -G $APP_USER

# Crear el directorio de la aplicación y asignar permisos
RUN mkdir -p ${APP_HOME} && \
    chown -R $APP_USER:$APP_USER ${APP_HOME}

# Establecer el directorio de trabajo
WORKDIR ${APP_HOME}

# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder --chown=$APP_USER:$APP_USER /app/main ${APP_HOME}/main

# Cambiar al usuario no root
USER $APP_USER

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]