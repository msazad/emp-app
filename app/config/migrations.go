package config



func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Employee{}); err != nil {
		return fmt.Errorf("employee model migration failed due to : %v", err)
	}
	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return fmt.Errorf("admin model migration failed due to : %v", err)
	}
	log.Println("migration successfull...")

	if err := CheckAndSeedAdmin(db); err != nil {
		return fmt.Errorf("admin creation failed due to : %v", err)
	}
	log.Println("Added admin successfully..!")
	return nil
}