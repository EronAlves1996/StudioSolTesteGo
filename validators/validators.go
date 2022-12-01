package validators

import "strings"

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

func minUppercase(sample string, value int) bool {
	toCompare := strings.Split(strings.ToUpper(sample), "")
	sampleSplitted := strings.Split(sample, "")
	var equals int = 0

	for i := 0; i < len(toCompare); i++ {
		if toCompare[i] == sampleSplitted[i] {
			equals++
		}
	}
	return value <= equals
}

var Validators = map[string]func(string, int) bool{
	"minSize":      minSize,
	"minUppercase": minUppercase,
}
