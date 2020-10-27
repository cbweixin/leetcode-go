package main

type HitCounter struct {
	hits      []int
	threshold int
}

/** Initialize your data structure here. */
func Constructor2() HitCounter {
	return HitCounter{threshold: 5 * 60}
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	this.hits = append(this.hits, timestamp)

	for i, ts := range this.hits {
		if timestamp-ts >= this.threshold {
			this.hits = this.hits[i:] //leaves a straggler
			break
		}
	}
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	n := len(this.hits)

	for _, ts := range this.hits {
		if timestamp-ts >= this.threshold {
			n--
		} else {
			return n
		}
	}
	return n
}

func main() {

}
