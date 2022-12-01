package validators

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

func testRuneForRange(charcode rune, lowerLim int32, upperLim int32) bool {
	return charcode >= lowerLim && charcode <= upperLim
}

func checkIfMeetsMinimumCharcodesOnRange(sample string, lowerLim, upperLim int32, minRequirement int) bool {
	var equals int = 0
	runes := []rune(sample)
	runesLength := len(runes)
	for i := 0; i < runesLength; i++ {
		charcode := runes[i]
		if testRuneForRange(charcode, lowerLim, upperLim) {
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

func minDigit(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, 48, 57, value)
}

func minSpecialChars(sample string, value int) bool {
	var equals int = 0
	runes := []rune(sample)
	runesLength := len(runes)
	for i := 0; i < runesLength; i++ {
		charcode := runes[i]
		if !testRuneForRange(charcode, 92, 122) && !testRuneForRange(charcode, 65, 90) && !testRuneForRange(charcode, 48, 57) {
			equals++
		}
	}
	return equals >= value
}

var Validators = map[string]func(string, int) bool{
	"minSize":         minSize,
	"minUppercase":    minUppercase,
	"minLowerCase":    minLowerCase,
	"minDigit":        minDigit,
	"minSpecialChars": minSpecialChars,
}
