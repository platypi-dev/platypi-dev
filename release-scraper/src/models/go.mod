module platypi.dev/release-scraper/models

go 1.23.0

require (
	gorm.io/driver/sqlite v1.5.6
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.23 // indirect
	golang.org/x/text v0.18.0 // indirect
)

replace platypi.dev/release-scraper/controllers v0.0.0 => ../controllers
