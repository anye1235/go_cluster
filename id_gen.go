package main

import(
	"fmt"
	"sync"
//	"time"
)

type UID struct{
	pChan chan int
	sync.WaitGroup
}

func new_id( id int ) int {
	return id
}

func (uid *UID) get() int {
	var value int
	select{
		case value =  <-uid.pChan:
			fmt.Printf("get curr id =%d\n", value)
//		default:
//			value = new_id(1)

	}

	go func(uid *UID){
		uid.Done()
		uid.put(value+1)
		//uid.Done()
	}(uid)
	return value
}

func (uid *UID) put(value  int) {			

	select{
	case  uid.pChan <- value:
		fmt.Printf("put new id = %d\n", value)
//	default:
//		fmt.Printf("put error : %d\n", value)
	}
}

func get_test( ){
	id := &UID{ pChan: make(chan int)}
	id.Add(20)
	go func(id *UID){ 
		id.pChan <- 1
	}(id)

	go func(id *UID){
		fmt.Println("start  test")
		//time.Sleep(time.Second*10)
		for i:=0;i<10; i++ {
			go func(i int){
				uid := id.get()
				fmt.Printf("g %d get uid=%d\n",i, uid) 
				id.Done()
			}(i)
		}
	}(id)
	
	id.Wait()
	close(id.pChan)

}

func main(){
//	var wg sync.WaitGroup
//	done := make( chan int)
//	id := make( chan id_gen)

   get_test()

//	wg.Wait()
//	<-done
}
