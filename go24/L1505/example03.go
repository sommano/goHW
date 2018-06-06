package example03

func translates(s string) string {
	switch s {
	case "en-US":
		return "Hello "
	case "fr-FR":
		return "Bonjour "
	case "it-IT":
		return "Ciao "
	default:
		return "hello "
	}
}

func Greeting(name, locale string) string {
	salutation := translates(locale)
	return (salutation + name)
}
