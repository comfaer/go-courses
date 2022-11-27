package converter

import "strings"

type converter struct {
	symbols []string
	values  []int
}

type Converter interface {
	IntToRoman(number int) string
}

func (conv *converter) IntToRoman(number int) string {
	var result strings.Builder

	for i, val := range conv.values {
		for number >= val {
			result.WriteString(conv.symbols[i])
			number -= val
		}
	}
	return result.String()
}

func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		values:  values,
		symbols: symbols,
	}
}
