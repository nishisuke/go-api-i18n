package i18n

import (
	"net/http"
	"strings"
)

func FirstAcceptLanguageRequest(r *http.Request) *http.Request {
	s := parseFirstAcceptLanguageTag(r.Header.Get("Accept-Language"))
	if s == "" {
		return r
	}

	l := NewLCIDFromTag(s)
	ctx := LCIDContext(r.Context(), l)
	return r.WithContext(ctx)
}

func parseFirstAcceptLanguageTag(v string) string {
	noSpace := strings.ReplaceAll(v, " ", "")
	arr := strings.Split(noSpace, ",")

	var ret string
	for _, lang := range arr {
		l := strings.SplitN(lang, ";", 2)[0]
		if l != "*" {
			ret = l
			break
		}
	}

	if ret == "" {
		return defaultLang
	}
	return ret
}
