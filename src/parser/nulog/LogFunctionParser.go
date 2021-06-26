package nulog

import (
	"github.com/xXNurioXx/golang-experimental-script-interpreter/advicer"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/controllers"
	"regexp"
)

func matches(line string) bool {
	match, _ := regexp.MatchString(`nulog(.+)`, line)
	return match
}

func compiles(line string) bool {
	match, _ := regexp.MatchString(`nulog\('(.+)'\);`, line)
	return match
}

func parse(line string) string {
	re := regexp.MustCompile(`nulog\('(.+)'\);`)
	data := re.FindSubmatch([]byte(line))
	return string(data[1])
}

func ReadLine(line string) bool {
	if !matches(line) {
		return false
	}

	if !compiles(line) {
		advicer.Error("Can't parse nulog: Invalid syntax")
		return false
	}

	value := parse(line)
	controllers.Log(value)
	return true
}
