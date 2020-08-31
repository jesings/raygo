package main

import (  
    "fmt"
    "os"
    "newgraph/logic"
)

const width, height = 640, 480;
func main() {
    fmt.Println("panics when dying like it should")
    img, _ := os.Create("img.ppm")
    //please panic if you die
    metadata := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
    img.WriteString(metadata)
    for x := 0; x < width; x++ {
        var pixels [width * 3]byte
        for y := 0; y < width; y++ {
            var _ logic.Vec3
            //r
            pixels[y * 3], pixels[y * 3 + 1], pixels[y * 3 + 2] = 0, 0, 255
        }
        img.Write(pixels[:]);
    }
    defer img.Close()
}
