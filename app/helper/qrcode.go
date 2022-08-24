package helper

import (
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type QRCodeStruct struct {
	Status string `json:status`
	QrCode string `json:qrCode`
}

func GenQR(uuidv4 string) *QRCodeStruct {
	qrcode, err := qrcode.New(uuidv4)
	if err != nil {
		SugarObj.Error(err)
		return &QRCodeStruct{
			Status: "fail",
			QrCode: uuidv4,
		}
	}
	path := "assets/" + uuidv4 + ".jpeg"
	refImage := "assets/Burger_King.png"

	standardWriter, err := standard.New(path,
		standard.WithHalftone(refImage),
		standard.WithQRWidth(40),
	)
	SugarObj.Info(standardWriter)

	if err != nil {
		SugarObj.Error(err)
		return &QRCodeStruct{
			Status: err.Error(),
			QrCode: uuidv4,
		}
	}

	err = qrcode.Save(standardWriter)
	if err != nil {
		SugarObj.Error(err)
		return &QRCodeStruct{
			Status: err.Error(),
			QrCode: uuidv4,
		}
	}
	resp := &QRCodeStruct{
		Status: "success",
		QrCode: uuidv4,
	}
	return resp
}
