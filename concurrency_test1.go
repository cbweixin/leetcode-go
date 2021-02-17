package main

import "./concurrency/wait_group_sample"
import "./concurrency/channel_sample"

func main() {
	wait_group_sample.Wg_ex1()
	channel_sample.Channel_test()
}
