package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"
)

func test() {

	startTime := "2021-08-06"
	var machineID int64 = 1

	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Printf("set base time err: %v\n", err)
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		fmt.Printf("make node err: %v\n", err)
		return
	}
	id := node.Generate().Int64()
	fmt.Println("generate id:", id)

}
