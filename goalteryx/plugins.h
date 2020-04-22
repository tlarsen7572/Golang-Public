#include <stdbool.h>

struct RecordData
{

};

typedef long ( _stdcall * T_II_Init)(void * handle, const wchar_t * pXmlRecordMetaInfo);
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
    const wchar_t *pIncomingConnectionType,
    const wchar_t *pIncomingConnectionName,
    struct IncomingConnectionInterface *r_IncConnInt);
typedef long ( _stdcall * T_PI_AddOutgoingConnection)(void * handle,
              const wchar_t *pOutgoingConnectionName,
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
