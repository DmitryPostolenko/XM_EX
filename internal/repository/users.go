package repository

import (
	"context"
	"log"

	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
)

type UserManager interface {
	SaveUser(models.User) error
	GetUserByField(field string, username string) (models.User, bool)
}

func (p *Pool) SaveUser(us models.User) error {
	ctx := context.Background()
	_, err := p.DB.NewInsert().Model(&us).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) GetUserByField(field string, value string) (models.User, bool) {
	ctx := context.Background()
	user := new(models.User)
	if err := p.DB.NewSelect().Model(user).Where("? = ?", bun.Ident(field), value).Scan(ctx); err != nil {
		log.Println("User does not exist")
		return models.User{}, false
	}
	return *user, true
}

func GetUsersRepository(db *bun.DB) *Pool {
	return &Pool{DB: db}
}
