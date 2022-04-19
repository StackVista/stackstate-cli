package printer

import (
	"strings"

	"github.com/gookit/color"
)

type Symbol struct {
	UnicodeChar string
	Fallback    string
}

var symbols = map[string]Symbol{
	"success": {UnicodeChar: "\u2705", Fallback: "success"},    // ✅
	"warn":    {UnicodeChar: "\u26A0\uFE0F", Fallback: "warn"}, // ⚠️
	"error":   {UnicodeChar: "\u274C", Fallback: "error"},      // ❌
	"info":    {UnicodeChar: color.Blue.Render("ⓘ"), Fallback: "info"},
}

func (p *StdPrinter) sprintSymbol(symbolName string) string {
	symbol := symbols[symbolName]
	if p.useColor {
		return symbol.UnicodeChar
	} else {
		return "[" + strings.ToUpper(symbol.Fallback) + "]"
	}
}
