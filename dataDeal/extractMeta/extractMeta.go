package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type VertexTag struct {
	VertexTag     string   `json:"vertex_tag"`
	PropertyIndex []string `json:"property_index,omitempty"`
}

type EdgeType struct {
	EdgeType      string   `json:"edge_type"`
	PropertyIndex []string `json:"property_index,omitempty"`
}

type Meta struct {
	VertexTagList []*VertexTag `json:"vertex_tag_list"`
	EdgeTypeList  []*EdgeType  `json:"edge_type_list"`
}

type Entity struct {
	ID         int64
	Tag        string
	UniqueName string
	Number     string
}

func main() {
	entityFile := flag.String("entity", "/your/targe/file", "the path of entity file")
	relationFile := flag.String("relation", "/your/targe/file", "the path of relation file")
	metaFile := flag.String("out", "/your/targe/file", "the path of output meta file")
	flag.Parse()

	fmt.Println("entity file: ", *entityFile)
	fmt.Println("relation file: ", *relationFile)
	fmt.Println("output meta file: ", *metaFile)

	tagSet, typeSet := ExtractMetaData(*entityFile, *relationFile)
	meta := Meta{}
	for tag := range tagSet {
		meta.VertexTagList = append(meta.VertexTagList, &VertexTag{VertexTag: tag})
	}
	for atype := range typeSet {
		meta.EdgeTypeList = append(meta.EdgeTypeList, &EdgeType{EdgeType: atype})
	}
	file, err := os.OpenFile(*metaFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("open file error")
	}
	defer file.Close()

	metaByte, err := json.Marshal(meta)
	if err != nil {
		log.Fatalf(err.Error())
	}
	file.Write(metaByte)

}

func ExtractMetaData(entityFile, relationFile string) (tagSet, typeSet map[string]struct{}) {

	start := time.Now()
	tagSet = ExtractTagFromEntity(entityFile)
	fmt.Printf("tag set:")
	for tag, _ := range tagSet {
		fmt.Printf(" %s", tag)
	}
	fmt.Println("\ntag size:", len(tagSet))

	typeSet = ExtractTypeFromRelation(relationFile)
	fmt.Printf("type set:")
	for atpye, _ := range typeSet {
		fmt.Printf(" %s", atpye)
	}
	fmt.Println("\ntype size:", len(typeSet))
	elapsed := time.Since(start)
	fmt.Printf("spend time：%s\n", elapsed)
	return tagSet, typeSet
}

func ModBuild() {
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
	fmt.Println("Are you sure to start data import with the above parameters?[Y/N]")
	var sel string
	fmt.Scanf("%s", sel)
	fmt.Printf("import KG done: %d entitys, %d relations", 7153420, 6620027)
}

func ExtractTagFromEntity(file string) map[string]struct{} {
	tagSet := make(map[string]struct{})

	fs, err := os.Open(file)
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
		tagSet[row[1]] = struct{}{}
	}
	return tagSet
}

func ExtractTypeFromRelation(file string) map[string]struct{} {
	typeSet := make(map[string]struct{})

	fs, err := os.Open(file)
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
		typeSet[row[2]] = struct{}{}
	}
	return typeSet
}
