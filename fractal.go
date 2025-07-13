package main

import (
	"fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "sync"
)

const maxIter = 1000

func mandelbrot(c complex128, z complex128) int {
    for n := 0; n < maxIter; n++ {
        z = z*z + c
        if real(z)*real(z)+imag(z)*imag(z) > 4 { return n }
    }
    return maxIter
}

func abs(num float64) float64 {
    if num < 0 { return -num }
    return num
}

func burningship(c complex128, z complex128) int {
    zx := real(c) + real(z)
    zy := imag(c) + imag(z)
    for n := 0; n < maxIter; n++ {
        if zx*zx+zy*zy > 4 { return n }
        xtemp := zx*zx - zy*zy + real(c)
        zy = abs(2*zx*zy) + imag(c)
        zx = xtemp
    }
    return maxIter
}

func generateFractal(width, height int, z complex128) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    var wg sync.WaitGroup
    for y := 0; y < height; y++ {
        wg.Add(1)
        go func(y int) {
            defer wg.Done()
            for x := 0; x < width; x++ {
                coord := complex(float64(x-720)/320, float64(y-960)/320)
                iter := burningship(coord, z)
                img.Set(x, y, color.RGBA{uint8(iter * iter * iter * 0xff / maxIter), uint8(iter * 0xff / maxIter), uint8(iter * 0xff / maxIter), 0xff})
            }
        }(y)
    }
    wg.Wait()
    return img
}

func saveImage(img image.Image, filename string) {
    file, _ := os.Create(filename)
    defer file.Close()
    png.Encode(file, img)
}

func main() {
    for param := float64(-1.5); param < 3.5; param += 0.02 {
        filename := fmt.Sprintf("mandelbrot-d%.2f.png", param + 1.5)
        saveImage(generateFractal(1280, 1280, complex(0, param)), filename)
	}
}
