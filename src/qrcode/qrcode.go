package main

import (
	"fmt"
	"image"
	"image/png"
  "os"
  "io"
)

func main() {
	fmt.Println("Hello QR Code")

	// qrcode := GenerateQRCode("555-2368")
	// ioutil.WriteFile("qrcode.png", qrcode, 0644)

  file,_ := os.Create("qrcode.png")
  defer file.Close()

  GenerateQRCode(file,"555-2368")
}

func GenerateQRCode(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	png.Encode(w, img)
  return nil
}
