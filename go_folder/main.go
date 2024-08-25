package main

import (
	"image"
	"log"
	"os"
	"strconv"

	"image/color"
	"image/jpeg"
	"image/png"
	"path/filepath"
)

// global varable
var folderpath string = "picture_folder"
var varable int = 0

// converting image to GrayScale
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

// main function
func main() {

	files, err := os.ReadDir(folderpath)
	if err != nil {
		log.Fatalf("Failed to read directory")
	}

	for _, file := range files {
		varable++
		filePaths := filepath.Join(folderpath, file.Name())

		imgfile, err := os.Open(filePaths)
		if err != nil {
			log.Fatalf("Failed to open image path: %v", err)
		}

		//Decode image to JPEG
		img, err := jpeg.Decode(imgfile)
		if err != nil {
			//handling error
			log.Fatalf("Failed to decode JPEG image: %v", err)
		}
		log.Printf("Image type: %T", img)

		//Working with grayScale image, e.g conver to png
		var imagename string = "image" + strconv.Itoa(varable) + ".png"
		f, err := os.Create(imagename)
		if err != nil {
			log.Fatal(err)
		}
		//closing files
		defer f.Close()

		grayImg := grayScaler(img)
		if err := png.Encode(f, grayImg); err != nil {
			log.Fatal(err)
		}

	}
}
