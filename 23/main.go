package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintAlphabet() {
	alphabet := "abcdefghijk"
	for _, v := range alphabet {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d", v)
	}
}

func PrintNumber() {
	for i:= 1; i <=10; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다\n", a, b, sum)
	wg.Done()
}

type Account struct {
	Balance int
}

var mutex sync.Mutex 
func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock() 
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func diningProblem(name string, first, second *sync.Mutex, firstname, secondname string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다. \n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstname)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondname)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과 %d\n", j.index, j.index * j.index)
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <- ch

	time.Sleep(time.Second)
	fmt.Println("Square: ", n*n)
	wg.Done()
}

func square2(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10*time.Second)

	for {
		select {
		case <- tick:
			fmt.Println("Tick")
		case <- terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <- ch:
			fmt.Println("Square: ", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 1
	// go PrintAlphabet()
	// go PrintNumber()
	// time.Sleep(3 * time.Second)

	//2
	// wg.Add(10)
	// for i := 0; i <10; i++ {
	// 	go SumAtoB(1, 100000000)
	// }
	// wg.Wait()

	// 3
	// account := &Account{10}
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for {
	// 			DepositAndWithdraw(account)
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// 4
	// rand.Seed(time.Now().UnixNano())
	// wg.Add(2)
	// fork := &sync.Mutex{}
	// spoon := &sync.Mutex{}

	// go diningProblem("A", fork, spoon, "포크", "스푼")
	// go diningProblem("B", spoon, fork,  "스푼", "포크")
	// wg.Wait()

	// 5
	// var jobList [10]Job
	// for i := 0; i < 10; i++ {
	// 	jobList[i] = &SquareJob{i}
	// }

	// wg.Add(10)

	// for i := 0; i < 10; i++ {
	// 	job := jobList[i]
	// 	go func() {
	// 		job.Do()
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// 6
	// ch := make(chan int)

	// wg.Add(1)
	// go square(&wg, ch)
	// ch <- 9
	// wg.Wait()

	//7
	ch := make(chan int)

	go square2(&wg, ch)
	ch <- 9
	fmt.Println("never print")
}