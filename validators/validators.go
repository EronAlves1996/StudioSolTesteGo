package validators

func minSize(sample string, value int) bool {
	return len([]rune(sample)) >= value
}

var Validators = map[string]func(string, int) bool{
	"minSize": minSize,
}
