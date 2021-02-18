package main

import "./concurrency/wait_group_sample"
import "./concurrency/channel_sample"
import "./concurrency/context_sample"

func main() {
	wait_group_sample.Wg_ex1()
	channel_sample.Channel_test()
	context_sample.Exit_test()
	channel_sample.Send_Recieve()
	channel_sample.Chan_With_Map()
	channel_sample.Chan_With_Map2()
}
