package modbus

func afterWriteCoilsHook(address uint16, quality uint16, buf []byte) {

}

func beforeWriteCoilsHook(address uint16, quality uint16, buf []byte) bool {
	return true
}

func beforeWriteHoldingsBytesHook(address uint16, quality uint16, valBuf []byte) bool {
	return true
}

func afterWriteHoldingsBytesHook(address uint16, quality uint16, valBuf []byte) {

}

func afterWriteInputsHook(address uint16, buf []uint16) {

}

func beforeWriteInputsHook(address uint16, buf []uint16) bool {
	return true
}

func afterWriteInputBytesHook(address uint16, quality uint16, buf []byte) {

}

func beforeWriteInputBytesHook(address uint16, quality uint16, buf []byte) bool {

	return true
}

func afterWriteHoldingsHook(address uint16, buf []uint16) {

}

func beforeWriteHoldingsHook(address uint16, buf []uint16) bool {
	return true
}

func beforeWriteDiscretesHook(address uint16, quality uint16, buf []byte) bool {
	return true
}

func afterWriteDiscretesHook(address uint16, quality uint16, buf []byte) {

}
