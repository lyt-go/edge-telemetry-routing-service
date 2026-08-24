package device

import (
	"fmt"

	"edgetelemetry/internal/model"
)

type Builder struct{}

func (Builder) Build(id string, labels map[string]string) (profile *model.DeviceProfile, err error) {
	profile = &model.DeviceProfile{ID: id, Labels: make(map[string]string)}
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("build profile: %v", recovered)
		}
	}()
	for key, value := range labels {
		if key == "panic" {
			panic("invalid label source")
		}
		profile.Labels[key] = value
	}
	profile.Ready = true
	return profile, nil
}
