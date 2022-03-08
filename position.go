package main

import "time"

/*
* GPS Positions
 */
type Position struct {
	Id            int
	Lat           float64
	Lng           float64
	Speed         float32
	Angle         uint16
	TDate         time.Time
	SDate         time.Time
	Sat           uint8
	Acc           uint8
	Charge        uint8
	Batt          uint8
	Fcut          uint8
	GF_ID         uint16 //Geofence ID
	GF_Name       string //Geofence Name
	POI_ID        uint16 //Point of Interest ID
	Poi           string //Point of Interest Name
	Address       string //Region, Street
	Park_Info     string
	ParkTimestamp uint64
}

func NewPosition() *Position {
	return &Position{}
}
func Parse(p *Position) {

}
