package protocol

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestGetBytesFromMsg_C2SSay(t *testing.T) {
	msg := NewMsgC2SSay(12345, General, "PlayerOne", "Hello world")

	data, err := GetBytesFromMsg(msg)
	if err != nil {
		t.Fatalf("GetBytesFromMsg: unexpected error: %v", err)
	}

	expectedLen := binary.Size(msg)
	if len(data) != expectedLen {
		t.Errorf("GetBytesFromMsg: got len %d, want %d", len(data), expectedLen)
	}

	// Helper output should match the message's own GetBytes()
	direct := msg.GetBytes()
	if len(data) != len(direct) {
		t.Errorf("GetBytesFromMsg length %d != msg.GetBytes() length %d", len(data), len(direct))
	}
	if len(data) > 0 && !reflect.DeepEqual(data, direct) {
		t.Error("GetBytesFromMsg output differs from msg.GetBytes()")
	}
}

func TestReadMsgFromBytes_C2SSay(t *testing.T) {
	msg := NewMsgC2SSay(999, Whisper, "Sender", "Secret message")
	data := msg.GetBytes()

	var decoded MsgC2SSay
	err := ReadMsgFromBytes(data, &decoded)
	if err != nil {
		t.Fatalf("ReadMsgFromBytes: unexpected error: %v", err)
	}

	if !reflect.DeepEqual(decoded, msg) {
		t.Errorf("ReadMsgFromBytes: decoded message differs from original:\n got  %+v\n want %+v", decoded, msg)
	}
}

func TestGetBytesFromMsg_ReadMsgFromBytes_RoundTrip_C2SSay(t *testing.T) {
	original := NewMsgC2SSay(42, Shout, "Shouter", "Hello everyone!")

	data, err := GetBytesFromMsg(original)
	if err != nil {
		t.Fatalf("GetBytesFromMsg: %v", err)
	}

	var decoded MsgC2SSay
	err = ReadMsgFromBytes(data, &decoded)
	if err != nil {
		t.Fatalf("ReadMsgFromBytes: %v", err)
	}

	if !reflect.DeepEqual(decoded, original) {
		t.Errorf("round-trip: decoded != original:\n decoded %+v\n original %+v", decoded, original)
	}
}

func TestReadMsgFromBytes_C2SSay_TooShortData(t *testing.T) {
	msg := NewMsgC2SSay(1, General, "A", "B")
	data := msg.GetBytes()

	// Decode with truncated data should error
	var decoded MsgC2SSay
	err := ReadMsgFromBytes(data[:len(data)/2], &decoded)
	if err == nil {
		t.Error("ReadMsgFromBytes: expected error when data is too short, got nil")
	}
}
