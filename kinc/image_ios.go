// +build ios

package kinc

func (texture *Texture) Upload(data []byte, stride int32) {
	kinc_g4_texture_upload(texture.ref, data, stride)
}
