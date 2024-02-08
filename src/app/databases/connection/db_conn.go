package dbconn

import (
	"database/sql"
	"excelvis/app/config"
	middlewares "excelvis/app/service/middleware"
	"time"
)

/**=======================================================================================================================
*!                                                   CONNECTION'S
*=======================================================================================================================**/

var (
	dbCms          = config.DbExcelvis
	dbMaxIdleConns = 10
	dbMaxConns     = 100
)

func ExcelvisConnection(ip string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbCms)
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "CmsConnection", "", &err)

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "CmsConnectionPing", "", &err)

		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
