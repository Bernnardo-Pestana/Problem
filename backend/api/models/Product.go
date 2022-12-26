package models

import (
	"errors"
	"time"
	"github.com/jinzhu/gorm"
)


type Product struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  string    `gorm:"size:255;not null;unique" json:"name"`
	Price float32    `gorm:"size:255;not null;" json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {

	var err error
	err = db.Debug().Create(&p).Error

	if err != nil {
		return &Product{}, err
	}
	return p, nil
}


func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, err
}


func (p *Product) FindProductByID(db *gorm.DB, uid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("User Not Found")
	}
	return p, err
}

func (p *Product) UpdateProduct(db *gorm.DB, pid uint32) (*Product,error) {
	var err error

	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"Name":  p.Name,
			"Price":  p.Price,
		},
	)
	if db.Error != nil {
		return &Product{}, db.Error
	}

	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product)DeleteProduct(db *gorm.DB, pid uint32) (int64,error){
	
	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).Delete(&Product{})
	
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

