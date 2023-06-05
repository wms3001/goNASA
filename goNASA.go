package goNASA

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/wms3001/goNASA/apod"
	"github.com/wms3001/goNASA/base"
	"github.com/wms3001/goNASA/cme"
	"github.com/wms3001/goNASA/donki"
	"github.com/wms3001/goNASA/earth"
	"github.com/wms3001/goNASA/eonet"
	"github.com/wms3001/goNASA/epic"
	"github.com/wms3001/goNASA/images"
	"github.com/wms3001/goNASA/insight"
	"github.com/wms3001/goNASA/marsPhotos"
	"github.com/wms3001/goNASA/neo"
	"strings"
)

type GoNASA struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

func (goNASA *GoNASA) DayAstronomyPicture(apo *apod.Apod) {
	url := goNASA.Url + "/planetary/apod?date=" + apo.AdReq.Date +
		"&start_date=" + apo.AdReq.Start_date +
		"&end_date=" + apo.AdReq.End_date +
		"&count=" + apo.AdReq.Count +
		"&thumbs=" + apo.AdReq.Thumbs +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	apRes := apod.ApdRes{}
	json.Unmarshal([]byte(res), &apRes)
	apo.AdRes = apRes
}

func (goNASA *GoNASA) NeoFeed(ne *neo.NEO) {
	url := goNASA.Url + "/neo/rest/v1/feed?" +
		"start_date=" + ne.NeoReq.Start_date +
		"&end_date=" + ne.NeoReq.End_date +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	neRes := neo.NeRes{}
	json.Unmarshal([]byte(res), &neRes)
	ne.NeoRes = neRes
}

func (goNASA *GoNASA) NeoLookup(neoId string) neo.NeoLink {
	url := goNASA.Url + "/neo/rest/v1/neo/" + neoId + "?" +
		"api_key=" + goNASA.Token
	res := goNASA.Get(url)
	neoLink := neo.NeoLink{}
	json.Unmarshal([]byte(res), &neoLink)
	return neoLink
}

func (goNASA *GoNASA) NeoBrowse(page string, pageSize string) neo.NeoLinkList {
	url := goNASA.Url + "/neo/rest/v1/neo/browse?" +
		"page=" + page + "&size=" + pageSize +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	neoLinkList := neo.NeoLinkList{}
	json.Unmarshal([]byte(res), &neoLinkList)
	return neoLinkList
}

func (goNASA *GoNASA) CME(startDate string, endDate string) []donki.CMEActive {
	url := goNASA.Url + "/DONKI/CME?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []donki.CMEActive
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) CMEAnalysis(cmea cme.CMEAnalysy) []cme.CMEAnaly {
	url := goNASA.Url + "/DONKI/CME?" +
		"startDate=" + cmea.StartDate + "&endDate=" + cmea.EndDate +
		"&mostAccurateOnly=" + cmea.MostAccurateOnly +
		"&speed=" + cmea.Speed +
		"&halfAngle=" + cmea.HalfAngle +
		"&catalog=" + cmea.Catalog +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.CMEAnaly
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) GeomagneticStorm(startDate string, endDate string) []cme.GST {
	url := goNASA.Url + "/DONKI/GST?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.GST
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) InterplanetaryShock(ipsReq cme.IPSReq) []cme.IPS {
	url := goNASA.Url + "/DONKI/IPS?" +
		"startDate=" + ipsReq.StartDate + "&endDate=" + ipsReq.EndDate +
		"&location=" + ipsReq.Location +
		"&catalog=" + ipsReq.Catalog +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.IPS
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) SolarFlare(startDate string, endDate string) []cme.FLR {
	url := goNASA.Url + "/DONKI/FLR?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.FLR
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) SolarEnergeticParticle(startDate string, endDate string) []cme.SEP {
	url := goNASA.Url + "/DONKI/SEP?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.SEP
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) MagnetopauseCrossing(startDate string, endDate string) []cme.MPC {
	url := goNASA.Url + "/DONKI/MPC?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.MPC
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) RadiationBeltEnhancement(startDate string, endDate string) []cme.RBE {
	url := goNASA.Url + "/DONKI/RBE?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.RBE
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) HightSpeedStream(startDate string, endDate string) []cme.HSS {
	url := goNASA.Url + "/DONKI/HSS?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.HSS
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) Notifications(startDate string, endDate string, ty string) []cme.Note {
	url := goNASA.Url + "/DONKI/notifications?" +
		"startDate=" + startDate + "&endDate=" + endDate +
		"&type=" + ty +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var cms []cme.Note
	json.Unmarshal([]byte(res), &cms)
	return cms
}

func (goNASA *GoNASA) Imagery(date string, lon string, lat string) string {
	url := goNASA.Url + "/planetary/earth/imagery?" +
		"date=" + date + "&lon=" + lon +
		"&lat=" + lat +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	return res
}

func (goNASA *GoNASA) Assets(date string, lon string, lat string, dim string) earth.Assets {
	url := goNASA.Url + "/planetary/earth/assets?" +
		"date=" + date + "&lon=" + lon +
		"&lat=" + lat +
		"&dim=" + dim +
		"&api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var assets earth.Assets
	json.Unmarshal([]byte(res), &assets)
	return assets
}

func (goNASA *GoNASA) Categories_v2() eonet.CategoriesInfo {
	url := goNASA.Url + "/api/v2.1/categories"
	res := goNASA.Get(url)
	var cates eonet.CategoriesInfo
	json.Unmarshal([]byte(res), &cates)
	return cates
}

func (goNASA *GoNASA) Category(cate eonet.CategoryReq) eonet.Category {
	url := goNASA.Url + "/api/v2.1/categories/" + cate.CategoryId + "?" +
		"source=" + cate.Source +
		"&status=" + cate.Status +
		"&limit=" + cate.Limit +
		"&days=" + cate.Days
	res := goNASA.Get(url)
	var cates eonet.Category
	json.Unmarshal([]byte(res), &cates)
	return cates
}

func (goNASA *GoNASA) Layers() eonet.Layers {
	url := goNASA.Url + "/api/v2.1/layers"
	res := goNASA.Get(url)
	var layers eonet.Layers
	json.Unmarshal([]byte(res), &layers)
	return layers
}

func (goNASA *GoNASA) Layer(id string) eonet.Layer {
	url := goNASA.Url + "/api/v2.1/layers/" + id
	res := goNASA.Get(url)
	var layer eonet.Layer
	json.Unmarshal([]byte(res), &layer)
	return layer
}

func (goNASA *GoNASA) Events(req eonet.EventReq) eonet.Events {
	url := goNASA.Url + "/api/v2.1/events?" +
		"source=" + req.Source +
		"&status=" + req.Status +
		"&limit=" + req.Limit +
		"&days=" + req.Days
	res := goNASA.Get(url)
	var events eonet.Events
	json.Unmarshal([]byte(res), &events)
	return events
}

func (goNASA *GoNASA) EPIC(day string) []epic.Epic {
	url := goNASA.Url + "/EPIC/api/natural/date/" + day + "?" +
		"api_key=" + goNASA.Token
	res := goNASA.Get(url)
	var epc []epic.Epic
	json.Unmarshal([]byte(res), &epc)
	for key, value := range epc {
		d1 := strings.Split(value.Date, " ")
		d2 := strings.Split(d1[0], "-")
		imageUrl := "https://api.nasa.gov/EPIC/archive/natural/" + d2[0] + "/" + d2[1] + "/" + d2[2] + "/png/" + value.Image + ".png?api_key=" + goNASA.Token
		epc[key].ImageUrl = imageUrl
	}
	return epc
}

func (goNASA *GoNASA) MarsPhotosBySole(sol string, camera string, page string) marsPhotos.MarsPhotos {
	url := goNASA.Url + "/mars-photos/api/v1/rovers/curiosity/photos?" + "api_key=" + goNASA.Token
	url += "&sol=" + sol
	if camera != "" {
		url += "&camera=" + camera
	}
	if page != "" {
		url += "&page=" + page
	}
	res := goNASA.Get(url)
	var marsPhotos marsPhotos.MarsPhotos
	json.Unmarshal([]byte(res), &marsPhotos)
	return marsPhotos
}

func (goNASA *GoNASA) MarsPhotosByDate(date string, camera string, page string) marsPhotos.MarsPhotos {
	url := goNASA.Url + "/mars-photos/api/v1/rovers/curiosity/photos?" + "api_key=" + goNASA.Token
	url += "&earth_date=" + date
	if camera != "" {
		url += "&camera=" + camera
	}
	if page != "" {
		url += "&page=" + page
	}
	res := goNASA.Get(url)
	var marsPhotos marsPhotos.MarsPhotos
	json.Unmarshal([]byte(res), &marsPhotos)
	return marsPhotos
}

func (goNASA *GoNASA) MarsWeather(feedtype string, ver string) insight.WarsWeather {
	url := goNASA.Url + "/insight_weather/?" + "api_key=" + goNASA.Token
	url += "&feedtype=" + feedtype
	url += "&ver=" + ver
	res := goNASA.Get(url)
	var marsPhotos insight.WarsWeather
	json.Unmarshal([]byte(res), &marsPhotos)
	return marsPhotos
}

func (goNASA *GoNASA) ImagesSearch(q string) images.NASAImage {
	url := goNASA.Url + "/search?q=" + q
	res := goNASA.Get(url)
	var nasaImage images.NASAImage
	json.Unmarshal([]byte(res), &nasaImage)
	return nasaImage
}

func (goNASA *GoNASA) ImagesAsset(id string) images.ImageAsset {
	url := goNASA.Url + "/asset/" + id
	res := goNASA.Get(url)
	var nasaImage images.ImageAsset
	json.Unmarshal([]byte(res), &nasaImage)
	return nasaImage
}

func (goNASA *GoNASA) ImageMetadata(id string) images.Metadata {
	url := goNASA.Url + "/metadata/" + id
	res := goNASA.Get(url)
	var nasaImage images.Metadata
	json.Unmarshal([]byte(res), &nasaImage)
	return nasaImage
}

func (goNASA *GoNASA) TechTransfer() Teah {
	url := goNASA.Url + "/techtransfer/patent/?engine&patent&patent_issued&software&Spinoff"
	res := goNASA.Get(url)
	var nasaImage Teah
	json.Unmarshal([]byte(res), &nasaImage)
	return nasaImage
}

func (goNASA *GoNASA) Get(url string) string {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		response := base.NASARsp{}
		response.Code = -1
		response.ServiceVersion = "v1"
		response.Msg = err.Error()
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}

func (goNASA *GoNASA) PostJson(url string, body string) string {
	client := resty.New()
	client.R().SetHeader("Content-Type", "application/json")
	if body != "" {
		client.R().SetBody(body)
	}
	resp, err := client.R().Post(url)
	if err != nil {
		response := base.NASARsp{}
		response.Code = -1
		response.ServiceVersion = "v1"
		response.Msg = err.Error()
		str, _ := json.Marshal(response)
		return string(str)
	} else {
		return string(resp.Body())
	}
}
