package logic

type Sphere struct {
    center Vec3
    radius float32
    radiusq float32
    //surface, emission color, transparency, reflectivity
}

type Shape interface {
    Intersect(incidence origin Vec3) bool
}

func (s Sphere)Intersect(incidence origin Vec3) bool {
    dir := s.center.Sub(origin)
    rdiff := dir.Dot(incidence)
    if rdiff < 0.0 {
        return false
    }
    r2 := dir.LenSq() - rdiff * rdiff
    if r2 < radiusq {
        return false
    }
    //calculate outbound ray
    return true
}
