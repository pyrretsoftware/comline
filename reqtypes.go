package comline

type RawComData struct {
	Payload any `json:"payload,omitempty"`
	Type string `json:"type,omitempty"`
	Error string `json:"error,omitempty"`
}
type RawComRequest struct {
	Action string `json:"action"`
	Payload map[string]string `json:"payload"`
	Key string `json:"key"`
}

type RawComRayInfo struct {
	RayVersion string `json:"version"`
	ProtocolVersion string `json:"protocolVersion"`
}

type RawComKeyInfo struct {
	Holder string `json:"holder"`
	Permissions []string `json:"permissions"`
}

type RawComResponse struct {
	Ray RawComRayInfo `json:"ray"`
	Key *RawComKeyInfo `json:"key"`
	Data RawComData `json:"response"`
}