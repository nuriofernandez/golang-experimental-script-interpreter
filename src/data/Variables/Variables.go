package Variables

import (
	"github.com/xXNurioXx/golang-experimental-script-interpreter/advicer"
	. "github.com/xXNurioXx/golang-experimental-script-interpreter/structs"
)

var Variables = []NugetVariable{}

func Get(name string) NugetVariable {
	for _, variable := range Variables {
		if variable.Name == name {
			return variable
		}
	}

	advicer.Error("Unable to retrieve a variable!")
	return NugetVariable{}
}

func Exists(name string) bool {
	for _, variable := range Variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

func Create(name string, value string) {
	if Exists(name) {
		advicer.Error("Unable to create a variable : Name already used")
		return
	}

	variable := NugetVariable{
		Type:  1,
		Name:  name,
		Value: value,
	}

	Variables = append(Variables, variable)
}
