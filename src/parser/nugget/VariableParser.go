package nugget

import (
	"github.com/xXNurioXx/golang-experimental-script-interpreter/advicer"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/data/Variables"
	"regexp"
)

func matches(line string) bool {
	match, _ := regexp.MatchString(`nugget .+`, line)
	return match
}

func compiles(line string) bool {
	match, _ := regexp.MatchString(`nugget (\w+) » '(.+)';`, line)
	return match
}

func parse(line string) (string, string) {
	re := regexp.MustCompile(`nugget (\w+) » '(.+)';`)
	data := re.FindSubmatch([]byte(line))
	return string(data[1]), string(data[2])
}

func ReadLine(line string) bool {
	if !matches(line) {
		return false
	}

	if !compiles(line) {
		advicer.Error("Can't parse nugget: Invalid syntax")
		return false
	}

	name, value := parse(line)
	Variables.Create(name, value)
	return true
}
