package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No se pudo cargar el archivo .env: %v", err)
	}
}

func main() {
	/*
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		r.Use(gin.Logger())
		r.Use(gin.Recovery())

		log.Printf("Conectando a PostgreSQL en %s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

		db, err := config.InitializeApp()
		if err != nil {
			log.Fatalf("Error fatal conectando a PostgreSQL: %v", err)
		}

		log.Println("Conexión exitosa a PostgreSQL")

		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Printf("Error al cerrar la conexión con PostgreSQL: %v", err)
			}
		}(db.DB)

		r.GET("/health", func(c *gin.Context) {
			dbStatus := false
			var dbError string

			if db != nil && db.DB != nil {
				err := db.DB.Ping()
				if err != nil {
					dbError = err.Error()
				}
				dbStatus = err == nil
			} else {
				dbError = "db instance is nil"
			}

			c.JSON(http.StatusOK, gin.H{
				"http_status": "ok",
				"postgres":    dbStatus,
				"db_error":    dbError,
				"timestamp":   time.Now().Unix(),
			})
		})

		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API Go funcionando correctamente",
				"time":    time.Now().Format(time.RFC3339),
			})
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		log.Printf("Servidor iniciado en el puerto %s", port)
		if err := r.Run(":" + port); err != nil {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	*/

}
