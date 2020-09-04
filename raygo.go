package main

import (  
    "fmt"
    "os"
    "newgraph/logic"
)

const WIDTH, HEIGHT = 640, 480;
func main() {
    fmt.Println("panics when dying like it should")
    img, _ := os.Create("img.ppm")
    //please panic if you die
    metadata := fmt.Sprintf("P6\n%d %d\n255\n", WIDTH, HEIGHT)
    img.WriteString(metadata)
    var pixels [WIDTH * HEIGHT * 3]byte
    offset := 0
    for x := 0; x < HEIGHT * WIDTH; x++ {
        var _ logic.Vec3
        //r
        pixels[offset], pixels[offset + 1], pixels[offset + 2] = 0, 0, 255
        offset += 3
    }
    img.Write(pixels[:]);
    defer img.Close()
}
