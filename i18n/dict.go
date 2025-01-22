package i18n

var enDict = map[messageKey]string{
	"greet": "Hello, %s!",
}

var jaDict = map[messageKey]string{
	"greet": "こんにちは、%s！",
}

func dict(lcid LCID) map[messageKey]string {
	switch lcid.Lang {
	case "en":
		return enDict
	default:
		return jaDict
	}
}
