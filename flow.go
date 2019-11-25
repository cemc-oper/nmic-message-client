package nmic_message_client

type Flow struct {
	MessageType   string
	Message       string
	DataFlow      string
	Send          string
	BusinessState string
}

func (f *Flow) FillMessage(m *MonitorMessage) {
	m.MessageType = f.MessageType
	m.Message = f.Message
	m.Fields.DataFlow = f.DataFlow
	m.Fields.Send = f.Send
	m.Fields.BusinessState = f.BusinessState
}

const (
	FlowMessageType = "RT.DPC.STATION.DI"
	SwfdpTarget     = "WMC_SWFDP"
)

func CreateFlow(
	messageType string,
	message string,
	target string,
	state string) Flow {
	return Flow{
		MessageType:   messageType,
		Message:       message,
		DataFlow:      "BDMAIN",
		Send:          target,
		BusinessState: state,
	}
}
