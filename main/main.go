package main

import (
	"fmt"
	"time"

	"./types"

	"github.com/golang/protobuf/proto"
)

func main() {
	decode(&types.OldOrder{
		Time:    int32(time.Now().Unix()),
		Userid:  1,
		PriceE4: 20.0,
		Desc:    "OldOrder",
	}, &types.OldOrder{}, "[normal] ")

	decode(&types.NewOrderAdd{
		Time:      int32(time.Now().Unix()),
		Userid:    1,
		PriceE4:   20.0,
		Desc:      "NewOrderAdd",
		OrderType: 1,
	}, &types.OldOrder{}, "[add] ")

	decode(&types.NewOrderDelete{
		Time:   int32(time.Now().Unix()),
		Userid: 1,
		Desc:   "NewOrderDelete",
	}, &types.OldOrder{}, "[delete] ")

	decode(&types.NewOrderModifyLabelNo{
		Time:    int32(time.Now().Unix()),
		Userid:  1,
		PriceE4: 20.0,
		Desc:    "NewOrderModifyLabelNo",
	}, &types.OldOrder{}, "[ModifyLabelNo] ")

	decode(&types.NewOrderModifyLabelTypeOne{
		Time:    int32(time.Now().Unix()),
		Userid:  1,
		PriceE4: "20.0",
		Desc:    "NewOrderModifyLabelTypeOne",
	}, &types.OldOrder{}, "[ModifyLabelTypeOne] ")

	decode(&types.NewOrderModifyLabelTypeTwo{
		Time:    int32(time.Now().Unix()),
		Userid:  1,
		PriceE4: 20,
		Desc:    "NewOrderModifyLabelTypeTwo",
	}, &types.OldOrder{}, "[ModifyLabelTypeTwo] ")

	decode(&types.NewOrderModifyLabelTypeThree{
		Time:    int32(time.Now().Unix()),
		Userid:  1,
		PriceE8: 20,
		Desc:    "NewOrderModifyLabelTypeThree",
	}, &types.OldOrder{}, "[ModifyLabelTypeThree] ")
}

func decode(req proto.Message, decodeObj proto.Message, logTag string) {
	// Marshal
	encoded, err := proto.Marshal(req)
	if err != nil {
		fmt.Printf(logTag+"Encode to protobuf data error: %v", err)
	}

	// Unmarshal
	err = proto.Unmarshal(encoded, decodeObj)
	if err != nil {
		fmt.Printf(logTag+"Unmarshal to struct error: %v", err)
	}

	fmt.Printf(logTag+"req: %v\n", req.String())
	fmt.Printf(logTag+"resp: %v\n", decodeObj.String())
}
