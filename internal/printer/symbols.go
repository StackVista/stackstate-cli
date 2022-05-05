package printer

import (
	"github.com/gookit/color"
)

type Symbol struct {
	UnicodeSymbol string
	Text          string
	TextColor     string
}

var symbols = map[string]Symbol{
	"success": {UnicodeSymbol: "\u2705", Text: "[SUCCESS]", TextColor: color.Green.Render("[SUCCESS]")},  // ✅
	"warn":    {UnicodeSymbol: "\u26A0\uFE0F", Text: "[WARN]", TextColor: color.Yellow.Render("[WARN]")}, // ⚠️
	"error":   {UnicodeSymbol: "\u274C", Text: "[ERROR]", TextColor: color.Red.Render("[ERROR]")},        // ❌
	"info":    {UnicodeSymbol: color.Blue.Render("ⓘ"), Text: "info"},
}

//nolint:gocritic
func (p *StdPrinter) sprintSymbol(symbolName string) string {
	symbol := symbols[symbolName]
	if p.useColor && p.useSymbols {
		return symbol.UnicodeSymbol
	} else if p.useColor {
		return symbol.TextColor
	} else {
		return symbol.Text
	}
}
