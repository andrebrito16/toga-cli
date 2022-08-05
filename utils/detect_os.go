package utils

import "runtime"

func DetectOs() string {
	os := runtime.GOOS

	return os
}
