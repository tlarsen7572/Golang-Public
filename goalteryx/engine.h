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