package main

import (
	"fmt"

	"github.com/KL-Lru/sample-web-service/pkg/repos"
)

func main() {
	repo, err := repos.NewRepository()
	if err != nil {
		fmt.Println(err)
		panic("Can't connect SQL")
	}

	fmt.Println(repo.UserList())
}
