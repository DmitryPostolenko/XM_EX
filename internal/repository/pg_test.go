package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
)

func DBConnection() *bun.DB {
	// DB CONFIGURATION AND SETTINGS
	cfg, _ := configuration.Load("../../internal/configuration/configuration.yml")
	pool := Pool{}
	var dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Host, cfg.DataBase.Port, cfg.DataBase.Name)
	}

	dbc := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbUrl)))
	pool.DB = bun.NewDB(dbc, pgdialect.New())

	dbRep := InitDatabase(&pool)
	dbConnect := dbRep.GetDB()

	return dbConnect
}

func TestDB(t *testing.T) {
	dbRep := DBConnection()
	DBT := fmt.Sprintf("%T", dbRep)
	if DBT != "*bun.DB" {
		t.Fatal("Wrong DB connection type")
	}
}
