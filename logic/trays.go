package logic

type Vec3 struct {
    x, y, z float32
}

func (v0 Vec3) Scale(factor float32) Vec3 {
    return Vec3{v0.x * factor, v0.y * factor, v0.z * factor}
}

func (v0 Vec3) Mul(v1 Vec3) Vec3 {
    return Vec3{v0.x * v1.x, v0.y * v1.y, v0.z * v1.z}
}

func (v0 Vec3) Dot(v1 Vec3) float32 {
    return v0.x * v1.x + v0.y * v1.y + v0.z * v1.z
}

func (v0 Vec3) Add(v1 Vec3) Vec3 {
    return Vec3{v0.x + v1.x, v0.y + v1.y, v0.z + v1.z}
}

func (v0 Vec3) Sub(v1 Vec3) Vec3 {
    return Vec3{v0.x - v1.x, v0.y - v1.y, v0.z - v1.z}
}
