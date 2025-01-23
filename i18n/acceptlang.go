package i18n

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func AcceptLanguageRequest(r *http.Request) *http.Request {
	l := parseAcceptLanguage(r.Header.Get("Accept-Language"))
	ctx := LCIDContext(r.Context(), l)
	return r.WithContext(ctx)
}

func parseAcceptLanguage(v string) LCID {
	noSpace := strings.ReplaceAll(v, " ", "")
	arr := strings.Split(noSpace, ",")

	isAcceptAny := false
	var ret LCID
	var currentQuality float64 = 0
	for _, lang := range arr {
		tagAndQuality := strings.SplitN(lang, ";", 2)
		tag := tagAndQuality[0]
		if tag == "*" {
			isAcceptAny = true
			continue
		}

		quality := 1.0
		if len(tagAndQuality) == 2 {
			q, err := parseQuality(tagAndQuality[1])
			if err != nil {
				slog.Warn("Parse Accept-Language quality failed", slog.String("val", tagAndQuality[1]), slog.String("error", err.Error()))
			} else {
				quality = q
			}
		}

		l := NewLCIDFromTag(tag)
		if _, ok := supportedLangs[l.Lang]; !ok {
			continue
		}
		if quality > currentQuality {
			ret = l
			currentQuality = quality
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

func parseQuality(v string) (float64, error) {
	q, found := strings.CutPrefix(v, "q=")
	if !found {
		return 1, nil
	}
	f, err := strconv.ParseFloat(q, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
