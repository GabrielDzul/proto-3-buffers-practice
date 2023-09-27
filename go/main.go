package main

import (
	"fmt"
	"reflect"

	pb "github.com/gabriel-dzul/protocol-3-buffers-course/go/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnun())
	//fmt.Println("This should be an Id:")
	//doOneOf(&pb.Result_Id{Id: 42})
	//fmt.Println("This should be a Message:")
	//doOneOf(&pb.Result_Message{Message: "a message"})
	//fmt.Println(doMaps())
	//doFile(doSimple())
	//jsonString := doToJSON(doComplex())
	//message := doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	//fmt.Println(jsonString)
	//fmt.Println(message)
	fmt.Println(doFromJSON(`{"id": 42, "unknown": "test"}`, reflect.TypeOf(pb.Simple{})))
}

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "A name",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDumy: &pb.DummyExample{Id: 42, Name: "My Name"},
		MultipleDumies: []*pb.DummyExample{
			{Id: 43, Name: "My Name 2"},
			{Id: 44, Name: "My Name 3"},
		},
	}
}

func doEnun() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doMaps() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"mi_id":   {Id: 41},
			"mi_id_2": {Id: 42},
			"mi_id_3": {Id: 43},
			"mi_id_4": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}
