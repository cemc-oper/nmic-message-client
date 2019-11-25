package nmic_message_client

type ModelInfo struct {
	ModelName         string
	ModelDataTypeCode string
	ModelDataType     string
}

func (modelInfo *ModelInfo) FillMessage(m *MonitorMessage) {
	m.Name = modelInfo.ModelName
	m.Fields.DataType = modelInfo.ModelDataTypeCode
	m.Fields.DataType = modelInfo.ModelDataType
	m.Fields.Receive = modelInfo.ModelName
}

func CreateGrapesGfsGmfModelInfo() ModelInfo {
	return ModelInfo{
		ModelName:         "GRAPES_GFS",
		ModelDataTypeCode: "DI.CODE.0001.0002",
		ModelDataType:     "nwpc_grapes_gfs",
	}
}
