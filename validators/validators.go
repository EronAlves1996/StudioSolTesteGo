package validators

import "strings"

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

func minUppercase(sample string, value int) bool {
	var equals int = 0
	for i := 0; i < len(strings.Split(sample, "")); i++ {
		charcode := []rune(sample)[i]
		if charcode >= 65 && charcode <= 90 {
			equals++
		}
	}
	return value <= equals
}

var Validators = map[string]func(string, int) bool{
	"minSize":      minSize,
	"minUppercase": minUppercase,
}
