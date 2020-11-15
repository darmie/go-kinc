package main

import (
	"fmt"

	"github.com/darmie/go-kinc/kinc"
)

var (
	win      *kinc.Window
	pipe     *kinc.Pipeline
	vertices *kinc.VertexBuffer
	indices  *kinc.IndexBuffer
)

func main() {
	win = kinc.Init("Shader", 1024, 768, nil, nil)
	g4 := &kinc.Graphics{}
	update := func() {
		fmt.Println("update")

		g4.Begin()
		g4.Clear(kinc.CLEAR_COLOR, 0, kinc.CLEAR_DEPTH, kinc.CLEAR_STENCIL)

		g4.SetPipeline(pipe)
		g4.SetVertexBuffer(vertices)
		g4.SetIndexBuffer(indices)
		g4.DrawIndexedVertices()

		g4.End()
		g4.SwapBuffers()
	}

	kinc.SetUpdateCallback(&update)
	vertexShader := kinc.ShaderInitFromSource("Tests/Empty/Sources/shader.frag.glsl", kinc.SHADER_TYPE_VERTEX)
	fragmentShader := kinc.ShaderInitFromSource("Tests/Empty/Sources/shader.frag.glsl", kinc.SHADER_TYPE_FRAGMENT)

	structure := kinc.VertexStructureInit()
	structure.Add("pos", kinc.VERTEX_DATA_FLOAT3)

	pipe = kinc.PipelineInit()
	pipe.InputLayout[0] = structure
	pipe.InputLayout[1] = nil
	pipe.VertexShader = vertexShader
	pipe.FragmentShader = fragmentShader
	pipe.Compile()

	vertices = kinc.VertexBufferInit(3, structure, kinc.USAGE_STATIC, 0)
	defer vertices.UnLock(3)
	_v := make([]float32, 9)
	_v[0] = -1
	_v[1] = -1
	_v[2] = 0.5
	_v[3] = 1
	_v[4] = -1
	_v[5] = 0.5
	_v[6] = -1
	_v[7] = 1
	_v[8] = 0.5
	vertices.Lock(_v, 0, 3)

	indices = kinc.IndexBufferInit(3)
	defer indices.UnLock()
	_i := make([]int32, 3)
	_i[0] = 0
	_i[1] = 1
	_i[2] = 2
	indices.Lock(_i)

	kinc.Start()
}
