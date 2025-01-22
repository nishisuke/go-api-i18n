package i18n

import "context"

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
