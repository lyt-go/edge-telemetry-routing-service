package device

import (
	"fmt"

	"edgetelemetry/internal/model"
)

type Builder struct{}

func (Builder) Build(id string, labels map[string]string) (profile *model.DeviceProfile, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			profile = nil
			err = fmt.Errorf("build profile: %v", recovered)
		}
	}()
	local := &model.DeviceProfile{ID: id, Labels: make(map[string]string)}
	for key, value := range labels {
		if key == "panic" {
			panic("invalid label source")
		}
		local.Labels[key] = value
	}
	local.Ready = true
	return local, nil
}
