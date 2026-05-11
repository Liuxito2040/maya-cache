package backend

import (
	"fmt"
	"time"
)

func Fetch(key string) []byte {
	time.Sleep(2 * time.Second)
	return fmt.Appendf(nil, "backend_value_for_%s", key)
}
