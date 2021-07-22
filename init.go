package tarantool

import (
	"github.com/cryptericon/decimal"
	"github.com/vmihailenco/msgpack/v5"
)

func init() {
	msgpack.RegisterExt(1, (*decimal.Decimal)(nil))
}
