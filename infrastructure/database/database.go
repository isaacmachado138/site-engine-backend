package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"myapp/config"
)

// SetupDatabase configura a conexão com o banco de dados MySQL
func SetupDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao MySQL: %w", err)
	}

	// Auto Migrate - cria tabelas baseadas nas entidades
	/*err = db.AutoMigrate(
		&entities.User{},
	)*/
	if err != nil {
		return nil, fmt.Errorf("falha na migração do banco: %w", err)
	}

	// Configurar o pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("falha ao acessar conexão SQL: %w", err)
	}

	// Definir número máximo de conexões abertas
	sqlDB.SetMaxOpenConns(25)

	// Definir número máximo de conexões ociosas
	sqlDB.SetMaxIdleConns(10)

	// Definir tempo máximo de vida da conexão
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Definir tempo máximo de ociosidade
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return db, nil
}
