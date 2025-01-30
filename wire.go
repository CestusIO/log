//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package log

import "code.cestus.io/tools/wire"

// ZapperLogProviderSet provides a zap logger
var ZapperLogProviderSet = wire.NewSet(
	ProvideZaprLogger,
)
