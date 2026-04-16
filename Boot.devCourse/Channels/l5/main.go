package main

func countReports(numSentCh chan int) int {
			running_total := 0 
	for {

		c,ok := <- numSentCh
		if !ok {
			break
		}else  {
			running_total +=c 
		}
	}
return  running_total
	}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
