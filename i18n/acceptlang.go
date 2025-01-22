package i18n

import (
	"net/http"
	"strings"
)

func AcceptLanguageRequest(r *http.Request) *http.Request {
	l := parseAcceptLanguageTag(r.Header.Get("Accept-Language"))
	ctx := LCIDContext(r.Context(), l)
	return r.WithContext(ctx)
}

func parseAcceptLanguageTag(v string) LCID {
	noSpace := strings.ReplaceAll(v, " ", "")
	arr := strings.Split(noSpace, ",")

	isAcceptAny := false
	var ret LCID
	for _, lang := range arr {
		tag := strings.SplitN(lang, ";", 2)[0]

		l := NewLCIDFromTag(tag)
		if _, ok := supportedLangs[l.Lang]; ok {
			ret = l
			break
		}

		if tag == "*" {
			isAcceptAny = true
		}
	}

	if ret.Lang != "" {
		return ret
	}

	if isAcceptAny {
		return defaultLCID
	}

	return fallbackLCID
}
