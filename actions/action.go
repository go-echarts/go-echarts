package actions

type Dispatcher interface {
	getAction() Action
}

type Action interface {
	GetActionType() string
	GetConfigs() ActionConfigs
}

type ActionConfigs map[string]interface{}
