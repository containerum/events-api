package atomicbool

import "sync/atomic"

func Create() (func(), func() bool) {
	var flag atomic.Value
	return func() {
			flag.Store(true)
		}, func() bool {
			var v, ok = flag.Load().(bool)
			return v && ok
		}
}
