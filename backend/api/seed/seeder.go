package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/Bernnardo-Pestana/Problem/api/models"
)


var users = []models.User{
	models.User{
		First_Name: "Robert",
		Last_Name: "Down",
		Type: "Seller",
		Email:    "robert@gmail.com",
		Password: "password",
	},
	models.User{
		First_Name: "Steven",
		Last_Name: "Victor",
		Type: "Customer",
		Email:    "steven@gmail.com",
		Password: "password",
	},
}

var products = []models.Product{
	models.Product{
		Name:"Pen",
		Price:1.25,
	},
	models.Product{
		Name:"Pencil",
		Price:1.50,
	},
}

var purchases = []models.Purchase{
	models.Purchase{
		
	},
	models.Purchase{
	
	},
}



func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{},&models.Product{},&models.Purchase{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{},&models.Product{},&models.Purchase{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}


	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
		purchases[i].Id_Product = products[i].ID

		purchases[i].Id_User = users[i].ID

		err = db.Debug().Model(&models.Purchase{}).Create(&purchases[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

	}
}