const baseUrl = "http://127.0.0.1:9091/api"
const Api = {
    loadTableColumnApi: `${baseUrl}/v2/xsheetServer/tablemeta/get`,
    createReportApi: `${baseUrl}/v2/xsheetServer/create`,
    getReportDataApi: `${baseUrl}/v2/xsheetServer/rawdatas/get`,
};

export {Api}
