package main

import (
	"bytes"
	"image/png"
	"testing"
  "errors"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte)(int, error){
  return 0, errors.New("Expected error")
}

// func TestGenerateQRCodeReturnsValue(t *testing.T) {
//   buffer := new(bytes.Buffer)
// 	GenerateQRCode(buffer, "555-2368")
//
// 	if buffer.Len() == 0 {
// 		t.Errorf("No QRCode Generated")
// 	}
// }

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
  buffer := new(bytes.Buffer)
  GenerateQRCode(buffer,"555-2368")

  if buffer.Len() == 0 {
    t.Errorf("No QRCode Generated")
  }
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not PNG: %s", err)
	}
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
  buffer := new(ErrorWriter)
  err := GenerateQRCode(buffer,"555-2368")

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}
}
