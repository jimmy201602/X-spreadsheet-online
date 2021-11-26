package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"sheetServerApi/global"
	"sheetServerApi/internal/middlewares/response"
	"sheetServerApi/internal/model/params"
	"sheetServerApi/internal/services"
)

// 生成报表
func CreateSheetDatas(c *gin.Context) {
	var req params.SheetParamsReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 400, "参数绑定错误", err)
		return
	}
	url, err := services.GenerateSheetFile(req)
	if err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 500, err.Error(), err)
		return
	}
	response.ResponseSuccess(c, 200, "生成成功", url)
}

// 根据ID获取表的模版数据
func GetTableRawData(c *gin.Context) {
	var req params.SheetRawdataReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 400, "参数绑定错误", err)
		return
	}
	data, err := services.GetExcelRawDatas(req.ID)
	if err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 500, err.Error(), err)
		return
	}
	response.ResponseSuccess(c, 200, "获取成功", data)
}

// 获取数据源的字段信息
func GetSheetTableMeta(c *gin.Context) {
	var req params.SheetTableReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 400, "参数绑定错误", err)
		return
	}
	data, err := services.GetTableMetaInfo(req.TableName)
	if err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 500, "生成失败", err)
		return
	}
	response.ResponseSuccess(c, 200, "生成成功", data)
}

// 获取表格的操作历史
func GetSheetHistory(c *gin.Context) {
	var req params.SheetHistoryReq
	if err := c.BindJSON(&req); err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 400, "参数绑定错误", err)
		return
	}
	data, err := services.GetSheetHistory(req)
	if err != nil {
		logrus.Fatal(err)
		response.ResponseError(c, 500, "获取失败", err)
		return
	}
	response.ResponseSuccess(c, 200, "获取成功", data)
}

func ReportDownload(c *gin.Context) {
	file_name := c.Param("file")
	if file_name == "" {
		c.JSON(404, gin.H{"code": "REPORT_NOT_FOUND", "message": "Report record not found"})
		return
	}
	if _, err := os.Stat(path.Join(global.AppSetting.ExcelFileDir, file_name)); err != nil {
		c.JSON(404, gin.H{"code": "REPORT_NOT_FOUND", "message": "Report record not found"})
		return
	}

	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + file_name}
	http.ServeFile(c.Writer, c.Request, path.Join(global.AppSetting.ExcelFileDir, file_name))

}
