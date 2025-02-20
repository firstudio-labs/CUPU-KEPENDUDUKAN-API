package test

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	//wd, _ := os.Getwd()
	//fmt.Println("Current working directory:", wd)

	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result entity.User
	if err := pool.Find(result); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
