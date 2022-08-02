package main

import (
	"fmt"
	"time"

	"github.com/KL-Lru/sample-web-service/pkg/repos"
)

func main() {
	repo, err := repos.NewRepository()
	if err != nil {
		fmt.Println(err)
		panic("Can't connect SQL")
	}

	timeout := time.After(10 * time.Second)
	ticker := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-timeout:
			fmt.Println("Can't connect SQL")
			return
		case <-ticker:
			err := repo.Ping()
			if err == nil {
				users, err := repo.UserList()
				if err == nil {
					fmt.Println(users)
					return
				} else {
					fmt.Println("Can't get user info")
					return
				}
			}
		}
	}
}
