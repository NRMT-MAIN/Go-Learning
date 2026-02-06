package main

import (
	"os"
	"text/template"
)

func textTemplateExample() {
	temp1 := template.New("example")

	temp1 , err := temp1.New("example").Parse("Welcome , {{.name}}! How are you doing ?")

	if err != nil {
		panic(err)
	}

	// Alternate Way : 
	//tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	data := map[string]interface{} {
		"name" : "Nirmit" , 
	}

	err = temp1.Execute(os.Stdout , data) //Welcome , Nirmit! How are you doing ?
	if err != nil {
		panic(err)
	}

}