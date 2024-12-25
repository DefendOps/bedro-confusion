package database

func MigrateModels() {
	db := GetDB()
	db.AutoMigrate(
		// &SCPackage{},
	)
}