# golang NASA api
## 简介
封装api
## 使用
```
go get github.com/wms3001/goNASA
```
## 实例
1. 宇宙每日一图
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
var ap apod.Apod
ap.AdReq.Date = "2023-05-26"
goNASA.DayAstronomyPicture(&ap)
tt, _ := json.Marshal(ap)
log.Println(string(tt))
```
2. NEO list
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
var ne neo.NEO
ne.NeoReq.Start_date = "2023-05-20"
ne.NeoReq.End_date = "2023-05-27"
goNASA.NeoFeed(&ne)
tt, _ := json.Marshal(ne)
log.Println(string(tt))
```
3. 获取单个NEO数据
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.NeoLookup("3542519")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
4. 获取NEOlist
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.NeoBrowse("10", "5")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
5. 日冕物质抛射
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.CME("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
6. 地磁风暴
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.GeomagneticStorm("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
7. 行星际冲击
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
var ipsReq cme.IPSReq
ipsReq.StartDate = "2023-05-01"
ipsReq.EndDate = "2023-05-30"
ipsReq.Location = "ALL"
ipsReq.Catalog = "ALL"
re := goNASA.InterplanetaryShock(ipsReq)
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
8. 太阳耀斑
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.SolarFlare("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
9. 太阳高能粒子
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.SolarEnergeticParticle("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
10. 磁层顶穿越
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.MagnetopauseCrossing("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
11. 辐射带增强
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.RadiationBeltEnhancement("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
12. Hight Speed Stream
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.HightSpeedStream("2023-05-01", "2023-05-10")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
13. Notifications
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.Notifications("2023-05-01", "2023-05-10", "ALL")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
14. Imagery
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.Imagery("2014-02-01", "100.75", "1.5")
tt, _ := json.Marshal(re)
log.Println(tt)
```
15. Assets
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://api.nasa.gov",
}
re := goNASA.Assets("2014-02-01", "-95.33", "29.78", "0.10")
tt, _ := json.Marshal(re)
log.Println(tt)
```
16. Categories
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.Categories_v2()
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
17. Category
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
var cateReq eonet.CategoryReq
cateReq.CategoryId = "8"
cateReq.Source = "InciWeb,EO"
cateReq.Status = "open"
re := goNASA.Category(cateReq)
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
18. Layers
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.Layers()
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
19. Layer
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.Layer("8")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
20. Events
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
var eventReq eonet.EventReq
eventReq.Source = "InciWeb"
re := goNASA.Events(eventReq)
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
21. Epic
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.EPIC("2023-05-30")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
22. mars photos by sole
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.MarsPhotosBySole("1000", "", "")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
23. mars photos by date
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.MarsPhotosBySole("2015-06-03", "", "")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
24. mars weather
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://eonet.gsfc.nasa.gov",
}
re := goNASA.MarsWeather("json", "1.0")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
25. imagesearch
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://images-api.nasa.gov",
}
re := goNASA.ImagesSearch("mars")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
26. ImagesAsset
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://images-api.nasa.gov",
}
re := goNASA.ImagesAsset("NHQ201905310026")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```
27. ImageMetadata
```go
var goNASA GoNASA = GoNASA{
Token: " ",
Url:   "https://images-api.nasa.gov",
}
re := goNASA.ImageMetadata("NHQ201905310026")
tt, _ := json.Marshal(re)
log.Println(string(tt))
```