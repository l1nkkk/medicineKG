package graph

import (
	"encoding/json"
	"github.com/l1nkkk/medicineKG/src/util"
	"sync"
)

var OutEdge uint8 = 1
var InEdge uint8 = 2

type Property struct {
	sync.Once
	PropRawData string `json:"prop_raw_data"`
	// TODO 支持interface
	propsMapping map[string]string
}

func NewProperty(prop string) *Property {
	return &Property{
		PropRawData: prop,
	}
}

func (p *Property) Get(key string) string {
	p.Do(func() {
		p.propsMapping = make(map[string]string)
		if err := json.Unmarshal([]byte(p.PropRawData), &p.propsMapping); err != nil {
			util.Logger.Error.Println("unmarshal propRawData error: ", err.Error())
		}
	})
	if v, ok := p.propsMapping[key]; !ok {
		return ""
	} else {
		return v
	}
}

type Vertex struct {
	TagID uint32    `json:"tag_id"`
	Tag   string    `json:"tag"`
	Vid   uint64    `json:"vid"`
	Prop  *Property `json:"prop"`
}

type Edge struct {
	TagID     uint32    `json:"tag_id"`
	Tag       string    `json:"tag"`
	Vid       uint64    `json:"vid"`
	EdgeType  string    `json:"edge_type"`
	PeerTagID uint32    `json:"peer_tag_id"`
	PeerTag   string    `json:"peer_tag"`
	PeerVid   uint64    `json:"peer_vid"`
	Rank      uint64    `json:"rank"`
	Prop      *Property `json:"prop"`
}

type SubGraph struct {
	V       Vertex  `json:"v"`
	InEdge  []*Edge `json:"in_edge"`
	OutEdge []*Edge `json:"out_edge"`
}

type WriteVertex struct {
	Tag        string
	UniqueName string
	Property   string
}

type WriteEdge struct {
	Tag            string
	UniqueName     string
	Direction      uint32 // 8
	EdgeType       string
	TagPeer        string
	UniqueNamePeer string
	Property       string
}
