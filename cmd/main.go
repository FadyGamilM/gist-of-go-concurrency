package main

import "sync"

func main() {
	usersIds := []int{10, 11, 12}
	wg := sync.WaitGroup{}

	for _, userId := range usersIds {
		// handle concurrency
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*
				then perform the business logic
				(without passing anything related
				to go concurrency)
			*/
			DoAsyncWorkFixed(userId)
		}()
	}

	wg.Wait()
}

func DoAsyncWorkFixed(userId int) error {
	// Do business logic

	return nil
}

func DoAsyncWork(wg *sync.WaitGroup, userId int) error {
	defer wg.Done()

	// Do business logic

	return nil
}
