package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {

	image, err := importImage("./car.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(image)
	exportImage("./rec.jpeg", image)
}

func importImage(loc string) (image.Image, error) {
	car, err := os.Open(loc)
	if err != nil {
		log.Fatal(err)
	}
	defer car.Close()

	return jpeg.Decode(car)
}

func exportImage(loc string, img image.Image) error {
	out, err := os.Create(loc)
	if err != nil {
		return err
	}
	defer out.Close()
	jpeg.Encode(out, img, nil)
	return nil

}
