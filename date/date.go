package date

import (
	"time"
)

// Date vars
var (
	AsiaJakartaTZ *time.Location
)

// Date constants
const (
	RFC3339UnixMillisecond string = "2006-01-02T15:04:05.000Z07:00"
)

func init() {
	AsiaJakartaTZ, _ = time.LoadLocation("Asia/Jakarta")
}
