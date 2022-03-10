package main

type GpsInfo struct {
	Id      int
	Id_user int
	Nopol   string
	Imei    string
	Phone   string
	Tdate   string
	Sdate   string
	Lat     float64
	Lng     float64
}

func NewGpsInfo() *GpsInfo {
	return &GpsInfo{}
}
