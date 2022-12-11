package main

import "testing"

var testDatastream = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

func TestFindStartOfPacket(t *testing.T) {
	marker, index := FindMarker(testDatastream, 4)

	if marker != "jpqm" && index != 7 {
		t.Errorf("Failed to find marker for start of packet")
	}
}

func TestFindStartOfMessage(t *testing.T) {
	marker, index := FindMarker(testDatastream, 14)

	if marker != "qmgbljsphdztnv" && index != 19 {
		t.Errorf("Failed to find marker for start of message")
	}
}
