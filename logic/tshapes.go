package logic

import (
    "github.com/chewxy/math32"
)

type Sphere struct {
    center Vec3
    radius, radiusq float32
    transparency, reflection float32
    surfacecolor, emissioncolor Vec3
    //transparency, reflectivity
}

func CtSphere(center Vec3, radius float32, surfacecolor Vec3, reflection, transparency float32, emission Vec3) Sphere {
    return Sphere {center, radius, radius * radius, transparency, reflection, surfacecolor, emission}
}

type Shape interface {
    Intersect(incidence, origin Vec3, rd0, rd1 *float32) bool
    Normal(intersection Vec3) Vec3
    Transparency() float32
    Reflection() float32
    SurfaceColor() Vec3
    EmissionColor() Vec3
    Distance(o Vec3) Vec3
}

func (s Sphere)Intersect(incidence, origin Vec3, rd0, rd1 *float32) bool {
    dir := s.center.Sub(origin)
    rdiff := dir.Dot(incidence)
    if rdiff < 0.0 {
        return false
    }
    r2 := dir.LenSq() - rdiff * rdiff
    if r2 < s.radiusq {
        return false
    }
    tdist := math32.Sqrt(s.radiusq - r2)
    *rd0 = rdiff - tdist
    *rd1 = rdiff + tdist
    return true
}
func (s Sphere)Normal(intersection Vec3) Vec3 {
    return intersection.Sub(s.center)
}
func (s Sphere)Transparency() float32 {
    return s.transparency
}
func (s Sphere)Reflection() float32 {
    return s.reflection
}
func (s Sphere)SurfaceColor() Vec3 {
    return s.surfacecolor
}
func (s Sphere)EmissionColor() Vec3 {
    return s.emissioncolor
}
func (s Sphere)Distance(o Vec3) Vec3 {
    return s.center.Sub(o)
}
