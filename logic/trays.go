package logic

import (
    "github.com/chewxy/math32"
)

type Vec3 struct {
    x, y, z float32
}

func CtVec(x, y, z float32) Vec3 {
    return Vec3{x, y, z}
}

func (v0 Vec3) IVals() (byte, byte, byte) {
    return byte(v0.x), byte(v0.y), byte(v0.z)
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

func (v0 Vec3) Neg() Vec3 {
    return Vec3{-v0.x, -v0.y, -v0.z}
}

func (v0 Vec3) LenSq() float32 {
    return v0.Dot(v0)
}

func (v0 Vec3) Len() float32 {
    return math32.Sqrt(v0.LenSq())
}

func (v0 Vec3) Clamp(f float32) Vec3 {
    var x, y, z float32 = f, f, f
    if v0.x < x {x = v0.x}
    if v0.y < y {y = v0.y}
    if v0.z < z {z = v0.z}
    return Vec3{x, y, z}
}

func (v0 Vec3) Norm() Vec3 {
    nor := v0.LenSq()
    if(nor > 0.0) {
        inor := 1.0 / math32.Sqrt(nor)
        return v0.Scale(inor)
    }
    return v0
}

func (v0 Vec3) Cross(v1 Vec3) Vec3 {
    return Vec3{
        v0.y * v1.z - v1.y * v0.z,
        v1.x * v0.z - v0.x * v1.z,
        v0.x * v1.y - v1.x * v0.y,
    }
}

