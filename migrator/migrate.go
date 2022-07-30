package migrate

import (
	mysql_health "health-app/repository/health/mysql"
	mysql_users "health-app/repository/users/mysql"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&mysql_users.Users{}, &mysql_health.Health{})

}
