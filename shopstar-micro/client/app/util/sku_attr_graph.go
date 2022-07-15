package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/vmihailenco/msgpack/v5"

)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type Graph struct {
	size 		 int // 顶点数据量

	vertex       []*Node // 存储顶点数据
	vertexOffset int     //已添加的顶点的偏移量
	vertexStr    map[int]string
	vertexPos    map[string]int // 记录无序的顶点数据中元素的位置

	attrs 		 map[int64]*GrapAtt // 存储顶点数据的同级
	multiAtts 	 map[string][]string

	matrix 		[]byte // 存储邻接矩阵的数据
}

type GrapAttrVal struct {
	Super  int64
	Title  string
	Ids    string
	Status byte
}

type GrapAtt struct {
	Title string
	Attrs map[string]*GrapAttrVal
}

type Node struct {
	Super     int64
	IsVisited bool

	data    string
	isTypes bool
}

func NewGraph(size int) (g *Graph) {
	g = &Graph{
		size:         size,
		vertexOffset: 0,
	}

	g.vertex = make([]*Node, size)
	g.vertexPos = make(map[string]int, size)
	g.vertexStr = make(map[int]string, size)

	g.attrs = make(map[int64]*GrapAtt)
	g.multiAtts = make(map[string][]string)

	g.matrix = make([]byte, size*size)


	return
}


// 添加属性与属性值
func (g *Graph) AddAttrs(attr string, attrVal *GrapAttrVal) {
	attrVal.Status = 0

	if g.attrs[attrVal.Super] == nil {
		g.attrs[attrVal.Super] = &GrapAtt{
			Title: attr,
		}
		g.attrs[attrVal.Super].Attrs = make(map[string]*GrapAttrVal)
	}

	g.attrs[attrVal.Super].Attrs[attrVal.Ids] = attrVal
}
// 记录顶点节点
func (g *Graph) AddNode(super int64, ids string, data string) {
	g.vertexPos[ids] = g.vertexOffset // 是记录属性的 row 或者 col

	g.vertexStr[g.vertexOffset] = ids

	g.vertex[g.vertexOffset] = &Node{ // 是记录属性的具体信息
		data:      data,
		isTypes:   false,
		IsVisited: false,
		Super:     super,
	}

	g.vertexOffset++
}
// 地图标记
func (g *Graph) AddVertex(val []string) error {
	if len(val) < 1 {
		return errors.New("数据不足无法定位")
	}

	for i := 0; i < len(val); i++ { // row
		row := g.vertexPos[val[i]]

		g.sameLevel(val[i], row)
		// 同级的添加
		for j := i + 1; j < len(val); j++ { // col
			if err := g.matrixs(row, g.vertexPos[val[j]]); err != nil {
				return err
			}
		}
	}

	if len(val) > 2{
		g.addmultiAtts(val)
	}

	return nil
}

func (g *Graph) addmultiAtts(val []string)  {
	var key strings.Builder
	key.Grow(len(val))
	for i := 2; i < len(val) ; i++ {

		for j := 0; j < len(val) - 1; j++ {

			if j + i > len(val) {
				return
			}

			for _, v := range val[j : j + i] {
				key.WriteString(v)
			}

			if tmp := val[j + i:]; len(tmp) > 0 {
				g.multiAtts[key.String()] = append(g.multiAtts[key.String()], tmp...)
			}

			if tmp := val[:j]; len(tmp) > 0 {
				g.multiAtts[key.String()] = append(g.multiAtts[key.String()], tmp...)
			}

			key.Reset()
		}
	}
}
// 同级登记
func (g *Graph) sameLevel(item string, row int) error  {
	if g.vertex[row].isTypes {
		return nil
	}
	// 获取同级->标记
	for _, v := range g.attrs[g.vertex[row].Super].Attrs {
		if v.Ids == item {
			continue
		}

		if err := g.matrixs(row, g.vertexPos[v.Ids]); err != nil {
			return err
		}
	}

	g.vertex[row].isTypes = true

	return nil
}

func (g *Graph) matrixValue(row, col int) byte {
	return g.matrix[row*g.size+col]
}
// 对称标记
func (g *Graph) matrixs(row, col int) error {

	if row < 0 || row >= g.size || col < 0 || col > g.size {
		return errors.New("数据异常，无法定位")
	}

	g.matrix[row*g.size+col] = 1

	g.matrix[col*g.size+row] = 1

	return nil
}
// 属性的筛选；分为多属性
func (g *Graph) Collection(val ...string) (types map[int64]*GrapAtt) {
	// 可选和不可选
	types = g.attrs
	// 这是在没有，
	if len(val) == 1 && val[0] == "" {
		for i, attr := range types {
			for k, _ := range attr.Attrs {
				types[i].Attrs[k].Status = 1
			}
		}
		return
	}
	// 这是多个类型选中的时候
	if len(val) > 1 {

		var key strings.Builder

		for _, v := range val {
			types[g.vertex[ g.vertexPos[v]].Super].Attrs[v].Status = 1
			key.WriteString(v)
		}

		for _, v := range g.multiAtts[key.String()] {
			types[g.vertex[g.vertexPos[v]].Super].Attrs[v].Status = 1
		}

		return
	}
	// 这是只选中一个元素的时候处理
	vertexs := g.vertexStr // 减少运算量

	col := g.vertexPos[val[0]]
	types[g.vertex[col].Super].Attrs[vertexs[col]].Status = 1

	g.vertex[col].IsVisited = true

	for i, _ := range vertexs {
		if g.vertex[i].IsVisited {
			continue
		}

		if _, ok := vertexs[i]; ok && g.matrixValue(i, col) == 0 {
			delete(vertexs, i)
		} else {
			types[g.vertex[i].Super].Attrs[vertexs[i]].Status = 1
		}
	}

	g.vertex[col].IsVisited = false

	return
}

// 打印
func (g *Graph) Print() {

	fmt.Print("\t")
	for i := 0; i < g.size; i++ {
		fmt.Print(g.vertex[i].data, "\t")
	}

	fmt.Println()

	for i := 0; i < g.size; i++ {
		fmt.Print(g.vertex[i].data, "\t")

		for j := 0; j < g.size; j++ {
			fmt.Print(g.matrix[i*g.size+j], "\t")
		}
		fmt.Println()
	}
}

type CopyGrap struct {
	Size int // 顶点数据量

	Vertex       []*Node // 存储顶点数据
	VertexStr    map[int]string
	VertexPos    map[string]int // 记录无序的顶点数据中元素的位置

	Attrs map[int64]*GrapAtt // 存储顶点数据的同级
	MultiAtts map[string][]string

	Matrix []byte // 存储邻接矩阵的数据
}
func newCopyGraph(g *Graph) *CopyGrap {
	return &CopyGrap{
		Size:      g.size,
		Vertex:    g.vertex,
		VertexStr: g.vertexStr,
		VertexPos: g.vertexPos,
		Attrs:     g.attrs,
		MultiAtts: g.multiAtts,
		Matrix:    g.matrix,
	}
}
func (g *CopyGrap) newGraph() *Graph {
	return &Graph{
		size:         g.Size,
		vertex:       g.Vertex,
		vertexStr:    g.VertexStr,
		vertexPos:    g.VertexPos,
		attrs:        g.Attrs,
		multiAtts:    g.MultiAtts,
		matrix:       g.Matrix,
	}
}

func (g *Graph) ExportJson() string {
	d, _ := json.Marshal(newCopyGraph(g))
	return string(d)
}
func LoadGrapJson(d []byte) (*Graph, error){
	tmp := new(CopyGrap)
	if err := json.Unmarshal(d, tmp); err != nil {
		return nil, err
	}
	return tmp.newGraph(), nil
}

func (g *Graph) ExportGob() []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer) // 创建编码器
	encoder.Encode(newCopyGraph(g))

	return buffer.Bytes()
}
func LoadGrapGob(d []byte) (*Graph, error){
	decoder := gob.NewDecoder(bytes.NewReader(d)) // 创建解析器

	tmp := new(CopyGrap)

	if err := decoder.Decode(tmp); err != nil {
		return nil, err
	}

	return tmp.newGraph(), nil
}

func (g *Graph) ExportMsgPack() []byte {
	b, _ := msgpack.Marshal(newCopyGraph(g))
	return b
}
func LoadGrapMsgPack(d []byte) (*Graph, error){
	tmp := new(CopyGrap)

	if err := msgpack.Unmarshal(d, tmp) ; err != nil {
		return nil, err
	}

	return tmp.newGraph(), nil
}