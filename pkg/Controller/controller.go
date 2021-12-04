package Controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"home.page.gohtml")
}
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"about.page.gohtml")
}


func renderTemplate(w http.ResponseWriter,tmpl string) {
	parsedTemplate,_ := template.ParseFiles("./Views/"+tmpl)
	err:=parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("Error: ",err)
	}

}