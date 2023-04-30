package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/l1nkkk/medicineKG/src/rpc"
	"github.com/l1nkkk/medicineKG/src/util"
	"net/http"
	"strconv"
)

// https://www.kancloud.cn/shuangdeyu/gin_book/949419
// https://gin-gonic.com/docs/examples/

func pingHandle(c *gin.Context) {
	var (
		code int
		msg  string
		data interface{}
	)
	treat, ok := c.Params.Get("treat")
	if !ok {
		code, msg = util.HttpParaError, "parameter 'treat' can't not be found"
		constructResp(c, code, msg, data)
		return
	}
	rpcResp, err := rpc.Gmgr.Ping(treat)
	if err != nil {
		code, msg = util.HttpInternalError, err.Error()
		constructResp(c, code, msg, data)
		return
	}
	code, msg, data = util.HttpOK, "", rpcResp
	constructResp(c, code, msg, data)
	return
}

func getVertexHandle(c *gin.Context) {
	var (
		code int
		msg  string
		data interface{}
	)
	tag, _ := c.GetQuery("tag")
	propType, _ := c.GetQuery("prop_type")
	keyword, _ := c.GetQuery("keyword")
	if tag == "" || propType == "" || keyword == "" {
		code, msg = util.HttpParaError, "parameter error"
		constructResp(c, code, msg, data)
		return
	}
	vertex, err := rpc.Gmgr.GetSingleVertex(tag, propType, keyword)
	if err != nil {
		code, msg = util.HttpInternalError, err.Error()
		constructResp(c, code, msg, data)
		return
	}

	code, msg, data = util.HttpOK, "", vertex
	constructResp(c, code, msg, data)
	return
}

func getEdgeListHandle(c *gin.Context) {
	var (
		code int
		msg  string
		data interface{}
	)
	tagIdStr, _ := c.GetQuery("tag_id")
	vidStr, _ := c.GetQuery("vid")
	tagId, err := strconv.ParseUint(tagIdStr, 10, 32)
	if err != nil {
		code, msg = util.HttpParaError, "parameter error"
		constructResp(c, code, msg, data)
		return
	}
	vid, err := strconv.ParseUint(vidStr, 10, 64)
	if err != nil {
		code, msg = util.HttpParaError, "parameter error"
		constructResp(c, code, msg, data)
		return
	}
	edgeList, err := rpc.Gmgr.GetEdgeList(uint32(tagId), vid)
	if err != nil {
		code, msg = util.HttpInternalError, err.Error()
		constructResp(c, code, msg, data)
		return
	}
	data, err = json.Marshal(edgeList)
	if err != nil {
		code, msg = util.HttpInternalError, err.Error()
		constructResp(c, code, msg, data)
		return
	}
	code, msg, data = util.HttpOK, "", edgeList
	constructResp(c, code, msg, data)
	return
}

func getSubGraphHandle(c *gin.Context) {
	var (
		code int
		msg  string
		data interface{}
	)
	tagIdStr, _ := c.GetQuery("tag_id")
	vidStr, _ := c.GetQuery("vid")
	tagId, err := strconv.ParseUint(tagIdStr, 10, 32)
	if err != nil {
		code, msg = util.HttpParaError, "parameter error"
		constructResp(c, code, msg, data)
		return
	}
	vid, err := strconv.ParseUint(vidStr, 10, 64)
	if err != nil {
		code, msg = util.HttpParaError, "parameter error"
		constructResp(c, code, msg, data)
		return
	}
	sg, err := rpc.Gmgr.GetSubGraph(uint32(tagId), vid)
	if err != nil {
		code, msg = util.HttpInternalError, err.Error()
		constructResp(c, code, msg, data)
		return
	}
	code, msg, data = util.HttpOK, "", sg
	constructResp(c, code, msg, data)
	return
}

func constructResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, util.HttpResp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
