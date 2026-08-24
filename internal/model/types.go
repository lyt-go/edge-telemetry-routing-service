package model

type Reading struct {
	DeviceID string
	Value    int
}

type DeviceProfile struct {
	ID     string
	Labels map[string]string
	Ready  bool
}

type Alert struct {
	ID      string
	Version int
	State   string
}
