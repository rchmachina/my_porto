package handler

import (


	"html/template"
	"net/http"

	. "my_porto/Config"

	"github.com/jinzhu/gorm"
)

// import "fmt"

func Home(w http.ResponseWriter, r* http.Request){
	var tmpl, err = template.ParseFiles("Views/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	var data = map[string]interface{}{
		"title": "Welcome to my portofolio",
		"Home" : "home",
		"About" : "About",
		"Project" : "Project",
		"Contact" : "contact",
		"Name" : "hasbielRch",
		"Job" : "Programmer | Math enthusiast | Tech enthusiast",
		"wa_contact" : "https://wa.me/6285811117978?text=I'm%20interested%20working%20with%20you",
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


var db *gorm.DB

func Test(w http.ResponseWriter, r* http.Request){

	var Porto []Portofolio

	
	var tmpl, err = template.ParseFiles("Views/test.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data = map[string]interface{}{
		"title": db.Find(&Porto),
	}
	

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// json.NewEncoder(w).Encode(Porto)


}