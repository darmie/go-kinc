package graphics4

// +build cgo

// #include "vertexstructure.h"
// static bool ToBool(int i)
// {
//  	if(i == 1) return true;
// 		return false;
// }
// static int FromBool(bool b)
// {
//  	if(b) return 1;
// 		return 0;
// }
import "unsafe"
import "C"

const (
	None       VertexData = 0
	Float1     VertexData = 1
	Float2     VertexData = 2
	Float3     VertexData = 3
	Float4     VertexData = 4
	Float4x4   VertexData = 5
	Short2Norm VertexData = 6
	Short4Norm VertexData = 7
	Color      VertexData = 8
)

type VertexData int

const MaxVertexElements = 16

type VertexElement struct {
	Name string
	Data VertexData
}

type VertexStructure struct {
	Elements  []*VertexElement
	Size      int
	Instanced bool
}

func InitVertexStructure() *VertexStructure {
	_structure = C.kinc_g4_vertex_structure_t{}

	C.kinc_g4_vertex_structure_init(unsafer.Pointer(&_structure))

	return vertexStructureFromCStruct(_structure)
}

func (s *VertexStructure) Add(name string, data VertexData) {
	_s := vertexStructureToCStruct(s)
	C.kinc_g4_vertex_structure_add(unsafe.Pointer(_s), C.CString(name), C.int(data))

	s.Size++
	s.Elements = append([]*VertexElements, &VertexElement{
		Name: name,
		Data: data,
	})
}

func vertexStructureToCStruct(structure *VertexStructure) *C.kinc_g4_vertex_structure_t {
	_structure = C.kinc_g4_vertex_structure_t{}

	for i := 0; i < MaxVertexElements; i++ {
		_structure.elements.name = C.CString(structure.Elements[i].Name)
		_structure.elements.data = C.int(structure.Elements[i].Data)
	}

	_structure.size = C.int(structure.Size)
	if structure.Instanced {
		_structure.instanced = C.ToBool(C.int(1))
	} else {
		_structure.instanced = C.ToBool(C.int(0))
	}

	return (*C.kinc_g4_vertex_structure_t)(unsafe.Pointer(&_structure))
}

func vertexStructureFromCStruct(structure *C.kinc_g4_vertex_structure_t) *VertexStructure {
	_structure := &VertexStructure{}
	_structure.Elements = make([]*VertexStructure{}, MaxVertexElements)
	for i := 0; i < MaxVertexElements; i++ {
		_structure.Elements[i].Name = C.GoString(C.CString(structure.elements[i].name))
		_structure.Elements[i].Data = (VertexData)(C.int(structure.elements[i].data))
	}
	structure.Instanced = C.FromBool(_structure.instanced) == C.int(1)

	return structure
}
