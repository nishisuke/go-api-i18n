package i18n

import (
	"context"
	"strings"
)

const defaultLang = "ja"

var fallbackLCID = LCID{defaultLang, "JP"}

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
		lang = defaultLang
	}

	return LCID{
		Lang:    lang,
		Country: country,
	}
}

type lcidKey struct{}

func lcid(ctx context.Context) LCID {
	if lcid, ok := ctx.Value(lcidKey{}).(LCID); ok {
		return lcid
	}
	return fallbackLCID
}

func LCIDContext(ctx context.Context, lcid LCID) context.Context {
	return context.WithValue(ctx, lcidKey{}, lcid)
}
