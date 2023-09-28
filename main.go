package main

import (
	"net/http"
)

// func init() {
// 	godotenv.Load()

// 	// if os.Getenv(`DEBUG`) == `Y` {
// 	// 	db := model.NewDatabase()
// 	// 	db.MigrateTables()
// 	// 	connDB := db.ConnDB()
// 	// 	sqlDB, _ := connDB.DB()

// 	// 	defer sqlDB.Close()

// 	// }
// }

func main() {
	// port := os.Getenv("PORT")
	port := ""
	if port == "" {
		port = "5000"
	}
	// test

	http.ListenAndServe(":"+port, RouterEngine())
	// http.ListenAndServe("", RouterEngine())
}
