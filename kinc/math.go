package kinc

type Matrix4x4 struct {
	M   [16]float32
	ref *kinc_matrix4x4
}

type Matrix3x3 struct {
	M   [9]float32
	ref *kinc_matrix3x3
}

func (mat *Matrix3x3) Get(x int32, y int32) float32 {
	return kinc_matrix3x3_get(mat.ref, x, y)
}

func (mat *Matrix3x3) Set(x int32, y int32, value float32) {
	kinc_matrix3x3_set(mat.ref, x, y, value)
	mat.M = mat.ref.m
}

func (mat *Matrix3x3) Transpose() {
	kinc_matrix3x3_transpose(mat.ref)
}

func Matrix3x3Identity() *Matrix3x3 {
	mat := &Matrix3x3{
		ref: &kinc_matrix3x3{},
	}
	*mat.ref = kinc_matrix3x3_identity()
	mat.M = mat.ref.m

	return mat
}

func Matrix3x3RotaionX(alpha float32) *Matrix3x3 {
	mat := &Matrix3x3{
		ref: &kinc_matrix3x3{},
	}
	*mat.ref = kinc_matrix3x_rotation_x(alpha)
	mat.M = mat.ref.m

	return mat
}

func Matrix3x3RotaionY(alpha float32) *Matrix3x3 {
	mat := &Matrix3x3{
		ref: &kinc_matrix3x3{},
	}
	*mat.ref = kinc_matrix3x_rotation_y(alpha)
	mat.M = mat.ref.m

	return mat
}

func Matrix3x3RotaionZ(alpha float32) *Matrix3x3 {
	mat := &Matrix3x3{
		ref: &kinc_matrix3x3{},
	}
	*mat.ref = kinc_matrix3x_rotation_z(alpha)
	mat.M = mat.ref.m

	return mat
}

func (mat *Matrix4x4) Get(x int32, y int32) float32 {
	return kinc_matrix4x4_get(mat.ref, x, y)
}

func (mat *Matrix4x4) Transpose() {
	kinc_matrix4x4_transpose(mat.ref)
}

type Float32x4 kinc_float32x4

func Float32x4Load(a float32, b float32, c float32, d float32) Float32x4 {
	return Float32x4(kinc_float32x4_load(a, b, c, d))
}

func Float32x4LoadAll(t float32) Float32x4 {
	return Float32x4(kinc_float32x4_load_all(t))
}

func (t Float32x4) Get(index int32) float32 {
	return kinc_float32x4_get(kinc_float32x4(t), index)
}
