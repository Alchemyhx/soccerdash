package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/latifrons/soccerdash"
)

func main() {
	r := soccerdash.Reporter{
		Id:            "node_1",
		TargetAddress: "127.0.0.1:8080",
	}

	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Seconds().Do(func() {
		r.Report("TxPoolNum", "1", false)
	})

	s.Every(10).Seconds().Do(func() {
		r.Report("ConnNum", "10", false)
		r.Report("IsProducer", false, false)
	})

	s.Every(30).Seconds().Do(func() {
		r.Report("NodeDelay", "30", false)
	})

	s.Every(1).Minute().Do(func() {
		r.Report("Version", "1.0", false)
		r.Report("NodeName", "123", false)
	})

	s.Every(1).Minute().Do(func() {
		SeqMsg := `{
			"Type":1,
			"Hash":"0xb8ed9e555a2f1df24150f443c6fa1b695366508cbb6eb85041e7f6db3f9613e2",
			"ParentsHash":[
				"0xf1efee80c371a20090bf27e6ae9d015ad25cd9ff0fb3cf040458d02b7d51c8ee"
			],
			"AccountNonce":10330,
			"Height":10330,
			"PublicKey":"BIDG6ARHwZspJ1gUJTuHrjre0pyL7YJvG7E+UnbujaOVym+AjVi6jZIhptr92lwqi0xkXyXnI10gHt8RtnRHzrI=",
			"Signature":"0xeaa5753954c3021e5108caf14b7ff614234f7f1bbf3938d644fb20b665b2dc951a427d5be5553a3e7ed669d134a18d27052d5c7a99062f4dd3e687c2107939c001",
			"MineNonce":0,
			"Weight":10330,
			"Version":0,
			"Issuer":"0x7349f7a6f622378d5fb0e2c16b9d4a3e5237c187",
			"BlsJointSig":"0x",
			"BlsJointPubKey":"0x",
			"StateRoot":"0xe7537563bd2288cf95cd1d84ea24bf1bc5f642c7a1b266cceffca814c79d859c",
			"Proposing":false,
			"Timestamp":1591769404219
		}`
		r.Report("LatestSequencer", SeqMsg, false)
	})

	s.StartBlocking()
}
