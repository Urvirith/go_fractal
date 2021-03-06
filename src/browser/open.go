package browser

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func OpenBrowser(url string) error {
	time.Sleep(2 * time.Second)

	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
