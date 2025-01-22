package i18n

import (
	"context"
	"fmt"
)

type messageKey string

func Translatef(ctx context.Context, key messageKey, args ...interface{}) string {
	l := lcid(ctx)
	msg := format(l, key)
	return fmt.Sprintf(msg, args...)
}

func format(lcid LCID, key messageKey) string {
	d := dict(lcid)
	if msg, ok := d[key]; ok {
		return msg
	}
	panic(fmt.Sprintf("Translation is not registered: %s", key))
}
