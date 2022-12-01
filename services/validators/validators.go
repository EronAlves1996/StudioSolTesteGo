package validators

type charType struct {
	lowerLim int32
	upperLim int32
}

var charTypeMap = map[string]charType{
	"lowerCase": {lowerLim: 97, upperLim: 122},
	"upperCase": {lowerLim: 65, upperLim: 90},
	"digit":     {lowerLim: 48, upperLim: 57},
}

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

func testRuneForRange(charcode rune, cType charType) bool {
	return charcode >= cType.lowerLim && charcode <= cType.upperLim
}

func checkIfMeetsMinimumCharcodesOnRange(sample string, cType charType, minRequirement int) bool {
	var equals int = 0
	runes := []rune(sample)
	runesLength := len(runes)
	for i := 0; i < runesLength; i++ {
		charcode := runes[i]
		if testRuneForRange(charcode, cType) {
			equals++
		}
	}
	return equals >= minRequirement
}

func minLowerCase(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, charTypeMap["lowerCase"], value)
}

func minUppercase(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, charTypeMap["upperCase"], value)
}

func minDigit(sample string, value int) bool {
	return checkIfMeetsMinimumCharcodesOnRange(sample, charTypeMap["digit"], value)
}

func minSpecialChars(sample string, value int) bool {
	var equals int = 0
	runes := []rune(sample)
	runesLength := len(runes)
	for i := 0; i < runesLength; i++ {
		charcode := runes[i]
		if !testRuneForRange(charcode, charTypeMap["lowerCase"]) &&
			!testRuneForRange(charcode, charTypeMap["upperCase"]) &&
			!testRuneForRange(charcode, charTypeMap["digit"]) {
			equals++
		}
	}
	return equals >= value
}

// Map that describes the keys that corresponds to the validator function and return a boolean
var Validators = map[string]func(string, int) bool{
	"minSize":         minSize,
	"minUppercase":    minUppercase,
	"minLowerCase":    minLowerCase,
	"minDigit":        minDigit,
	"minSpecialChars": minSpecialChars,
	"noRepeated": func(s string, i int) bool {
		return true
	},
}
