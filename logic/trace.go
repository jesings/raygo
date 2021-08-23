package logic

import (
    "github.com/chewxy/math32"
)

const TRACEDEPTH = 5

func mix(a, b, mix float32) float32 {
    return b * mix + a * (1. - mix)
}

/*normal needs to depend on shape, currently doesn't*/
func Trace(origin, incidence Vec3, shapes []Shape, depth int32) Vec3 {
    var surfacecolor Vec3
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
    var bias float32 = 1e-4
    internal := incidence.Dot(inorm) > 0.0
    inside := false
    if internal {
      inorm = inorm.Neg()
      inside = true
    }
    if ((closeShape.Transparency() > 0.0) || (closeShape.Reflection() > 0.0)) && depth != TRACEDEPTH {
      fratio := -incidence.Dot(inorm)
      facer := 1 - fratio
      fresnel := mix(facer * facer * facer, 1, 0.1)
      reflray := incidence.Sub(inorm.Scale(2 * incidence.Dot(inorm))).Norm();
      reflected := Trace(intersection.Add(inorm.Scale(bias)), reflray, shapes, depth + 1);
      refracted := Vec3{0.,0.,0.}
      if closeShape.Transparency() > 0.0 {
        var eta float32 = 1.1
        if !inside {
          eta = 1/1.1
        }
        cosi := -intersection.Dot(reflray)
        k := 1 - eta * eta * (1 - cosi * cosi)
        refractray := reflray.Scale(eta).Add(inorm.Scale(eta * cosi - math32.Sqrt(k))).Norm()
        refracted = Trace(intersection.Sub(inorm.Scale(bias)), refractray, shapes, depth + 1)
      }
      surfacecolor = reflected.Scale(fresnel).Add(refracted.Scale((1 - fresnel) * closeShape.Transparency())).Mul(closeShape.SurfaceColor())
    } else {
      for _, shape := range shapes {
        if shape.EmissionColor().x > 0 {
          transmission := Vec3{1.,1.,1.}
          lightdir := shape.Distance(intersection).Norm()
          for _, shape2 := range shapes {
            if shape != shape2 {
              var f0, f1 float32
              if shape2.Intersect(intersection.Add(inorm.Scale(bias)), lightdir, &f0, &f1) {
                transmission = Vec3{0.,0.,0.}
                break
              }
            }
          }
          var noneg float32 = 0.
          possneg := inorm.Dot(lightdir)
          if possneg > noneg {
            noneg = possneg
          }
          surfacecolor = surfacecolor.Add(shape.SurfaceColor().Mul(transmission).Scale(possneg).Mul(shape.EmissionColor()))
        }
      }
    }

    return surfacecolor.Mul(closeShape.EmissionColor())
}
