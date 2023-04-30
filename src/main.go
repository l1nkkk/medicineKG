package main

import (
	"github.com/l1nkkk/medicineKG/src/rpc"
	"github.com/l1nkkk/medicineKG/src/util"
	"github.com/l1nkkk/medicineKG/src/web"
)

func main() {
	util.InitLog()
	rpc.InitGRPC()
	web.WebMgr.RegisterRouter()
	web.WebMgr.Run()
}
