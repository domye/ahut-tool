package models

// IMSResponse IMS响应结构体
type IMSResponse struct {
	Code int     `json:"Code"`
	Msg  string  `json:"Msg"`
	Data IMSData `json:"Data"`
}

// IMSData IMS数据结构体
type IMSData struct {
	RoomID    string  `json:"room_id"`
	AllAmp    float64 `json:"AllAmp"`
	UsedAmp   float64 `json:"UsedAmp"`
	RemainAmp float64 `json:"RemainAmp"`
}
