package main

type GpsManager struct {
	dbManger *DbManager
	gps      map[string]*GpsInfo
	gpsById  map[int]*GpsInfo
}

func NewGpsManager(db *DbManager) *GpsManager {
	return &GpsManager{
		dbManger: db,
	}
}

func (gm *GpsManager) Load() {
	rows, err := gm.dbManger.Db.Query(
		"select " +
			"id,id_user,nopol,imei,phone,tdate,sdate,lat,lng" +
			" from view_device")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var gpsInfo = NewGpsInfo()
		err2 := rows.Scan(&gpsInfo.Id, &gpsInfo.Id_user, &gpsInfo.Nopol, &gpsInfo.Imei, &gpsInfo.Phone, &gpsInfo.Tdate, &gpsInfo.Sdate, &gpsInfo.Lat, &gpsInfo.Lng)

		if err2 != nil {
			continue
		}

		gm.gps[gpsInfo.Imei] = gpsInfo
	}

}
