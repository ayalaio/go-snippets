package singleton

type channelsingleton struct {
	val int
}

var addCh chan bool = make(chan bool)
var getCountCh chan chan int = make(chan chan int)
var quitCh chan bool = make(chan bool)

var instance channelsingleton

func GetInstance() *channelsingleton {
	return &instance
}

func (c *channelsingleton) AddOne() {
	addCh <- true
}

func (c *channelsingleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh
	return <-resCh
}

func (c *channelsingleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
		for {
			select {
			case <-addCh:
				count++
			case ch := <-getCountCh:
				ch <- count
			case <-quitCh:
				break
			}
		}
	}(addCh, getCountCh, quitCh)
}
