package ProConsum

import (
	"fmt"
	"sync"
)

type Config struct {
	TaskNum         int
	ConcurrentNum   int
	ProductFunc     func(interface{}) interface{}
	ProductFuncArgs []interface{}
	ConsumeFunc     func(interface{})
}

func Product_Consumer2(c *Config){
	//taskNum := conf.TaskNum
	//task := make(chan interface{}, conf.ConcurrentNum)
	task := make(chan interface{})
	quit := make(chan int, 2)
	defer func() {
		close(quit)
		//fmt.Println("quit close")
		//close(task)

		//fmt.Println("task close")
		err := recover()
		if err != nil {
			fmt.Println("err: ", err)
		}
	}()

	var wg1 sync.WaitGroup
	go func() {
		product := c.ProductFunc
		for _, args := range c.ProductFuncArgs {
			wg1.Add(1)
			go func(args interface{}) {
				//fmt.Println("args: ", args)
				task <- product(args)
				wg1.Done()
			}(args)
		}
		wg1.Wait()
		close(task)
		//fmt.Println("task close")
		quit <- 1
	}()
	fmt.Println("start")

	go func() {
		var wg2 sync.WaitGroup
		//wg2.Add(taskNum)
		//go func() {
		//	wg2.Wait()
		//	closeConsume <- 1
		//}()
		for {
			select {
			case canvasId, ok := <-task:
				//fmt.Println("canvasId : ", canvasId, "  ok : ", ok)
				if !ok {
					wg2.Wait()
					quit <- 1
					return
				}
				wg2.Add(1)
				go func() {
					//fmt.Println("consume ")
					c.ConsumeFunc(canvasId)
					wg2.Done()
				}()
			}
		}
	}()
	<-quit
	<-quit

}