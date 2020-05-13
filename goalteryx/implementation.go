package main

type MyNewPlugin struct {
	ToolId int
}

func (plugin *MyNewPlugin) Init(toolId int, config string) bool {
	plugin.ToolId = toolId
	OutputMessage(plugin.ToolId, 1, config)
	return true
}

func (plugin *MyNewPlugin) PushAllRecords(recordLimit int) int {
	return 1
}

func (plugin *MyNewPlugin) Close(hasErrors bool) {

}

func (plugin *MyNewPlugin) AddIncomingConnection(connectionType string, connectionName string) IncomingInterface {
	return nil
}

func (plugin *MyNewPlugin) AddOutgoingConnection(connectionName string) bool {
	return true
}
