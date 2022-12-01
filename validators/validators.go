package validators

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

func checkIfMeetsMinimumCharcodesOnRange(sample string, lowerLim, upperLim int32, minRequirement int) bool {
	var equals int = 0
	runes := []rune(sample)
	for i := 0; i < len(runes); i++ {
		charcode := runes[i]
		if charcode >= lowerLim && charcode <= upperLim {
			equals++
		}
	}
	return equals >= minRequirement
}

func minLowerCase(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, 97, 122, value)
}

func minUppercase(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, 65, 90, value)
}

var Validators = map[string]func(string, int) bool{
	"minSize":      minSize,
	"minUppercase": minUppercase,
	"minLowerCase": minLowerCase,
}
