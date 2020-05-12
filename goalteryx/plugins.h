#include <stdbool.h>
#include <stddef.h>

// Plugin definitions

struct RecordData
{

};

typedef long ( _stdcall * T_II_Init)(void * handle, void * pXmlRecordMetaInfo);
typedef long ( _stdcall * T_II_PushRecord)(void * handle, const struct RecordData * pRecord);
typedef void ( _stdcall * T_II_UpdateProgress)(void * handle, double dPercent);
typedef void ( _stdcall * T_II_Close)(void * handle);
typedef void ( _stdcall * T_II_Free)(void * handle);

struct IncomingConnectionInterface
{
	int sizeof_IncomingConnectionInterface;
	void * handle;
	T_II_Init			pII_Init;
	T_II_PushRecord		pII_PushRecord;
	T_II_UpdateProgress pII_UpdateProgress;
	T_II_Close			pII_Close;
	T_II_Free			pII_Free;
};

typedef void ( _stdcall * T_PI_Close)(void * handle, bool bHasErrors);
typedef long ( _stdcall * T_PI_PushAllRecords)(void * handle, __int64 nRecordLimit);
typedef long ( _stdcall * T_PI_AddIncomingConnection)(void * handle,
    void * pIncomingConnectionType,
    void * pIncomingConnectionName,
    struct IncomingConnectionInterface *r_IncConnInt);
typedef long ( _stdcall * T_PI_AddOutgoingConnection)(void * handle,
              void * pOutgoingConnectionName,
              struct IncomingConnectionInterface *pIncConnInt);

struct PluginInterface
{
	int								sizeof_PluginInterface;
	void *							handle;
	T_PI_Close						pPI_Close;
	T_PI_AddIncomingConnection		pPI_AddIncomingConnection;
	T_PI_AddOutgoingConnection		pPI_AddOutgoingConnection;
	T_PI_PushAllRecords				pPI_PushAllRecords;
};

// Engine definitions

typedef void AlteryxThreadProc(void *pData);
struct PreSortConnectionInterface;
struct EngineInterface {
    const wchar_t (*CreateTempFileName)(const wchar_t *pExt);
    const wchar_t (*CreateTempFileName2)(const wchar_t *pExt, int nOptions);
    const wchar_t (*GetInitVar)(const wchar_t *pVar);
    int (*IsLicensed)(const wchar_t *pDll, const wchar_t *pEntryPoint);
    long (*OutputMessage)(int nToolID, int nStatus, const wchar_t *pMessage);
    long (*OutputToolProgress)(int nToolID, double dPercentProgress);
    long (*PreSort)(int nToolID, const wchar_t *pSortInfo, struct IncomingConnectionInterface *pOrigIncConnInt, struct IncomingConnectionInterface ** r_ppNewIncConnInt, struct PreSortConnectionInterface ** r_ppPreSortConnInt);
    void (*QueueThread)(AlteryxThreadProc pProc, void *pData);
};

// For the glue

void * GetPlugin();

long PiPushAllRecords(void * handle, __int64 recordLimit);
void PiClose(void * handle, bool hasErrors);
long PiAddIncomingConnection(void * handle, void * connectionType, void * connectionName, struct IncomingConnectionInterface * incomingInterface);
long PiAddOutgoingConnection(void * handle, void * connectionName, struct IncomingConnectionInterface * incomingInterface);
long IiInit(void * handle, void * recordInfoIn);

long __declspec(dllexport) AlteryxGoPlugin(int nToolID,
	void * pXmlProperties,
	struct EngineInterface *pEngineInterface,
	struct PluginInterface *r_pluginInterface);