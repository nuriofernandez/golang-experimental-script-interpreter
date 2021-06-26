package parser

import (
	"github.com/xXNurioXx/golang-experimental-script-interpreter/parser/nugget"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/parser/nulog"
)

func Interpret(line string) {
	nugget.ReadLine(line)
	nulog.ReadLine(line)
}
