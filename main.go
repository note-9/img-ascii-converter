package main

import (
    "fmt"
    "image"
	"image/color"
    _ "image/jpeg"
    _ "image/png"
    "os"
	"github.com/nfnt/resize"
	"golang.org/x/term"
)

func main() {
    filePath := "/home/note/Pictures/Wallpaper/wallhaven-6dr6rx_1920x1080.png"
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Error opening file %v\n", err)
        return
    }
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil{
		fmt.Printf("Error decoding image %v\n", err)
		return
	}
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println("Stdout is not a terminal.")
		return
	}

	// Get the terminal size
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Printf("Error getting terminal size: %v\n", err)
		return
	}
	
	newWidth := uint(width)
	newHeight := uint(0) 

	m := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)//lanczos3 interpolation
	
	charset := " .,-~+=oO@"

	ascii := ""

	for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
		for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			brightness := c.Y 
			
			charIndex := int(float64(brightness) / 256.0 * float64(len(charset)))
			if charIndex >= len(charset) {
				charIndex = len(charset) - 1
			}
			
			ascii += string(charset[charIndex])
		}
		ascii += "\n"
	}
	fmt.Println(ascii)
}
