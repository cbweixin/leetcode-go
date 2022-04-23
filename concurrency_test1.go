package main

import (
	"./concurrency/wait_group_sample"
)
import "./concurrency/channel_sample"
import "./concurrency/context_sample"

func main() {
	wait_group_sample.Wg_ex1()
	channel_sample.Channel_test()
	context_sample.Exit_test()
	channel_sample.Send_Recieve()
	channel_sample.Chan_With_Map()
	channel_sample.Chan_With_Map2()
	channel_sample.Chan_With_Map3()
	channel_sample.Close_test()
	channel_sample.Single_Direction_Channel()
	channel_sample.Chan_Conv()
	channel_sample.Chan_With_For()
	channel_sample.Select_test()
	channel_sample.Select_test2()
	channel_sample.Select_test3()
	channel_sample.Blocking_Channel()
	channel_sample.Timer_test()
	channel_sample.Timer_test2()
	channel_sample.Ticker_test()

	// without this still works fine, weird.
	//time.Sleep(5*time.Second)
}
