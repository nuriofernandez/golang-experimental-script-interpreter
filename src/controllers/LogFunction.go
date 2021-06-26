package controllers

import (
	"fmt"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/data/Variables"
	. "github.com/xXNurioXx/golang-experimental-script-interpreter/structs"
	"regexp"
	"strings"
)

func findVariables(message string) []NugetVariable {
	var variables []NugetVariable
	pat := regexp.MustCompile(`#\{(.[^}]+)}`)
	matches := pat.FindAllStringSubmatch(message, -1)
	for _, match := range matches {
		name := match[1]
		variables = append(variables, Variables.Get(name))
	}
	return variables
}

func replaceVariables(message string) string {
	variables := findVariables(message)
	for _, variable := range variables {
		name := fmt.Sprintf("#{%s}", variable.Name)
		message = strings.ReplaceAll(message, name, variable.Value)
	}
	return message
}

func Log(message string) {
	fmt.Println(replaceVariables(message))
}
