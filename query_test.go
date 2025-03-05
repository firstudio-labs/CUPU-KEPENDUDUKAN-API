package JARITMAS_API

import (
	"context"
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/cfg"
	"github.com/firstudio-lab/JARITMAS-API/internal/entity"
	"log"
	"testing"
)

func TestFIndAllMemberByKK(t *testing.T) {
	ctx := context.Background()
	pool, err := cfg.GetPool(cfg.GetConfig())
	fmt.Println(err)

	kk := 3323092402072308
	var familyMember []entity.Citizen
	if err := pool.WithContext(ctx).Where("kk = ?", kk).Find(&familyMember).Error; err != nil {
		//ERROR KK  NOFOUND OR APA LAH HITU
		log.Fatalf("error query")

	}

	if len(familyMember) == 0 {
		t.Logf("No family members found for KK: %d", kk)
	}

	fmt.Println(familyMember)
}

func TestFIndAllCitizens(t *testing.T) {
	ctx := context.Background()
	pool, _ := cfg.GetPool(cfg.GetConfig())

	var citizens []entity.Citizen
	if err := pool.WithContext(ctx).Find(&citizens).Error; err != nil {
		log.Fatalf("error query")
	}

	fmt.Println(citizens)

}
