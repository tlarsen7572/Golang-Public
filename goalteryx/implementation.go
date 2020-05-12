package main

type MyNewPlugin struct {
}

func (*MyNewPlugin) Init(toolId int, config string) bool {
	return true
}

func (*MyNewPlugin) PushAllRecords(recordLimit int) int {
	return 1
}

func (*MyNewPlugin) Close(hasErrors bool) {

}

func (*MyNewPlugin) AddIncomingConnection(connectionType string, connectionName string) IncomingInterface {
	return nil
}

func (*MyNewPlugin) AddOutgoingConnection(connectionName string) bool {
	return true
}
