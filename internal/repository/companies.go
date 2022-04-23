package repository

import (
	"context"
	"log"

	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
)

type CompaniesManager interface {
	SaveCompany(models.Company) error
	ListCompanies() (models.Company, bool)
	FindCompany(field string) (models.Company, bool)
	//UpdateCompany(field string) error
	DeleteCompany(cid string) bool
}

func (p *Pool) SaveCompany(msg models.Company) error {
	ctx := context.Background()
	_, err := p.DB.NewInsert().Model(&msg).Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *Pool) FindCompany(field string, value string) (models.Company, bool) {
	ctx := context.Background()
	company := new(models.Company)
	if err := p.DB.NewSelect().Model(company).Where("? = ?", bun.Ident(field), value).Scan(ctx); err != nil {
		log.Println("Company does not exist")
		return models.Company{}, false
	}
	return *company, true
}

func (p *Pool) DeleteCompany(cid string) bool {
	ctx := context.Background()
	company := new(models.Company)

	_, err := p.DB.NewDelete().Model(company).Where("id = ?", cid).Exec(ctx)
	if err != nil {
		return false
	}
	return true

}

func (p *Pool) ListCompanies() ([]models.Company, bool) {
	ctx := context.Background()
	msg := make([]models.Company, 0)
	if err := p.DB.NewSelect().Model(&msg).Where("true").Scan(ctx); err != nil {
		log.Println("No companies found")
		return []models.Company{}, false
	}
	return msg, true
}

func GetCompaniesRepository(db *bun.DB) *Pool {
	return &Pool{DB: db}
}
