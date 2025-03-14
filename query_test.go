package JARITMAS_API

//
//import (
//	"context"
//	"fmt"
//	"github.com/firstudio-lab/JARITMAS-API/cfg"
//	"testing"
//)
//
//type TransferKK struct {
//	//PEMOHON
//	KKRequest int64
//	//DATA KEPINDAHAN
//	Reason        string
//	ProvinceID    int
//	DistrictID    int
//	SubDistrictID int
//	VillageID     int
//	RT            string
//	RW            string
//	Address       string
//
//	FamilyMove    []string //ANNGOTA KELUARGA YG MOVE
//	NewHeadFamily int64    // JIIA TIDAK ADA NULL, JIKA ADA NANTI UPDATE NIK TERSEBUT JADI KEEPLA KELUARHA
//}
//
//type StatusMove string
//
//const (
//	AddToFamilyCard StatusMove = iota + 1
//	CreateNewFamilyCard
//	AllFamilyMove
//	NoMove
//)
//
//func (s StatusMove) ToString() string {
//	switch s {
//
//	}
//}
//
//func TestFIndAllMemberByKK(t *testing.T) {
//	ctx := context.Background()
//	pool, err := cfg.GetPool(cfg.GetConfig())
//	fmt.Println(err)
//
//}
