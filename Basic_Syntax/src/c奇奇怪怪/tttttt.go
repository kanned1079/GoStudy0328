package main

var goRoutineNums = 8

func write(numChan chan int) {
	for i := 1; i <= 100; i++ {
		numChan <- i
	}
	close(numChan)
}

func main() {

}
