package dependencies

import (
	"runtime"
)

func GetPath() string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	case "darwin":
		return ""
	case "linux":
		return ""
	default:
		return ""
	}
}
