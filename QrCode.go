package tools

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/skip2/go-qrcode"
)

// QrCreateEncode 生成二维码返回数组
func QrCreateEncode(content string, size int) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, size)
}

// QrCreateEncodeFile 生成二维码写入文件
func QrCreateEncodeFile(content string, size int, fileName string) error {
	return qrcode.WriteFile(content, qrcode.Medium, size, fileName)
}

// QrCreateEncode2 生成二维码写入文件(窄白边)
func QrCreateEncode2(content, fileName string, size int) error {
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return err
	}
	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}
	// create the output file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// encode the barcode as png
	err = png.Encode(file, qrCode)
	if err != nil {
		return err
	}
	return nil
}
