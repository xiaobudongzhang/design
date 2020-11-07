package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type signle struct {

}

var signleInstance *signle

func getInstance()*signle  {
	if signleInstance != nil {
		return signleInstance
	}
	lock.Lock()
	defer lock.Unlock()
	if signleInstance != nil {
		return signleInstance
	}
	signleInstance = &signle{}
	fmt.Println("creating")
	return signleInstance
}