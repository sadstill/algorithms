package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	queue := append(this.queue, t)
	window := t - 3000
	for queue[0] < window {
		queue = queue[1:]
	}
	return len(queue)
}
