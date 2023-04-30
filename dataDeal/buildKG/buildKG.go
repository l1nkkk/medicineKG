package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/l1nkkk/medicineKG/src/graph"
	"github.com/l1nkkk/medicineKG/src/rpc"
	"github.com/l1nkkk/medicineKG/src/util"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// -entity="/Users/l1nkkk/dataset/medicine/drug_neo4j_data/entity.csv" -relation="/Users/l1nkkk/dataset/medicine/drug_neo4j_data/relation.csv" -meta="/Users/l1nkkk/project/go/medicineKG/data/meta.json"
// -entity="/Users/l1nkkk/project/go/medicineKG/data/testentity.csv" -relation="/Users/l1nkkk/project/go/medicineKG/data/testrelation.csv" -meta="/Users/l1nkkk/project/go/medicineKG/data/meta.json"

const Batch = 10000

func main() {
	entity := flag.String("entity", "/your/targe/file", "the path of entity file")
	relation := flag.String("relation", "/your/targe/file", "the path of relation file")
	meta := flag.String("meta", "/your/targe/file", "the path of meta file")
	db := flag.String("db", "dbname", "db name")
	mode := flag.String("m", "append", "mode: append or trunk")
	flag.Parse()
	fmt.Println("entity file: ", *entity)
	fmt.Println("relation file: ", *relation)
	fmt.Println("meta file: ", *meta)
	fmt.Println("db name: ", *db)
	fmt.Println("mode: ", *mode)
	fmt.Println("Are you sure to start data-import-procedure with the above parameters?[Y/N]")
	var sel string
	fmt.Scanf("%s", &sel)
	sel = string(bytes.ToUpper([]byte(sel)))
	switch sel {
	case "Y":
		start := time.Now()
		rpc.InitGRPC()
		util.InitLog()
		b := NewBuildKG(*entity, *relation, *meta, *db, *mode)
		b.Do()
		elapsed := time.Since(start)
		fmt.Printf("spend time：%s\n", elapsed)
	case "N":
		fmt.Println("exist")
	default:
		log.Fatalf("error param")
	}
}

type KGBuilder struct {
	entityFile   string
	relationFile string
	metaFile     string
	dbName       string
	mode         string

	//
	idToTag             map[string]string
	firstKnowledgeGroup int64
}

func NewBuildKG(entityFile, relationFile, metaFile, dbName, mode string) *KGBuilder {
	return &KGBuilder{
		entityFile:          entityFile,
		relationFile:        relationFile,
		metaFile:            metaFile,
		dbName:              dbName,
		mode:                mode,
		idToTag:             make(map[string]string, 5000000),
		firstKnowledgeGroup: 4469557,
	}
}

func (kg *KGBuilder) Do() {
	kg.PushMeta()
	kg.PushVertex()
	kg.PushEdge()
}

func (kg *KGBuilder) PushMeta() {
	file, err := os.Open(kg.metaFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := fileInfo.Size()
	metaData := make([]byte, fileSize)
	_, err = file.Read(metaData)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(buffer))
	rpc.Gmgr.CreateKG("medical", string(metaData))
}

func (kg *KGBuilder) PushVertex() {
	wvList := []*graph.WriteVertex{}

	fs, err := os.Open(kg.entityFile)
	if err != nil {
		log.Fatalf("open file error")
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	r.FieldsPerRecord = -1
	count := 0
	for {
		count++
		row, err := r.Read()
		if count == 1 {
			continue
		}
		if err != nil && err != io.EOF {
			log.Printf("can not read, err is %+v", err)
			continue
		}
		if err == io.EOF {
			break
		}
		//if len(row) != 4 {
		//	fmt.Println(row)
		//	continue
		//}
		wvList = append(wvList, &graph.WriteVertex{
			Tag:        row[1],
			UniqueName: row[0],
			Property:   fmt.Sprintf("{\"unique_name\":\"%s\", \"name\":\"%s\" }", row[0], row[2]),
		})
		idnum, _ := strconv.ParseInt(row[0], 10, 64)
		if idnum < kg.firstKnowledgeGroup {
			kg.idToTag[row[0]] = row[1]
		}
		if (count % Batch) == 0 {
			util.Logger.Info.Println("write vertex: ", count)
			rpc.Gmgr.PutVertexListReq(wvList)
			wvList = []*graph.WriteVertex{}
		}
	}
	if (count % Batch) != 0 {
		rpc.Gmgr.PutVertexListReq(wvList)
	}
}

func (kg *KGBuilder) PushEdge() {
	weList := []*graph.WriteEdge{}

	fs, err := os.Open(kg.relationFile)
	if err != nil {
		log.Fatalf("open file error")
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	r.FieldsPerRecord = -1
	count := 0
	for {
		count++
		row, err := r.Read()
		if count == 1 {
			continue
		}
		if err != nil && err != io.EOF {
			log.Printf("can not read, err is %+v", err)
			continue
		}
		if err == io.EOF {
			break
		}
		if len(row) != 3 {
			fmt.Println(row)
			continue
		}
		weList = append(weList, &graph.WriteEdge{
			Tag:            kg.getTag(row[0]),
			UniqueName:     row[0],
			Direction:      uint32(graph.OutEdge),
			EdgeType:       row[2],
			TagPeer:        kg.getTag(row[1]),
			UniqueNamePeer: row[1],
			Property:       "",
		})
		if count%Batch == 0 {
			util.Logger.Info.Println("write edge: ", count)
			rpc.Gmgr.PutEdgeListReq(weList)
			weList = []*graph.WriteEdge{}
		}
	}
	if count%Batch != 0 {
		rpc.Gmgr.PutEdgeListReq(weList)
	}
}

func (kg *KGBuilder) getTag(id string) string {
	idnum, _ := strconv.ParseInt(id, 10, 64)
	if idnum >= kg.firstKnowledgeGroup {
		return "知识组"
	}
	return kg.idToTag[id]
}
