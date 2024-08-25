package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/color"
	"image/jpeg"
	"image/png"
)

func grayScaler(img image.Image) *image.Gray {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			R, G, B, _ := img.At(x, y).RGBA()
			//Luma: Y = 0.2126*R + 0.7152*G + 0.0722*B
			Y := (0.2126*float64(R) + 0.7152*float64(G) + 0.0722*float64(B)) * (255.0 / 65535)
			grayPix := color.Gray{uint8(Y)}
			grayImg.Set(x, y, grayPix)
		}
	}

	return grayImg
}

func main() {

	resp, err := os.Open("picture_folder/kk.jfif")
	if err != nil {
		fmt.Println("err")
	}

	//Decode image to JPEG
	img, err := jpeg.Decode(resp)
	if err != nil {
		//handling error
		log.Fatalf("Failed to decode JPEG image: %v", err)
	}
	log.Printf("Image type: %T", img)

	//converting image to GrayScale

	//Working with grayScale image, e.g conver to png
	f, err := os.Create("five_years.png")
	if err != nil {
		// handling error
		log.Fatal(err)
	}
	defer f.Close()

	grayImg := grayScaler(img)
	if err := png.Encode(f, grayImg); err != nil {
		log.Fatal(err)
	}

}
