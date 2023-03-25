package bootstrap

import (
	"context"
	"fmt"
	"log"

	database "github.com/dany0814/go-hexagonal/internal/core/adapters/outgoing"
	"github.com/dany0814/go-hexagonal/internal/core/application"
	"github.com/dany0814/go-hexagonal/internal/platform/server"
	mysqldb "github.com/dany0814/go-hexagonal/internal/platform/storage/mysql"
	"github.com/dany0814/go-hexagonal/pkg/config"
)

func Run() error {

	err := config.LoadConfig()
	if err != nil {
		return err
	}
	fmt.Println("Web server ready!")

	ctx := context.Background()
	db, err := config.ConfigDb(ctx)

	if err != nil {
		log.Fatalf("Database configuration failed: %v", err)
	}

	userRepository := mysqldb.NewUserRepository(db, config.Cfg.DbTimeout)
	userAdapter := database.NewUserAdapter(userRepository)
	userService := application.NewUserService(userAdapter)

	ctx, srv := server.NewServer(context.Background(), config.Cfg.Host, config.Cfg.Port, config.Cfg.ShutdownTimeout, server.AppService{
		UserService: userService,
	})

	return srv.Run(ctx)
}
