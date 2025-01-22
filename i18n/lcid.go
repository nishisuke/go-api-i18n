package i18n

import (
	"strings"
)

const jaLang = "ja"
const enLang = "en"
const defaultLang = jaLang
const fallbackLang = enLang

var supportedLangs = map[string]struct{}{
	jaLang: {},
	enLang: {},
}

var defaultLCID = LCID{defaultLang, "JP"}
var fallbackLCID = LCID{enLang, "US"}

type LCID struct {
	Lang    string
	Country string
}

func NewLCIDFromTag(tag string) LCID {
	arr := strings.SplitN(tag, "-", 2)
	if len(arr) == 1 {
		return NewLCID(arr[0], "")
	}

	return NewLCID(arr[0], arr[1])
}

func NewLCID(lang, country string) LCID {
	lang = strings.TrimSpace(lang)
	country = strings.TrimSpace(country)

	if lang == "" {
		lang = fallbackLang
	}

	return LCID{
		Lang:    lang,
		Country: country,
	}
}
