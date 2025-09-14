package comline

type Action interface {
	//FormatAction returns the action string (eg. process:get) as well as the payload that should be used for the request
	FormatAction() (string, map[string]string)
}

type Authentication interface {
	//GetKey returns the key that should be used for the request
	GetKey() string
}

type Request struct {
	Action
	Authentication
}