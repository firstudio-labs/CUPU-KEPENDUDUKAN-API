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

func TestFindAllMemberBySimilarName(t *testing.T) {
	ctx := context.Background()
	pool, err := cfg.GetPool(cfg.GetConfig())
	if err != nil {
		t.Fatalf("Failed to get database pool: %v", err)
	}

	namePattern := "Muly" // Example: Find members with a name similar to 'Muly'
	var familyMember []entity.Citizen

	// Using LIKE operator for MySQL and selecting only the full_name field
	if err := pool.WithContext(ctx).Select("full_name").Where("full_name LIKE ?", "%"+namePattern+"%").Find(&familyMember).Error; err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}

	if len(familyMember) == 0 {
		t.Logf("No family members found with similar name to: %s", namePattern)
	} else {
		fmt.Println(familyMember)
	}
}

func TestFindAllJobsBySimilarName(t *testing.T) {
	ctx := context.Background()
	pool, err := cfg.GetPool(cfg.GetConfig())
	if err != nil {
		t.Fatalf("Failed to get database pool: %v", err)
	}

	namePattern := "K" // Example: Find members with a name similar to 'Muly'
	var jobSimilarName []entity.Job

	// Using LIKE operator for MySQL and selecting only the full_name field
	if err := pool.WithContext(ctx).Select("name").Where("name LIKE ?", "%"+namePattern+"%").Find(&jobSimilarName).Error; err != nil {
		log.Fatalf("Error querying the database: %v", err)
	}

	if len(jobSimilarName) == 0 {
		t.Logf("No family members found with similar name to: %s", namePattern)
	} else {
		fmt.Println(jobSimilarName)
	}
}
