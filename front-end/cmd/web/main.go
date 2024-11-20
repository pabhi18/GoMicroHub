package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	basePath, err := getBashPath()
	if err != nil {
		panic(err)
	}
	partials := []string{
		filepath.Join(basePath, "templates/base.layout.gohtml"),
		filepath.Join(basePath, "templates/header.partial.gohtml"),
		filepath.Join(basePath, "templates/footer.partial.gohtml"),
	}

	var templateSlice []string
	testPageFile := filepath.Join(basePath, fmt.Sprintf("templates/%s", t))
	templateSlice = append(templateSlice, testPageFile)

	for _, x := range partials {

		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBashPath() (string, error) {
	basePath, err := os.Getwd()
	if err != nil {
		return "", errors.New(err.Error())
	}
	fmt.Println("Base Path:", basePath) // Debug print
	return basePath, nil
}
