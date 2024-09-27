package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// global variable
var folderpath string = "picture_folder"

// converting image to GrayScale
func grayScaler(img image.Image) *image.Gray {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			R, G, B, _ := img.At(x, y).RGBA()
			// Luma: Y = 0.2126*R + 0.7152*G + 0.0722*B
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
		log.Fatalln("Failed to read directory:", err)
	}

	var varable int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1) // Increment wait group counter
		go func(file os.DirEntry) {
			defer wg.Done() // Decrement the counter when the goroutine completes

			mu.Lock()
			varable++
			imagename := "image" + strconv.Itoa(varable) + ".png"
			mu.Unlock()

			filePath := filepath.Join(folderpath, file.Name())
			imgfile, err := os.Open(filePath)
			if err != nil {
				log.Fatalf("Failed to open image path: %v", err)
			}
			defer imgfile.Close() // Close the image file

			// Decode image to JPEG
			img, err := jpeg.Decode(imgfile)
			if err != nil {
				log.Fatalf("Failed to decode JPEG image: %v", err)
			}
			log.Printf("Processing image: %s", file.Name())

			// Create output PNG file
			f, err := os.Create(imagename)
			if err != nil {
				log.Fatalf("Failed to create output file: %v", err)
			}
			defer f.Close() // Close the output file

			grayImg := grayScaler(img)
			if err := png.Encode(f, grayImg); err != nil {
				log.Fatalf("Failed to encode PNG image: %v", err)
			}
		}(file)
	}

	wg.Wait() // Wait for all goroutines to finish
}
