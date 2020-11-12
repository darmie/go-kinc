package main

import (
	"fmt"

	"github.com/darmie/go-kinc/Sources/kinc"
	g4 "github.com/darmie/go-kinc/Sources/kinc/graphics4"
)

var (
	win      *kinc.Window
	pipe     *g4.Pipeline
	vertices *g4.VertexBuffer
	indices  *g4.IndexBuffer
)

func update() {
	fmt.Println("update")

	g4.Begin(win.Index)
	g4.Clear(g4.ClearColor, 0, g4.ClearDepth, g4.ClearStencil)

	g4.SetPipeline(pipe)
	g4.SetVertexBuffer(vertices)
	g4.SetIndexBuffer(indices)
	g4.DrawIndexedVertices()

	g4.End(win.Index)
	g4.SwapBuffers()
}

func main() {
	win = kinc.Init("Shader", 1024, 768, nil, nil)

	kinc.SetUpdateCallback(update)
	vertexShader := g4.InitShader("shader.vert", g4.Vertex)
	fragmentShader := g4.InitShader("shader.frag", g4.Fragment)

	structure := g4.InitVertexStructure()
	structure.Add("pos", g4.Float3)

	pipe = g4.InitPipeline()
	pipe.InputLayout[0] = structure
	pipe.InputLayout[1] = nil
	pipe.VertexShader = vertexShader
	pipe.FragmentShader = fragmentShader
	pipe.Compile()

	vertices = g4.InitVertexBuffer(3, structure, g4.Static, 0)
	defer vertices.Unlock(3)
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

	indices = g4.InitIndexBuffer(3)
	defer indices.Unlock()
	_i := make([]float32, 3)
	_i[0] = 0
	_i[1] = 1
	_i[2] = 2
	indices.Lock(_i)

	kinc.Start()
}
