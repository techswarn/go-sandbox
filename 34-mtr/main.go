package main

import (
	"fmt"
	"github.com/smuzoey/go-mtr"
	"time"
)

func main() {
	tracer, err := go_mtr.NewTrace(go_mtr.Config{
		ICMP:        true,
		UDP:         false,
		MaxUnReply:  8,
		NextHopWait: time.Millisecond * 200,
	})
	if err != nil {
		panic(err)
	}
	t, err := go_mtr.GetTrace(&go_mtr.Trace{
		SrcAddr: go_mtr.GetOutbondIP(),
		DstAddr: "143.244.131.85",
		SrcPort: 80,
		DstPort: 80,
		MaxTTL:  30,
		Retry:   2,
	})
	if err != nil {
		panic(err)
	}
	res, _ := tracer.BatchTrace([]go_mtr.Trace{*t}, 1)
	for _, r := range res {
		fmt.Println(r.Marshal())
		fmt.Println(r.MarshalAggregate())
	}
}