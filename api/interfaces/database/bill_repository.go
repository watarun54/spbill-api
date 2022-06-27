package database

import (
	"github.com/watarun54/spbill-api/server/domain"
)

type BillRepository struct {
	SqlHandler
}

func (repo *BillRepository) FindOne(p domain.Bill) (bill domain.Bill, err error) {
	if err = repo.Debug().Preload("Payer").Preload("Payees").Find(&bill, p.ID).Error; err != nil {
		return
	}
	return
}

func (repo *BillRepository) FindAll(b domain.Bill) (bills domain.Bills, err error) {
	if err = repo.Debug().Where(&b).Preload("Payer").Preload("Payees").Find(&bills).Error; err != nil {
		return
	}
	return
}

func (repo *BillRepository) Store(b domain.Bill) (bill domain.Bill, err error) {
	if err = repo.Debug().Create(&b).Error; err != nil {
		return
	}
	if err = repo.Debug().Preload("Payer").Preload("Payees").Find(&bill, b.ID).Error; err != nil {
		return
	}
	return
}

func (repo *BillRepository) Update(b domain.Bill) (bill domain.Bill, err error) {
	if err = repo.Debug().Set("gorm:save_associations", false).Take(&domain.Bill{ID: b.ID}).Updates(&b).Error; err != nil {
		return
	}
	if err = repo.Debug().Model(&b).Association("Payees").Replace(b.Payees).Error; err != nil {
		return
	}
	if err = repo.Debug().Preload("Payer").Preload("Payees").Find(&bill, b.ID).Error; err != nil {
		return
	}
	return
}

func (repo *BillRepository) DeleteById(bill domain.Bill) (err error) {
	if err = repo.Debug().Model(&bill).Association("Payees").Replace(&[]domain.User{}).Error; err != nil {
		return
	}
	if err = repo.Debug().Delete(&bill).Error; err != nil {
		return
	}
	return
}
