package concurrency

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int
}

func NewAccount(balance int) *Account {
	return &Account{balance: balance}
}

//Race condition: When more than one goroutine accesses a function at the same time we could have a Race condition
//It means that our function could be modified by another goroutine and cause an runtime error (inconsistent data)
func (a *Account) DoDeposit(deposit int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()

	//Only one process can work with the data, all other processes (goroutines) must wait
	lock.Lock()
	a.balance = deposit
	lock.Unlock()
}

func (a *Account) DoRetire(mount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	wg.Done()
	lock.Lock()
	if a.balance-mount < 0 {
		fmt.Printf("\nThe mount exceeds actual balance")
	} else {
		a.balance = a.balance - mount
	}
	lock.Unlock()
}

func (a *Account) SeeMyBalance(lock *sync.RWMutex) int {
	lock.RLock()
	balance := a.balance
	lock.RUnlock()
	return balance
}

//mutex: allows us to block a certain block of code while a process is modifying data
func Testing5() {
	var wg sync.WaitGroup
	//Read and write with mutex
	var lock sync.RWMutex
	myAccount := NewAccount(0)
	fmt.Printf("\n my new account with %d balance", myAccount.balance)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go myAccount.DoDeposit((i+1)*100, &wg, &lock)
	}
	wg.Wait()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go myAccount.DoRetire((i+1)*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Printf("\n my actual %d balance", myAccount.balance)
}
