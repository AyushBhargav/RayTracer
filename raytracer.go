package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"ayushbhargav.github.com/go-raytracer/math3"
)

const (
	width          = 1024
	height         = 768
	outputFilename = "out.ppm"
)

func main() {
	render()
}

func render() {
	framebuffer := make([]math3.Vec, width*height)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			framebuffer[i+j*width] = math3.Vec{
				X: float64(j) / height, Y: float64(i) / width, Z: 0.0,
			}
		}
	}

	// Write to output file
	f, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	writeIfFine(writer, fmt.Sprintf("P3\n%d %d\n255\n", width, height))
	for i := 0; i < height*width; i++ {
		for j := 0; j < 3; j++ {
			b := int(255 * math.Max(0.0, math.Min(1.0, framebuffer[i].Get(j))))
			writeIfFine(writer, fmt.Sprintf("%d ", b))
		}
		writeIfFine(writer, "\n")
	}
	writer.Flush()
}

func writeIfFine(writer *bufio.Writer, s string) {
	_, err := writer.WriteString(s)
	if err != nil {
		panic(err)
	}
}
