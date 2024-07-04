package main

import (
	"api-notification/initializers"
	model "api-notification/models"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseInit()
}

func main() {

	models := []struct {
		model interface{}
		table string
	}{
		{&model.Message{}, "messages"},
	}

	for _, m := range models {
		if initializers.DB.Migrator().HasTable(m.model) {
			if err := initializers.DB.Migrator().DropTable(m.model); err != nil {
				log.Printf("Error dropping table %s: %v\n", m.table, err)
			}
		}
	}

	for _, m := range models {
		if err := initializers.DB.AutoMigrate(m.model); err != nil {
			log.Printf("Error migrating model %T: %v\n", m.model, err)
		}
	}
	log.Println("Database migration completed successfully")
}
