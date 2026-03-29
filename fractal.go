package main

import (
//    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "sync"
    "math/rand"
    "time"
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

func buddhabrot(width, height int, samples int) image.Image {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    counts := make([][]int, height)
    for i := range counts {
        counts[i] = make([]int, width)
    }

    rand.Seed(time.Now().UnixNano())

    scale := 320.0
    offsetX := float64(width) / 2
    offsetY := float64(height) / 2

    for s := 0; s < samples; s++ {
        // Random point in complex plane
        cr := rand.Float64()*3.5 - 2.5
        ci := rand.Float64()*3.0 - 1.5
        c := complex(cr, ci)

        var orbit []complex128
        z := complex(0, 0)

        escaped := false
        for i := 0; i < maxIter; i++ {
            z = z*z + c
            orbit = append(orbit, z)
            if real(z)*real(z)+imag(z)*imag(z) > 4 {
                escaped = true
                break
            }
        }

        if !escaped {
            continue
        }

        // Plot orbit
        for _, z := range orbit {
            x := int(real(z)*scale + offsetX)
            y := int(imag(z)*scale + offsetY)
            if x >= 0 && x < width && y >= 0 && y < height {
                counts[y][x]++
            }
        }
    }

    // Normalize and color
    maxCount := 0
    for y := range counts {
        for x := range counts[y] {
            if counts[y][x] > maxCount {
                maxCount = counts[y][x]
            }
        }
    }

    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            c := counts[y][x]
            val := uint8(255 * float64(c) / float64(maxCount))
            img.Set(x, y, color.RGBA{val, val, val, 255})
        }
    }

    return img
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
    img := buddhabrot(1280, 1280, 5000000)
    saveImage(img, "buddhabrot.png")
    // saveImage(generateFractal(1280, 1280, complex(0, 0)), "mandelbrot.png")
}
