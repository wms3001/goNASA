package goNASA

import (
	"encoding/json"
	"github.com/wms3001/goNASA/apod"
	"github.com/wms3001/goNASA/cme"
	"github.com/wms3001/goNASA/eonet"
	"github.com/wms3001/goNASA/neo"
	"log"
	"testing"
)

var goNASA GoNASA = GoNASA{
	Token: "H67QEefLCw13HzP6L1ggqOeQgFbI69aVTdVsFJxN",
	Url:   "https://api.nasa.gov",
}

func TestGoNASA_DayAstronomyPicture(t *testing.T) {
	var ap apod.Apod
	ap.AdReq.Date = "2023-05-26"
	goNASA.DayAstronomyPicture(&ap)
	tt, _ := json.Marshal(ap)
	log.Println(string(tt))
}

func TestGoNASA_NeoFeed(t *testing.T) {
	var ne neo.NEO
	ne.NeoReq.Start_date = "2023-05-20"
	ne.NeoReq.End_date = "2023-05-27"
	goNASA.NeoFeed(&ne)
	tt, _ := json.Marshal(ne)
	log.Println(string(tt))
}

func TestGoNASA_NeoLookup(t *testing.T) {
	re := goNASA.NeoLookup("3542519")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_NeoBrowse(t *testing.T) {
	re := goNASA.NeoBrowse("10", "5")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_CME(t *testing.T) {
	re := goNASA.CME("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_CMEAnalysis(t *testing.T) {
	var cmea cme.CMEAnalysy
	cmea.StartDate = "2023-05-01"
	cmea.EndDate = "2023-05-10"
	cmea.MostAccurateOnly = "true"
	cmea.Catalog = "ALL"
	cmea.Speed = "500"
	cmea.HalfAngle = "30"
	re := goNASA.CMEAnalysis(cmea)
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_GeomagneticStorm(t *testing.T) {
	re := goNASA.GeomagneticStorm("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_InterplanetaryShock(t *testing.T) {
	var ipsReq cme.IPSReq
	ipsReq.StartDate = "2023-05-01"
	ipsReq.EndDate = "2023-05-30"
	ipsReq.Location = "ALL"
	ipsReq.Catalog = "ALL"
	re := goNASA.InterplanetaryShock(ipsReq)
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_SolarFlare(t *testing.T) {
	re := goNASA.SolarFlare("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_SolarEnergeticParticle(t *testing.T) {
	re := goNASA.SolarEnergeticParticle("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_MagnetopauseCrossing(t *testing.T) {
	re := goNASA.MagnetopauseCrossing("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_RadiationBeltEnhancement(t *testing.T) {
	re := goNASA.RadiationBeltEnhancement("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_HightSpeedStream(t *testing.T) {
	re := goNASA.HightSpeedStream("2023-05-01", "2023-05-10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Notifications(t *testing.T) {
	re := goNASA.Notifications("2023-05-01", "2023-05-10", "ALL")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Imagery(t *testing.T) {
	re := goNASA.Imagery("2014-02-01", "100.75", "1.5")
	tt, _ := json.Marshal(re)
	log.Println(tt)
}

func TestGoNASA_Assets(t *testing.T) {
	re := goNASA.Assets("2014-02-01", "-95.33", "29.78", "0.10")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Categories_v2(t *testing.T) {
	goNASA.Url = "https://eonet.gsfc.nasa.gov"
	re := goNASA.Categories_v2()
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Category(t *testing.T) {
	goNASA.Url = "https://eonet.gsfc.nasa.gov"
	var cateReq eonet.CategoryReq
	cateReq.CategoryId = "8"
	cateReq.Source = "InciWeb,EO"
	cateReq.Status = "open"
	re := goNASA.Category(cateReq)
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Layers(t *testing.T) {
	goNASA.Url = "https://eonet.gsfc.nasa.gov"
	re := goNASA.Layers()
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Layer(t *testing.T) {
	goNASA.Url = "https://eonet.gsfc.nasa.gov"
	re := goNASA.Layer("8")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_Events(t *testing.T) {
	goNASA.Url = "https://eonet.gsfc.nasa.gov"
	var eventReq eonet.EventReq
	eventReq.Source = "InciWeb"
	re := goNASA.Events(eventReq)
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_EPIC(t *testing.T) {
	re := goNASA.EPIC("2023-05-30")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_MarsPhotosBySole(t *testing.T) {
	re := goNASA.MarsPhotosBySole("1000", "", "")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_MarsPhotosByDate(t *testing.T) {
	re := goNASA.MarsPhotosBySole("2015-06-03", "", "")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_MarsWeather(t *testing.T) {
	re := goNASA.MarsWeather("json", "1.0")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_ImagesSearch(t *testing.T) {
	goNASA.Url = "https://images-api.nasa.gov"
	re := goNASA.ImagesSearch("mars")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_ImagesAsset(t *testing.T) {
	goNASA.Url = "https://images-api.nasa.gov"
	re := goNASA.ImagesAsset("NHQ201905310026")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}

func TestGoNASA_ImageMetadata(t *testing.T) {
	goNASA.Url = "https://images-api.nasa.gov"
	re := goNASA.ImageMetadata("NHQ201905310026")
	tt, _ := json.Marshal(re)
	log.Println(string(tt))
}
