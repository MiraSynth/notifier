package push

import (
	"os/exec"
)

func (p *pushService) Push(data *PushData) (err error) {
	return linuxNotify(data, "normal")
}

func (p *pushService) PushCritical(data *PushData) error {
	return linuxNotify(data, "critical")
}

func linuxNotify(data *PushData, urgency string) (err error) {
	cmd := exec.Command("notify-send", "-i", data.Icon, data.Title, data.Text, "-a", data.Context, "-u", urgency)

	err = cmd.Start()
	if err != nil {
		logError(data, urgency, 1, err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		exiterr, ok := err.(*exec.ExitError)
		if ok {
			logError(data, urgency, exiterr.ExitCode(), err)
			return err
		}

		logError(data, urgency, 1, err)
	}

	return err
}
