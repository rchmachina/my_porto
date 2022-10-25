package Config

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"


	_ "github.com/lib/pq"
)
  
var db *gorm.DB
var err error



const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hasbiR26"
	dbname   = "my_porto"
  )
  

type Portofolio struct{
	gorm.Model
	
	id int `gorm: "primary_key"`
	name string
	category string
	path_image string
	created string
	
	
}


var (
	portofolio1 = &Portofolio{id:69, name:"percobaain it", category: "science",path_image: "/static/img2.png" ,created: time.Now().Format("2006-01-02")}
)


func Connect() gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	  if err != nil {
		fmt.Println("cant connect")
		panic(err)
	  }
	  defer db.Close()


 	 fmt.Println("Successfully connected!")


	 db.Create(portofolio1)

	return gorm.DB{}
  }



