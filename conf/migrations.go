package conf

import log "github.com/sirupsen/logrus"

// MakeMigrations auto migrate the passed models.
func MakeMigrations() {
	if err := Db.Debug().AutoMigrate(); err != nil {
		log.Error("Could not make migrations")
		return
	}
	return
}
