package main


type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	lastPosition := len(q.items) - 1
	if lastPosition == -1 {
		return -1
	}
	item := q.items[0]
	if lastPosition == 0 {
		q.items = []int{}
	} else {
		q.items = q.items[1:]
	}
	return item
}


func main(){
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)


	println(q.Dequeue())
	println(q.Dequeue())
	println(q.Dequeue())
}
