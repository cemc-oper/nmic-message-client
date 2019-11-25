package nmic_message_client

import "time"

type MonitorMessage struct {
	MessageType string         `json:"type"`       // topic: RT.DPC.STATION.DI
	Name        string         `json:"name"`       // 消息名称: GRAPES_GFS
	Message     string         `json:"message"`    // 消息描述: GRAPES_GFS_GMF Message
	OccurTime   int64          `json:"occur_time"` // 时间戳
	Fields      ProdFileFields `json:"fields"`
}

type ProdFileFields struct {
	DataType         string `json:"DATA_TYPE"`        // 四级编码，GRAPES_GFS: DI.CODE.0001.0002
	DataType1        string `json:"DATA_TYPE_1"`      // 数据源：nwpc_grapes_gfs
	DataTime         string `json:"DATA_TIME"`        // 出产品数据的业务时次，分钟级别：2019-10-23 16:31:00
	FileName         string `json:"FILE_NAME"`        // 文件名称
	BusinessState    string `json:"BUSINESS_STATE"`   // 业务状态：1（正常）、0（错误）
	Receive          string `json:"RECEIVE"`          // 上游系统编码：GRAPES_GFS
	Send             string `json:"SEND"`             // 资料去向，存储系统标识：WMC_SWFDP
	DataFlow         string `json:"DATA_FLOW"`        // 业务流程标识：BDMAIN（大数据平台主流程）
	ProdType         string `json:"type"`             // 产品类型，GRIB2数据：prod_grib2
	AbsoluteDataName string `json:"absoluteDataName"` // 文件绝对路径
	StartTime        string `json:"start_time"`       // 起报时间,YYYYMMDDHH
	ForecastTime     string `json:"forecast_time"`    // 时效, FFF
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func CreateProdFileMessage(
	modelInfo ModelInfo,
	productionInfo ProductionInfo,
	flow Flow) MonitorMessage {
	message := MonitorMessage{
		OccurTime: makeTimestamp(),
	}

	modelInfo.FillMessage(&message)
	productionInfo.FillMessage(&message)
	flow.FillMessage(&message)

	return message
}
