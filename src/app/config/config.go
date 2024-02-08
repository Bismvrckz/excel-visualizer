package config

import (
	"excelvis/app/utils"

	_ "github.com/joho/godotenv/autoload"
)

var (
	BaseURL    string = utils.GetEnv("BASE_URL", "")
	DbExcelvis string = utils.GetEnv("DB_EXCELVIS_URL", "")
)
