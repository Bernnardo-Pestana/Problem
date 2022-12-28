package models

import (
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)


type Purchase struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Id_User  uint32    `gorm:"uint" json:"id_user"`
	Id_Product uint32    `gorm:"uint" json:"id_product"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}



func (p *Purchase) SavePurchase(db *gorm.DB) (*Purchase, error) {

	var err error
	err = db.Debug().Create(&p).Error

	if err != nil {
		return &Purchase{}, err
	}
	return p, nil
}

func (p *Purchase) FindAllPurchases(db *gorm.DB) (*[]Purchase, error) {
	var err error
	purchases := []Purchase{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&purchases).Error
	if err != nil {
		return &[]Purchase{}, err
	}
	return &purchases, err
}

func (p *Purchase) FindPurchasesByID(db *gorm.DB, uid uint32) (*Purchase, error) {
	var err error
	err = db.Debug().Model(Purchase{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Purchase{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Purchase{}, errors.New("User Not Found")
	}
	return p, err
}

func (p *Purchase) UpdatePurchases(db *gorm.DB, pid uint32) (*Purchase,error) {
	var err error

	db = db.Debug().Model(&Purchase{}).Where("id = ?", pid).Take(&Purchase{}).UpdateColumns(
		map[string]interface{}{
			"Id_User":  p.Id_User,
			"Id_Product":  p.Id_Product,
		},
	)
	if db.Error != nil {
		return &Purchase{}, db.Error
	}

	err = db.Debug().Model(&Purchase{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Purchase{}, err
	}
	return p, nil
}

func (p *Purchase)DeletePurchases(db *gorm.DB, pid uint32) (int64,error){
	
	db = db.Debug().Model(&Purchase{}).Where("id = ?", pid).Take(&Purchase{}).Delete(&Purchase{})
	
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}