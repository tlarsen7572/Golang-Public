#include "plugins.h"


extern __declspec(dllexport) long _stdcall AlteryxGoPlugin(int nToolID,
	const wchar_t *pXmlProperties,
	struct EngineInterface *pEngineInterface,
	struct PluginInterface *r_pluginInterface)
{
    r_pluginInterface->handle = GetPlugin();
    r_pluginInterface->pPI_PushAllRecords = PiPushAllRecords;
};
