package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/polar9527/learning-go/grpc/protobuf-go/simple"
)

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	smR := &simplepb.SimpleMessage{}
	readToFile("simple.bin", smR)
	fmt.Println(smR.String())

}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written.")
	return err
}

func readToFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protobuf struct", err)
		return err
	}

	return err
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 3, 4, 6},
	}

	fmt.Println(sm.String())
	sm.Name = "My new name"
	fmt.Println(sm.String())
	return &sm
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to json", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't convert from json", err)
	}
}

func JSONDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println(sm2.String())
}

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	JSONDemo(sm)
}
