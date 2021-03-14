package main

import (
	"Concurrent/PubSub"
	"context"
)
func Get(ctx context.Context,k key)
func main(){
	//Simple.Run()
	//Simple.Channel_run()
	//Simple.Channel_run2()
	//Simple.Channnel_run3()
	//Simple.Wait_group()
	//ProConsum.Product_Consumer()
	ctx:=context.WithValue(context.Background(),key("asong"),"hello")
	Get(ctx,key("asong"))
	Get(ctx,key("song"))
	PubSub.Run()
}
