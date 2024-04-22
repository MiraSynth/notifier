package push

import (
	_ "embed"
	"fmt"
	"os/exec"
)

//go:embed assets/show-notification.ps1
var notificationPowershellFunction string

func (p *pushService) Push(data *PushData) (string, error) {
	cmd := exec.Command("powershell", "-Command", notificationPowershellFunction, fmt.Sprintf(`Show-Notification "%s" "%s" "%s" "%s"`, data.Context, data.Title, data.Text, data.Icon), "")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), err
}

func (p *pushService) PushCritical(data *PushData) (string, error) {
	return "", nil
}
