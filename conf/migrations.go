package conf

import "fmt"

// MakeMigrations auto migrate the passed models.
func MakeMigrations() {
	if err := Db.Debug().AutoMigrate(); err != nil {
		fmt.Printf("Could not make migrations: %s\n", err)
		return
	}
	return
}
