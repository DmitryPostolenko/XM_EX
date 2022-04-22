package database

import (
	"log"

	"github.com/uptrace/bun"
)

type DatabaseManager interface {
	GetDB() *bun.DB
}

type Pool struct {
	DB *bun.DB
}

func (p *Pool) GetDB() *bun.DB {
	return p.DB
}

// To make it simplier to set up DB for testing purposes, not using mirgations.
func InitDatabase(p *Pool) *Pool {
	_, err := p.DB.Query("create database xmex")
	if err != nil {
		log.Println("DB already exists")
	}
	_, err = p.DB.Query(
		`create table if not exists users
			(
				id uuid primary key,
				username varchar,
				password varchar,
			);
			create unique index if not exists users_id
				on users (id);
			create unique index  if not exists users_username
				on users (username);
			
			create table if not exists companies
			(
				id uuid primary key,
				name varchar,
				code varchar,
				country varchar,
				website varchar,
				phone varchar,
			);
			create unique index if not exists companies_id
				on companies (id);
			create unique index  if not exists companies_name
				on companies (name);
			create unique index  if not exists companies_code
				on companies (code);
			create unique index  if not exists companies_country
				on companies (country);
			create unique index  if not exists companies_website
				on companies (website);
			create unique index  if not exists companies_phone
				on companies (phone);
			`)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
