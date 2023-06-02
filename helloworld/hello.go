package helloworld

import "fmt"

const (
	Spanish = "Spanish"
	French  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s, %s", getPrefix(language), name)
}

func getPrefix(language string) string {
	switch language {
	case Spanish:
		return "Hola"
	case French:
		return "Bonjour"
	default:
		return "Hello"
	}
}
