package logic

import (
    "github.com/chewxy/math32"
)

const TRACEDEPTH = 5

func Trace(incidence, origin Vec3, shapes []Shape, depth int32) Vec3 {
    infconst := math32.Inf(1)
    prox := infconst
    var closeShape Shape = nil
    for _, shape := range shapes {
        rd0, rd1 := infconst, infconst
        if shape.Intersect(incidence, origin, &rd0, &rd1) {
            if rd0 < 0.0 {
                rd0 = rd1
            }
            if rd0 < prox {
                prox = rd0
                closeShape = shape
            }
        }
    }
    if closeShape == nil {
        return Vec3{0.,0.,0.}/*background color*/
    }
    intersection := origin.Add(incidence.Scale(prox))
    inorm := closeShape.Normal(intersection).Norm()
    //var bias float32 = 1e-4
    internal := incidence.Dot(inorm) > 0.0
    if internal {
      inorm = inorm.Neg()
    }
    if (closeShape.Transparency() > 0.0) || (closeShape.Reflection() > 0.0) {
      fratio := incidence.Dot(inorm)
      //fresnel := 
      _ = fratio
    }

    /*normal needs to depend on shape, currently doesn't*/
    return Vec3{1.,1.,1.}/*background color*/
}
