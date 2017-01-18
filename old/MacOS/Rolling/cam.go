package main

import "github.com/lazywei/go-opencv/opencv"
import jpeg "image/jpeg"
import "bytes"

func GetPic() []byte {
	//Shows green led
	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("can not open camera")
	}
	img := cap.RetrieveFrame(1)
	cap.Release()
	pngimg := img.ToImage();
	//fout, _ := os.Create("/Users/Mac/Desktop/asdf.jpg")

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, pngimg, nil);
	byteImg := buf.Bytes()
	return byteImg
}

	//Read
	/*out, _ := os.Create("/Users/Mac/Desktop/a.jpeg")
	out.Write(img)*/
