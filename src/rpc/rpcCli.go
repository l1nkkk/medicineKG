package rpc

import (
	"context"
	"fmt"
	"github.com/l1nkkk/medicineKG/src/graph"
	"github.com/l1nkkk/medicineKG/src/rpc/proto"
	"github.com/l1nkkk/medicineKG/src/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GrpcManger struct {
	taiKouServer proto.TaiKouClient
	conn         *grpc.ClientConn
}

var (
	Gmgr *GrpcManger
)

func InitGRPC() {
	conn, err := grpc.Dial("127.0.0.1:18888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	Gmgr = &GrpcManger{conn: conn}
	Gmgr.taiKouServer = proto.NewTaiKouClient(conn)
}

func (g *GrpcManger) Close() {
	if err := g.conn.Close(); err != nil {
		log.Fatalln(err)
	}
}

func (g *GrpcManger) Ping(treat string) (string, error) {
	resp, err := Gmgr.taiKouServer.Ping(context.Background(), &proto.PingReq{Treat: treat})
	if err != nil {
		util.Logger.Error.Println(resp.S.Msg)
		return "", err
	}
	return resp.Msg, nil
}

func (g *GrpcManger) GetSingleVertex(tag, propType, keyword string) (*graph.Vertex, error) {
	resp, err := Gmgr.taiKouServer.GetSingleVertex(context.Background(), &proto.GetSingleVertexReq{
		Tag:      &tag,
		PropType: &propType,
		KeyWord:  &keyword,
	})
	if err != nil {
		util.Logger.Error.Println(err.Error())
		return nil, err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return nil, fmt.Errorf("%s", *resp.S.Msg)
	}
	rtnv := graph.Vertex{
		TagID: resp.Ver.TagID,
		Tag:   *resp.Ver.TagStr, // maybe panic
		Vid:   resp.Ver.Vid,
		Prop:  graph.NewProperty(resp.Ver.Property),
	}
	//util.Logger.Info.Println(resp)
	return &rtnv, nil
}

func (g *GrpcManger) GetEdgeList(tagId uint32, vid uint64) ([]*graph.Edge, error) {
	resp, err := Gmgr.taiKouServer.GetEdgeList(context.Background(), &proto.GetEdgeListReq{
		TagID: tagId,
		Vid:   vid,
	})
	if err != nil {
		util.Logger.Error.Println(err.Error())
		return nil, err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return nil, fmt.Errorf("%s", *resp.S.Msg)
	}
	rtnEdges := make([]*graph.Edge, 0)
	for _, e := range resp.EdgeList {
		rtnEdges = append(rtnEdges, &graph.Edge{
			TagID:     e.TagID,
			Tag:       *e.TagStr,
			Vid:       e.Vid,
			EdgeType:  *e.EdgeTypeStr,
			PeerTagID: e.PeerTagID,
			PeerTag:   *e.PeerTagStr,
			PeerVid:   e.PeerVid,
			Rank:      e.Rank,
			Prop:      graph.NewProperty(e.Property),
		})
	}
	return rtnEdges, nil
}

func (g *GrpcManger) GetSubGraph(tagId uint32, vid uint64) (*graph.SubGraph, error) {
	resp, err := Gmgr.taiKouServer.GetSubGraph(context.Background(), &proto.GetSubGraphReq{
		TagID: tagId,
		Vid:   vid,
	})

	if err != nil {
		util.Logger.Error.Println(err.Error())
		return nil, err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return nil, fmt.Errorf("%s", *resp.S.Msg)
	}
	rtnsg := graph.SubGraph{}
	rtnsg.V = graph.Vertex{
		Tag:  *resp.Sg.Ver.TagStr, // maybe panic
		Vid:  resp.Sg.Ver.Vid,
		Prop: graph.NewProperty(resp.Sg.Ver.Property),
	}

	for _, ie := range resp.Sg.InEdge {
		rtnsg.InEdge = append(rtnsg.InEdge, &graph.Edge{
			TagID:     ie.TagID,
			Tag:       *ie.TagStr,
			Vid:       ie.Vid,
			EdgeType:  *ie.EdgeTypeStr,
			PeerTagID: ie.PeerTagID,
			PeerTag:   *ie.PeerTagStr,
			PeerVid:   ie.PeerVid,
			Rank:      ie.Rank,
			Prop:      graph.NewProperty(ie.Property),
		})
	}

	for _, oe := range resp.Sg.OutEdge {
		rtnsg.InEdge = append(rtnsg.OutEdge, &graph.Edge{
			TagID:     oe.TagID,
			Tag:       *oe.TagStr,
			Vid:       oe.Vid,
			EdgeType:  *oe.EdgeTypeStr,
			PeerTagID: oe.PeerTagID,
			PeerTag:   *oe.PeerTagStr,
			PeerVid:   oe.PeerVid,
			Rank:      oe.Rank,
			Prop:      graph.NewProperty(oe.Property),
		})
	}

	return &rtnsg, nil
}

func (g *GrpcManger) CreateKG(kgName, meta string) error {
	resp, err := Gmgr.taiKouServer.CreateKG(context.Background(), &proto.CreateKGReq{
		KgName: kgName,
		Meta:   meta,
	})
	if err != nil {
		util.Logger.Error.Println(err.Error())
		return err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return fmt.Errorf("%s", *resp.S.Msg)
	}
	return nil
}

func (g *GrpcManger) PutVertexListReq(vertexList []*graph.WriteVertex) error {
	var wv []*proto.WriteVertex
	for _, v := range vertexList {
		wv = append(wv, &proto.WriteVertex{
			Tag:        v.Tag,
			UniqueName: v.UniqueName,
			Property:   v.Property,
		})
	}

	resp, err := Gmgr.taiKouServer.PutVertexList(context.Background(), &proto.PutVertexListReq{
		VertexList: wv,
	})
	if err != nil {
		util.Logger.Error.Println(err.Error())
		return err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return fmt.Errorf("%s", *resp.S.Msg)
	}

	return nil
}

func (g *GrpcManger) PutEdgeListReq(edgeList []*graph.WriteEdge) error {
	var we []*proto.WriteEdge
	for _, e := range edgeList {
		we = append(we, &proto.WriteEdge{
			VertexTag:      e.Tag,
			VertexName:     e.UniqueName,
			Direction:      e.Direction,
			EdgeType:       e.EdgeType,
			VertexTagPeer:  e.TagPeer,
			VertexNamePeer: e.UniqueNamePeer,
			Props:          e.Property,
		})
	}

	resp, err := Gmgr.taiKouServer.PutEdgeList(context.Background(), &proto.PutEdgeListReq{
		EdgeList: we,
	})
	if err != nil {
		util.Logger.Error.Println(err.Error())
		return err
	}
	if resp.S.Code != 0 {
		util.Logger.Error.Println(*resp.S.Msg)
		return fmt.Errorf("%s", *resp.S.Msg)
	}
	return nil
}
