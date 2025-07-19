package channel

import "testing"

func TestOneProducerAndMultiReceivers(t *testing.T) {
	CloseChannelCase1AndCase2()
}

func TestMultiProducersAndOneSender(t *testing.T) {
	CloseChannelCase3()
}

func TestMultiProducersAndSenders(t *testing.T) {
	CloseChannelCase4()
}
