package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rismanitalst/UTS_tumbler_store/config"
	"github.com/rismanitalst/UTS_tumbler_store/routes"
)

func main() {
	// 1. Load environment variables dari .env file
	err := godotenv.Load() // Tanpa argumen, dia cari file ".env" di folder yang sama
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}

	// 2. Baru ambil variabelnya
	secret := os.Getenv("JWT_SECRET")
	credPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	log.Println("JWT_SECRET:", secret)
	log.Println("GOOGLE_APPLICATION_CREDENTIALS:", credPath)
	// 2. Inisialisasi Firebase Admin SDK
	config.InitFirebase()
	// 3. Inisialisasi database + AutoMigrate
	config.InitDatabase()
	// 4. Setup Gin router dengan semua routes
	router := routes.SetupRouter()
	// 5. Jalankan server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di http://localhost:%s", port)
	log.Printf("Health check: http://localhost:%s/v1/health", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
