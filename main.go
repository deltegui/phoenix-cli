package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/deltegui/locomotive-cli/store"
)

//go:generate go run ./generators/files.go

var projectName string

func main() {
	name := flag.String("name", "", "Your project name. Needed")
	projectType := flag.String("type", "api", "Project type. Can be 'mpa' 'webpack' or 'api'")
	flag.Parse()
	if len(*name) == 0 {
		log.Fatalln("Invalid project name")
	}
	projectName = *name
	os.Mkdir(projectName, os.ModePerm)
	createDefaultProject()
	switch *projectType {
	case "webpack":
		createWebpackProject()
		break
	case "api":
		createAPIProject()
		break
	case "mpa":
		createMpaProject()
		break
	default:
		log.Fatalln("Invalid project type")
	}
	log.Println("You are ready to GO! 🚂")
}

func createDefaultProject() {
	log.Println("Generating project...")
	createDir("/src")
	createDir("/src/configuration")
	createDir("/src/controllers")
	createDir("/src/domain")

	writeFile("/src/configuration/config.go", "config")

	writeFile("/src/controllers/injector.go", "injector")

	writeFile("/src/domain/error.go", "error")
	writeFile("/src/domain/gateways.go", "gateways")

	writeFile("/config.json", "configjson")
	writeFile("/logo", "logo")

	writeFile("/.gitignore", "gitignore")
}

func createMpaProject() {
	log.Println("Creating MPA project...")
	createDir("/static")
	createDir("/templates")
	createDir("/templates/errors")
	writeFile("/templates/errors/notfound.html", "notfound.html")
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "mpamain")
	writeFile("/src/controllers/error.controller.go", "errorcontroller")
}

func createWebpackProject() {
	log.Println("Creating Webpack project...")
	createDir("/static")
	createDir("/templates")
	createDir("/templates/errors")
	writeFile("/templates/errors/notfound.html", "notfound.html")
	writeFile("/static/index.js", "webpackindexjs")
	writeFile("/makefile", "webpackmakefile")
	writeFile("/main.go", "mpamain")
	writeFile("/package.json", "packagejson")
	writeFile("/webpack.config.js", "webpackconf")
	writeFile("/src/controllers/error.controller.go", "errorcontroller")
}

func createAPIProject() {
	log.Println("Creating API project...")
	writeFile("/makefile", "mpamakefile")
	writeFile("/main.go", "apimain")
	writeFile("/src/controllers/error.controller.go", "apierrorcontroller")
}

func writeFile(path, templName string) {
	output, err := os.Create(fmt.Sprintf("%s%s", projectName, path))
	if err != nil {
		log.Fatalf("Cannot create file: %s\n", err)
	}
	defer output.Close()
	tmpl := template.New("a")
	tmpl, err = tmpl.Parse(store.Get(templName))
	if err != nil {
		panic(err)
	}
	tmpl.Execute(output, projectName)
}

func createDir(path string) {
	os.Mkdir(fmt.Sprintf("%s%s", projectName, path), os.ModePerm)
}
