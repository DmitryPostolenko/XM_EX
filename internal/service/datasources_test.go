package service

import (
	"fmt"
	"testing"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
)

func TestSetUpDB(t *testing.T) {
	conf, _ := configuration.Load("../../internal/configuration/configuration.yml")
	dbConn := setUpDB(conf)

	DBT := fmt.Sprintf("%T", dbConn)
	if DBT != "*bun.DB" {
		t.Fatal("Wrong DB connection type")
	}
}

func TestSetUpRedis(t *testing.T) {
	conf, _ := configuration.Load("../../internal/configuration/configuration.yml")
	dbConn := setUpRedis(conf)

	RT := fmt.Sprintf("%T", dbConn)
	if RT != "*redis.Client" {
		t.Fatal("Wrong redis connection config")
	}
}

func TestSetUpPort(t *testing.T) {
	setPort := "8000"
	p := setUpPort(setPort)
	if p != setPort {
		t.Fatal("Port setting up failed")
	}
}
