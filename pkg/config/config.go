package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Database config
	DbUser    string        `default:"enviadoc"`
	DbPass    string        `default:"enviadoc"`
	DbHost    string        `default:"0.0.0.0"`
	DbPort    string        `default:"3306"`
	DbName    string        `default:"enviadoc"`
	DbTimeout time.Duration `default:"10s"`
	// Server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"20s"`
	// Security token
	// Secret      string `env:"TK_SECRET,required"`
	// ExpiredHour int64  `env:"TK_EXPIREDya HOUR,required"`
}

var Cfg config

func LoadConfig() error {
	err := envconfig.Process("IRIS", &Cfg)
	if err != nil {
		return err
	}
	return nil
}

func ConfigDb(ctx context.Context) (*sql.DB, error) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", Cfg.DbUser, Cfg.DbPass, Cfg.DbHost, Cfg.DbPort, Cfg.DbName)
	fmt.Println("uri: ", mysqlURI)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		fmt.Println("Failed database connection")
		panic(err)
	}

	fmt.Println("Successfully Connected to MySQL database")

	db.SetConnMaxLifetime(time.Minute * 4)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
