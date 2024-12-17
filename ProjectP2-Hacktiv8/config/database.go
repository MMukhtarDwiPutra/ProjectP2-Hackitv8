package config

import (
	"fmt"
	"log"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB menginisialisasi koneksi ke database PostgreSQL menggunakan GORM.
// Koneksi dibuat menggunakan konfigurasi yang sudah ditentukan, baik untuk lingkungan lokal maupun Supabase.
func InitDatabase() *gorm.DB {
	// // Konfigurasi untuk koneksi ke database
	db_host := "localhost"
	db_user := "postgres"
	db_password := "12345678"
	db_name := "GC3P2"
	db_port := "5432"

	// Konfigurasi untuk Supabase
	// db_host := "aws-0-ap-southeast-1.pooler.supabase.com"
	// db_user := "postgres.morsauyvkvtbwbgiyzvm"
	// db_password := "SupabasePtr16"
	// db_name := "postgres"
	// db_port := "6543"

	// Membuat DSN (Data Source Name) untuk koneksi ke PostgreSQL
	// Format: "host=<host> user=<user> password=<password> dbname=<dbname> port=<port> sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		db_host,
		db_user,
		db_password,
		db_name,
		db_port,
	)

	// Membuka koneksi ke database menggunakan GORM
	var err error
	 db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true, // Disables prepared statement caching
    }), &gorm.Config{})
    
	if err != nil {
		// Jika koneksi gagal, log error dan keluar dari aplikasi
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil
	}

	// Mengembalikan objek koneksi database GORM yang berhasil dibuka
	return db
}
