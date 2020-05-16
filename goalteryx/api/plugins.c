#include "plugins.h"

void callEngineOutputMessage(struct EngineInterface *pEngineInterface, int toolId, int status, void * message) {
	pEngineInterface->pOutputMessage(pEngineInterface->handle, toolId, status, message);
}
