package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aristorinjuang/blog/controllers"
	_ "github.com/go-sql-driver/mysql"
)

/*
	func init() {
		controllers.Sessions = make(map[string]*models.User)

		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		models.Db, models.Err = sql.Open("mysql", os.Getenv("DATABASE"))
		if models.Err != nil {
			log.Fatal(models.Err)
		}

		models.Err = models.Db.Ping()
		if models.Err != nil {
			log.Fatal(models.Err)
		}
	}
*/
func main() {
	// get current directory
	dir, _ := os.Getwd()
	log.Println(dir)
	http.HandleFunc("/", controllers.Index)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(dir+"/assets"))))
	log.Println(http.ListenAndServe(":3000", nil))
}
