package main

import (  
    "fmt"
    "os"
    "newgraph/logic"
    "github.com/chewxy/math32"
)

const WIDTH, HEIGHT = 640, 480;
const INVWIDTH, INVHEIGHT = 1 / float32(WIDTH), 1 / float32(HEIGHT)
const FOV, ASPECTRATIO = 30., float32(WIDTH) / float32(HEIGHT)


func main() {
    img, _ := os.Create("img.ppm")
    
    var shapes []logic.Shape

    // position, radius, surface color, reflectivity, transparency, emission color
    shapes = append(shapes, 
      logic.CtSphere(logic.CtVec(0.0, -10004, -20), 10000, logic.CtVec(0.20, 0.20, 0.20), 0, 0.0, logic.CtVec(0.,0.,0.)),
      logic.CtSphere(logic.CtVec(0.0, 0, -20), 4, logic.CtVec(1.00, 0.32, 0.36), 1, 0.5, logic.CtVec(0.,0.,0.)),
      logic.CtSphere(logic.CtVec(5.0, -1, -15), 2, logic.CtVec(0.90, 0.76, 0.46), 1, 0.0, logic.CtVec(0.,0.,0.)),
      logic.CtSphere(logic.CtVec(5.0, 0, -25), 3, logic.CtVec(0.65, 0.77, 0.97), 1, 0.0, logic.CtVec(0.,0.,0.)),
      logic.CtSphere(logic.CtVec(-5.5, 0, -15), 3, logic.CtVec(0.90, 0.90, 0.90), 1, 0.0, logic.CtVec(0.,0.,0.)),
      logic.CtSphere(logic.CtVec(0.0, 20, -30), 3, logic.CtVec(0.00, 0.00, 0.00), 0, 0.0, logic.CtVec(3.,3.,3.)));

    fmt.Printf("%+v\n", shapes)
    //please panic if you die
    metadata := fmt.Sprintf("P6\n%d %d\n255\n", WIDTH, HEIGHT)
    img.WriteString(metadata)
    var pixels [WIDTH * HEIGHT * 3]byte

    offset := 0
    angle := math32.Tan(math32.Pi * 0.5 * FOV / 180.)
    for x := 0; x < HEIGHT * WIDTH; x++ {
        xx := (2. * ((float32(x % WIDTH) + 0.5) * INVWIDTH) - 1.) * angle * ASPECTRATIO
        yy := (1. - 2. * ((float32(x / WIDTH) + 0.5) * INVHEIGHT)) * angle
        raydir := logic.CtVec(xx, yy, -1.).Norm()
        colors := logic.Trace(logic.CtVec(0.,0.,0.), raydir, shapes, 0).Scale(255.)

        pixels[offset], pixels[offset + 1], pixels[offset + 2] = colors.IVals()
        offset += 3
    }

    img.Write(pixels[:]);
    defer img.Close()
}
