package db

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Optional tuning you can override when calling Connect(...)
type Config struct {
	MaxOpen     int
	MaxIdle     int
	MaxLifetime time.Duration
}

// Connect loads .env, reads DATABASE_URL, opens GORM, tunes the pool, and pings the DB.
func Connect(cfg ...Config) *gorm.DB {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("env DATABASE_URL is empty (set it in .env or environment)")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect database (gorm): %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm: %v", err)
	}

	// Defaults (can be overridden via Connect(Config{...}))
	c := Config{
		MaxOpen:     10,
		MaxIdle:     5,
		MaxLifetime: 30 * time.Minute,
	}
	if len(cfg) > 0 {
		c = cfg[0]
	}

	sqlDB.SetMaxOpenConns(c.MaxOpen)
	sqlDB.SetMaxIdleConns(c.MaxIdle)
	sqlDB.SetConnMaxLifetime(c.MaxLifetime)

	// Verify connectivity
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}
	log.Println("DB connected")
	return db
}

// Close cleanly closes the underlying *sql.DB.
func Close(db *gorm.DB) {
	if db == nil {
		return
	}
	if sqlDB, err := db.DB(); err == nil {
		_ = sqlDB.Close()
	}
}
