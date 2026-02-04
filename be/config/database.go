package config

import (
	"fmt"
	"os"
	"rialfu/wallet/database/entities"

	// "rialf/wallet/database/entities"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunExtension(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
}

func SetUpTestDatabaseConnection() *gorm.DB {
	fmt.Println("start setup database")

	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPass := getEnvOrDefault("DB_PASS", "password")
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbName := getEnvOrDefault("DB_NAME", "test_db")
	dbPort := getEnvOrDefault("DB_PORT", "5432")
	migrate := getEnvOrDefault("AUTO_MIGRATE_DB", "N")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", dbHost, dbUser, dbPass, dbName, dbPort)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	db.Config.NowFunc = func() time.Time {
		return time.Now().In(loc)
	}
	RunExtension(db)
	if migrate == "y" {
		db.AutoMigrate(&entities.User{}, &entities.InformationUser{}, &entities.Wallet{}, &entities.MasterBank{}, &entities.Transaction{},
			&entities.Deposit{}, &entities.WalletLedger{}, &entities.WalletLedger{}, &entities.Withdrawal{})
	}
	return db
}

// func SetUpInMemoryDatabase() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

// func SetUpTestSQLiteDatabase() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db
// }

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}
	return defaultValue
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
