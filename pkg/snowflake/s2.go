package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func test2() {
	startTime := "2021-08-06"
	var machineID uint16 = 1

	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		fmt.Printf("set base time err: %v\n", err)
		return
	}
	sonyMachineID = machineID
	sonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	})

	id, err := sonyFlake.NextID()
	if err != nil {
		fmt.Printf("generate id err: %v\n", err)
		return
	}
	fmt.Println("generate id:", id)

}
