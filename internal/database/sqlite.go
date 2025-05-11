package database

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetConnection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		var dsn string
		switch cfg.Driver {
		case "sqlite":
			dsn = cfg.Name
			db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		case "postgres":
			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		default:
			err = fmt.Errorf("unsupported database driver: %s", cfg.Driver)
		}

		if err != nil {
			return
		}

		err = db.AutoMigrate(
			&models.User{},
			&models.Brand{},
			&models.Type{},
			&models.Mode{},
			&models.EnergyClass{},
			&models.SplitSystem{},
			&models.SplitSystemImage{},
			&models.Review{},
		)
		err = populateDB(db)
		if err != nil {
			_ = fmt.Errorf("populateDB(db)")
			return
		}
	})

	if db == nil {
		return nil, err
	}

	return db, nil
}

func populateDB(db *gorm.DB) error {

	var brands = []models.Brand{
		{Name: "Tosot"}, {Name: "Lessar"}, {Name: "Midea"}, {Name: "Mitsubishi"},
		{Name: "Ballu"}, {Name: "Samsung"}, {Name: "LG"}, {Name: "Panasonic"},
		{Name: "Daikin"}, {Name: "Hitachi"}, {Name: "Dexp"}, {Name: "Centek"},
		{Name: "Electrolux"}, {Name: "Haier"}, {Name: "TCL"}, {Name: "Aceline"},
		{Name: "Bosch"}, {Name: "Daichi"}, {Name: "EcoClima"}, {Name: "Loriot"},
		{Name: "Marsa"}, {Name: "Pioneer"},
	}
	for _, b := range brands {
		db.FirstOrCreate(&b, models.Brand{Name: b.Name})
	}

	var types = []models.Type{
		{Name: "Настенные"}, {Name: "Кассетные"}, {Name: "Канальные"},
		{Name: "Напольно-потолочные"}, {Name: "Мультисплит-системы"},
	}
	for _, t := range types {
		db.FirstOrCreate(&t, models.Type{Name: t.Name})
	}

	var modes = []models.Mode{
		{Name: "охлаждение"}, {Name: "обогрев"}, {Name: "осушение"}, {Name: "вентиляция"},
	}
	for _, m := range modes {
		db.FirstOrCreate(&m, models.Mode{Name: m.Name})
	}

	var energyClasses = []models.EnergyClass{
		{Name: "A+++"}, {Name: "A++"}, {Name: "A+"}, {Name: "A"}, {Name: "B"}, {Name: "C"}, {Name: "D"},
	}
	for _, e := range energyClasses {
		db.FirstOrCreate(&e, models.EnergyClass{Name: e.Name})
	}

	// Добавим одну тестовую сплит-систему с привязкой к справочным данным
	var brand models.Brand
	db.First(&brand, "name = ?", "LG")

	var t models.Type
	db.First(&t, "name = ?", "Настенные")

	var coolingClass, heatingClass models.EnergyClass
	db.First(&coolingClass, "name = ?", "A++")
	db.First(&heatingClass, "name = ?", "A+")

	var modesList []models.Mode
	db.Where("name IN ?", []string{"охлаждение", "обогрев"}).Find(&modesList)

	system := models.SplitSystem{
		Title:                "Кондиционер настенный сплит-система DEXP AC-CD7ONF",
		BrandID:              brand.ID,
		TypeID:               t.ID,
		Price:                29990,
		HasInverter:          true,
		RecommendedArea:      25.0,
		CoolingPower:         2.5,
		EnergyClassCoolingID: coolingClass.ID,
		EnergyClassHeatingID: heatingClass.ID,
		MinNoiseLevel:        21.0,
		MaxNoiseLevel:        36.0,
		ExternalWeight:       28.5,
		ExternalWidth:        800,
		ExternalHeight:       545,
		ExternalDepth:        290,
		InternalWeight:       9.0,
		InternalWidth:        790,
		InternalHeight:       275,
		InternalDepth:        200,
	}

	db.FirstOrCreate(&system, models.SplitSystem{BrandID: system.BrandID, TypeID: system.TypeID, Price: system.Price})

	if len(modesList) > 0 {
		db.Model(&system).Association("Modes").Replace(&modesList)
	}

	fmt.Println("Database seeded successfully")
	return nil
}
