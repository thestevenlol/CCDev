package main

import (
	"fmt"
	"sync"
)

func produce(wg *sync.WaitGroup, mutex *sync.Mutex, buf chan int, nProductions *int) {
	defer wg.Done()
	mutex.Lock()
	buf <- 1
	mutex.Unlock()
	*nProductions--
	fmt.Println("Produced")
}

func consume(wg *sync.WaitGroup, mutex *sync.Mutex, buf chan int, nConsumers *int) {
	defer wg.Done()
	mutex.Lock()
	<-buf
	mutex.Unlock()
	*nConsumers--
	fmt.Println("Consumed")
}

func main() {
	nProductions := 5
	nConsumers := 5
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	buf := make(chan int)

	wg.Add(nProductions + nConsumers)
	for i := 0; i < nProductions; i++ {
		go produce(&wg, &lock, buf, &nProductions)
	}

	for i := 0; i < nConsumers; i++ {
		go consume(&wg, &lock, buf, &nConsumers)
	}

	wg.Wait()

}
