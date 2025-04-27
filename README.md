# API Go Contenerizada

Este proyecto contiene una API Go contenerizada junto con una base de datos PostgreSQL configurada para iniciarse con un solo comando.

## Estructura del Proyecto

- `main.go`: Código principal de la API
- `Dockerfile`: Configuración multi-etapa para construir la imagen de la API
- `.dockerignore`: Archivos a ignorar en la construcción de la imagen
- `docker-compose.yml`: Configuración para orquestar la API y PostgreSQL
- `go.mod`: Gestión de dependencias de Go
- `Makefile`: Comandos útiles para gestionar el proyecto
- `.env`: Variables de entorno para desarrollo local

## Requisitos Previos

- Docker
- Docker Compose

## Base de Datos

El proyecto está configurado para trabajar con PostgreSQL como motor de base de datos.

## Cómo Iniciar

Para iniciar todos los servicios, ejecute:
