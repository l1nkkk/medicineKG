package web

import "github.com/gin-gonic/gin"

type webManger struct {
	r *gin.Engine
}

var WebMgr webManger

func (w *webManger) RegisterRouter() {
	w.r = gin.Default()
	w.r.GET("/ping/:treat", pingHandle)
	w.r.GET("/vertex", getVertexHandle)
	w.r.GET("/edge", getEdgeListHandle)
	w.r.GET("/subgraph", getSubGraphHandle)
}

func (w *webManger) Run() {
	w.r.Run()
}
