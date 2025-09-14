package comline

import (
	"encoding/base64"
)

type CustomAction struct {
	Action  string
	Payload map[string]string
}

func (c CustomAction) FormatAction() (string, map[string]string) {
	return c.Action, c.Payload
}

type RouterRegisterAction struct {
	Route       string
	Destination string
}

func (r RouterRegisterAction) FormatAction() (string, map[string]string) {
	return "router:register", map[string]string{
		"route":       r.Route,
		"destination": r.Destination,
	}
}

type ProcessReadAction struct{}

func (p ProcessReadAction) FormatAction() (string, map[string]string) {
	return "process:read", map[string]string{}
}

type RouterDeregisterAction struct {
	Route string
}

func (r RouterDeregisterAction) FormatAction() (string, map[string]string) {
	return "router:deregister", map[string]string{
		"route": r.Route,
	}
}

type ConfigReadAction struct{}

func (a ConfigReadAction) FormatAction() (string, map[string]string) {
	return "config:read", map[string]string{}
}

type ConfigReadRawAction struct{}

func (a ConfigReadRawAction) FormatAction() (string, map[string]string) {
	return "config:readraw", map[string]string{}
}

type ConfigWriteAction struct {
	Config []byte
}

func (c ConfigWriteAction) FormatAction() (string, map[string]string) {
	return "config:write", map[string]string{
		"config": base64.StdEncoding.EncodeToString(c.Config),
	}
}

type ChannelRenrollAction struct {
	Project string
}

func (c ChannelRenrollAction) FormatAction() (string, map[string]string) {
	return "channel:renroll", map[string]string{
		"project": c.Project,
	}
}

type ChannelAuthAction struct{}

func (a ChannelAuthAction) FormatAction() (string, map[string]string) {
	return "channel:auth", map[string]string{}
}

type RayReloadAction struct{}

func (a RayReloadAction) FormatAction() (string, map[string]string) {
	return "ray:reload", map[string]string{}
}

type RaySystemctlRestartAction struct{}

func (a RaySystemctlRestartAction) FormatAction() (string, map[string]string) {
	return "ray:systemctl:restart", map[string]string{}
}

type RayUpdateAction struct{}

func (a RayUpdateAction) FormatAction() (string, map[string]string) {
	return "ray:update", map[string]string{}
}
