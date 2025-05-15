package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"myapp/config"
	"myapp/infrastructure/database"
	"myapp/presentation/routes"
)

func init() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis padrão")
	}
}

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Inicializar o banco de dados
	db, err := database.SetupDatabase(cfg)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Configurar o router Gin
	router := gin.Default()
	router.SetTrustedProxies(nil) // Remover o aviso de proxies confiáveis

	// Setup dependências
	deps := config.SetupDependencies(db)

	// Configurar rotas com os serviços
	routes.SetupRoutes(router, deps)

	// Iniciar o servidor em uma goroutine
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("Servidor iniciado na porta %s\n", cfg.ServerPort)

	go func() {
		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	// Monitorar sinais de encerramento (CTRL+C, etc.)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nEncerrando o servidor...")

	// Fechar o banco de dados com segurança
	sqlDB, _ := db.DB()
	sqlDB.Close()

	// Tempo para garantir que o servidor finalize corretamente
	time.Sleep(1 * time.Second)
	fmt.Println("Servidor finalizado com sucesso.")
}
