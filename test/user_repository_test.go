package test

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUser(t *testing.T) {
	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result []entity.User
	if err := pool.Find(&result).Error; err != nil {
		panic(err)
	}

	for i, v := range result {
		fmt.Println("HASIL NYA ", i, v)
	}
}

func TestGetSourceInternet(t *testing.T) {
	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result []entity.SourceInternet
	if err := pool.Find(&result).Error; err != nil {
		panic(err)
	}

	for i, v := range result {
		fmt.Println("HASIL NYA ", i, v)
	}

}

func TestGetPacketInternet(t *testing.T) {
	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result []entity.PacketInternet
	if err := pool.Find(&result).Error; err != nil {
		panic(err)
	}

	for i, v := range result {
		fmt.Println("HASIL NYA ", i, v)
	}

}

func TestGetComplaint(t *testing.T) {
	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result []entity.Complaint
	// Preload relasi User dan PacketInternet
	if err := pool.Preload("User").Preload("PacketInternet").Find(&result).Error; err != nil {
		panic(err)
	}

	for i, v := range result {
		fmt.Println("HASIL NYA ", i, v)
	}
}

func TestGetSubsPacket(t *testing.T) {
	pool, err := cfg.GetPool(cfg.GetConfig())
	assert.Nil(t, err)

	var result []entity.SubsPacket
	if err := pool.Find(&result).Error; err != nil {
		panic(err)
	}

	for i, v := range result {
		fmt.Println("HASIL NYA ", i, v)
	}
}
