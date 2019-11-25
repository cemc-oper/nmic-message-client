package nmic_message_client

import (
	"fmt"
	"path/filepath"
	"time"
)

type ProductionInfo struct {
	ProdType         string
	AbsoluteDataName string
	FileName         string
	StartTime        time.Time
	ForecastTime     time.Duration
	DataTime         time.Time
}

func (info *ProductionInfo) FillMessage(m *MonitorMessage) {
	m.Fields.ProdType = info.ProdType
	m.Fields.AbsoluteDataName = info.AbsoluteDataName
	m.Fields.FileName = info.FileName
	m.Fields.StartTime = info.StartTime.Format("2006010215")
	m.Fields.ForecastTime = fmt.Sprintf("%03d", int(info.ForecastTime.Hours()))
	m.Fields.DataTime = info.StartTime.Format("2006-01-02 15:04:05")
}

const (
	Grib2ProductionType = "prod_grib2"
)

func CreateProductionFileInfo(
	prodType string,
	filePath string,
	startTime time.Time,
	forecastTime time.Duration,
) ProductionInfo {

	return ProductionInfo{
		ProdType:         prodType,
		AbsoluteDataName: filePath,
		FileName:         filepath.Base(filePath),
		StartTime:        startTime,
		ForecastTime:     forecastTime,
		DataTime:         startTime,
	}
}
