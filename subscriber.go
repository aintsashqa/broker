package broker

import (
	"io"
)

type Subscriber interface {
	Proceed(io.Reader)
}
