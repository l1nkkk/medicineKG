package rpc

import (
	"context"
	"github.com/l1nkkk/medicineKG/src/rpc/proto"
	"testing"
)

func TestInitGRPC(t *testing.T) {
	InitGRPC()
	resp, err := Gmgr.taiKouServer.Ping(context.Background(), &proto.PingReq{Treat: "medicineKG"})
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log(resp.Msg)

}
