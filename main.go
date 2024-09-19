package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var base_directories = []string{"/cmd", "/internal", "/pkg", "/config", "/scripts", "/docs", "/migrations", "/test"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Project name?")
	project_name, err := reader.ReadString('\n')
	check(err)
	fmt.Println("Is this an api application? (y,n)")
	is_api, err := reader.ReadString('\n')
	check(err)

	project_name = strings.TrimSpace(project_name)
	is_api = strings.TrimSpace(is_api)

	project_name = strings.TrimSuffix(project_name, "\n")
	os.Mkdir(project_name, 0755)
	if is_api == "y" {
		base_directories = append(base_directories, "/api")
	}
	for _, dir := range base_directories {
		err = os.Mkdir(project_name+dir, 0755)
		check(err)
	}

	os.Mkdir(project_name+"/cmd/"+project_name+"/", 0755)
	err = os.WriteFile(project_name+"/cmd/"+project_name+"/"+"main.go", []byte("package main\n\nfunc main(){\n\n}"), 0644)
	check(err)
	err = os.WriteFile(project_name+"/cmd/"+project_name+"/"+"go.mod", []byte("module "+project_name+"\n\n"+"go 1.20"), 0644)
	check(err)
}
