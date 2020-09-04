package logic

import (
    "github.com/chewxy/math32"
)

type Sphere struct {
    center Vec3
    radius float32
    radiusq float32
    //surface, emission color, transparency, reflectivity
}

type Shape interface {
    Intersect(incidence, origin Vec3, rd0, rd1 *float32) bool
    Normal(intersection Vec3) Vec3
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
