package health

import (
	"mirasynth.stream/notifier/internal/server/health/live"
	"mirasynth.stream/notifier/internal/server/health/ready"
)

func RegisterController() (*live.HealthLiveController, *ready.HealthReadyController) {
	hlc := live.RegisterController("/api/v1/health/live")
	hrc := ready.RegisterController("/api/v1/health/reaady")

	return hlc, hrc
}
