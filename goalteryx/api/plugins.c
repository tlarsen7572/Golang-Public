#include "plugins.h"

void callEngineOutputMessage(struct EngineInterface *pEngineInterface, int toolId, int status, void * message) {
	pEngineInterface->pOutputMessage(pEngineInterface->handle, toolId, status, message);
}

long callInitOutput(struct IncomingConnectionInterface *connection, void * recordMetaInfoXml) {
    return connection->pII_Init(connection->handle, recordMetaInfoXml);
}

long callPushRecord(struct IncomingConnectionInterface *connection, void * record) {
    return connection->pII_PushRecord(connection->handle, record);
}
