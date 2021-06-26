package parser

import (
	"github.com/xXNurioXx/simple-golang-langparser/parser/nugget"
	"github.com/xXNurioXx/simple-golang-langparser/parser/nulog"
)

func Interpret(line string) {
	nugget.ReadLine(line)
	nulog.ReadLine(line)
}
