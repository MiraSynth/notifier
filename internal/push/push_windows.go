package push

import (
	_ "embed"
	"fmt"
	"os/exec"
)

//go:embed assets/show-notification.ps1
var notificationPowershellFunction string

func (p *pushService) Push(data *PushData) error {
	cmd := exec.Command("powershell", "-Command", notificationPowershellFunction, fmt.Sprintf(`Show-Notification "%s" "%s" "%s" "%s"`, data.Context, data.Title, data.Text, data.Icon), "")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
	}()

	return nil
}

func (p *pushService) PushCritical(data *PushData) error {
	return nil
}
