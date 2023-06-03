package store

import (
	"github.com/hello2dj/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
