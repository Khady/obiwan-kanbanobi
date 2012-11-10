package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"log"
	"obi-wan-kanbanobi/protocole"
)

func main() {
	cible := message.CIBLE_USERS
	cmd := message.CMD_Y
	test := &message.Msg{
		Cible: &cible,
		Users: &message.Msg_Users{
			Cmd: &cmd,
		},
	}
	// test := &example.Test{
	// 	Label: proto.String("hello world"),
	// 	Type:  proto.Int32(17),
	// 	Optionalgroup: &example.Test_OptionalGroup{
	// 		RequiredField: proto.String("good bye"),
	// 	},
	// }
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &message.Msg{}
	fmt.Println(data, len(data), string(data))
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetProjects() != newTest.GetProjects() {
		log.Fatalf("data mismatch %q != %q", test.GetProjects(), newTest.GetProjects())
	}
}
