package protoTest

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestProto(t *testing.T) {
	//引用生成的.pd.go文件中的接口
	msg := &TestMsg{
		Id:    proto.Int32(10),
		Name:  proto.String("Proto-Test"),
		Reps:  []int32{1, 2, 3},
		Union: &TestMsg_Names{"Union_Name"},
		//Union: &TestMsg_Number{9},
		Optionalgroup: &TestMsg_OptionalGroup{
			GroupId: proto.String("good bye"),
		},
	}

	//编码
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
	fmt.Println("Marshal: ", string(data))

	//解码
	newMsg := &TestMsg{}
	err = proto.Unmarshal(data, newMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshal: ", newMsg)

	//可以拿一组数据进行验证
	if msg.GetId() != newMsg.GetId() {
		fmt.Println("Not equal")
	} else {
		fmt.Println("Equal")
	}
}
