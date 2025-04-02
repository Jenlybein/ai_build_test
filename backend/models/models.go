package models

import (
	"time"
)

type LogQueryParams struct {
	StartTime    string `form:"start_time"`
	EndTime      string `form:"end_time"`
	UserID       string `form:"user_id"`
	AppID        string `form:"app_id"`
	ID           string `form:"id"`
	EntranceTime string `form:"entrance_time"`
	Platform     string `form:"platform"`
	PageSize     int    `form:"page_size,default=100"`
	PageIndex    int    `form:"page_index,default=1"`
}

// KV7 represents the structure of the kv_7 table
type KV7 struct {
	DataTime       *time.Time `ch:"data_time" json:"data_time"`
	WriteTime      *time.Time `ch:"write_time" json:"write_time"`
	TimeHour       string     `ch:"time_hour" json:"time_hour"`
	ID             string     `ch:"id" json:"id"`             // String 类型
	Time           int64      `ch:"time" json:"time"`           // Int64 类型
	Extra          string     `ch:"extra" json:"extra"`          // String 类型
	EntranceTime   int64      `ch:"entrance_time" json:"entrance_time"`   // Int64 类型
	EntranceID     string     `ch:"entrance_id" json:"entrance_id"`     // String 类型
	Stamp          int32      `ch:"stamp" json:"stamp"`          // Int32 类型
	AppID          string     `ch:"app_id" json:"app_id"`         // LowCardinality(String) 类型
	Platform       string     `ch:"platform" json:"platform"`       // LowCardinality(String) 类型
	UserID         string     `ch:"user_id" json:"user_id"`        // String 类型
	Version        string     `ch:"version" json:"version"`        // String 类型
	BuildID        string     `ch:"build_id" json:"build_id"`       // String 类型
	DeviceID       string     `ch:"device_id" json:"device_id"`      // String 类型
	Model          string     `ch:"model" json:"model"`          // String 类型
	OS             string     `ch:"os" json:"os"`             // LowCardinality(String) 类型
	OSVer          string     `ch:"os_ver" json:"os_ver"`         // String 类型
	SDKVer         string     `ch:"sdk_ver" json:"sdk_ver"`        // LowCardinality(String) 类型
	Category       string     `ch:"category" json:"category"`       // LowCardinality(String) 类型
	Action         string     `ch:"action" json:"action"`         // String 类型
	Label          string     `ch:"label" json:"label"`          // String 类型
	State          string     `ch:"state" json:"state"`          // String 类型
	Value          int32      `ch:"value" json:"value"`          // Int32 类型
	D1             string     `ch:"d1" json:"d1"`             // String 类型
	D2             string     `ch:"d2" json:"d2"`             // String 类型
	D3             string     `ch:"d3" json:"d3"`             // String 类型
	D4             string     `ch:"d4" json:"d4"`             // String 类型
	D5             string     `ch:"d5" json:"d5"`             // String 类型
	D6             string     `ch:"d6" json:"d6"`             // String 类型
	D7             string     `ch:"d7" json:"d7"`             // String 类型
	D8             string     `ch:"d8" json:"d8"`             // String 类型
	D9             string     `ch:"d9" json:"d9"`             // String 类型
	D10            string     `ch:"d10" json:"d10"`            // String 类型
	D11            string     `ch:"d11" json:"d11"`            // String 类型
	D12            string     `ch:"d12" json:"d12"`            // String 类型
	D13            string     `ch:"d13" json:"d13"`            // String 类型
	D14            string     `ch:"d14" json:"d14"`            // String 类型
	D15            string     `ch:"d15" json:"d15"`            // String 类型
	D16            string     `ch:"d16" json:"d16"`            // String 类型
	D17            string     `ch:"d17" json:"d17"`            // String 类型
	D18            string     `ch:"d18" json:"d18"`            // String 类型
	D19            string     `ch:"d19" json:"d19"`            // String 类型
	D20            string     `ch:"d20" json:"d20"`            // String 类型
	D21            string     `ch:"d21" json:"d21"`            // String 类型
	D22            string     `ch:"d22" json:"d22"`            // String 类型
	D23            string     `ch:"d23" json:"d23"`            // String 类型
	D24            string     `ch:"d24" json:"d24"`            // String 类型
	D25            string     `ch:"d25" json:"d25"`            // String 类型
	D26            string     `ch:"d26" json:"d26"`            // String 类型
	D27            string     `ch:"d27" json:"d27"`            // String 类型
	D28            string     `ch:"d28" json:"d28"`            // String 类型
	D29            string     `ch:"d29" json:"d29"`            // String 类型
	D30            string     `ch:"d30" json:"d30"`            // String 类型
	D31            string     `ch:"d31" json:"d31"`            // String 类型
	D32            string     `ch:"d32" json:"d32"`            // String 类型
	D33            string     `ch:"d33" json:"d33"`            // String 类型
	D34            string     `ch:"d34" json:"d34"`            // String 类型
	D35            string     `ch:"d35" json:"d35"`            // String 类型
	D36            string     `ch:"d36" json:"d36"`            // String 类型
	D37            string     `ch:"d37" json:"d37"`            // String 类型
	D38            string     `ch:"d38" json:"d38"`            // LowCardinality(String) 类型
	D39            string     `ch:"d39" json:"d39"`            // LowCardinality(String) 类型
	D40            string     `ch:"d40" json:"d40"`            // LowCardinality(String) 类型
	V1             int64      `ch:"v1" json:"v1"`             // Int64 类型
	V2             int64      `ch:"v2" json:"v2"`             // Int64 类型
	V3             int64      `ch:"v3" json:"v3"`             // Int64 类型
	V4             int64      `ch:"v4" json:"v4"`             // Int64 类型
	V5             int64      `ch:"v5" json:"v5"`             // Int64 类型
	V6             int64      `ch:"v6" json:"v6"`             // Int64 类型
	V7             int64      `ch:"v7" json:"v7"`             // Int64 类型
	V8             int64      `ch:"v8" json:"v8"`             // Int64 类型
	V9             int64      `ch:"v9" json:"v9"`             // Int64 类型
	V10            int64      `ch:"v10" json:"v10"`            // Int64 类型
	V11            int64      `ch:"v11" json:"v11"`            // Int64 类型
	V12            int64      `ch:"v12" json:"v12"`            // Int64 类型
	V13            int64      `ch:"v13" json:"v13"`            // Int64 类型
	V14            int64      `ch:"v14" json:"v14"`            // Int64 类型
	V15            int64      `ch:"v15" json:"v15"`            // Int64 类型
	V16            int64      `ch:"v16" json:"v16"`            // Int64 类型
	V17            int64      `ch:"v17" json:"v17"`            // Int64 类型
	V18            int64      `ch:"v18" json:"v18"`            // Int64 类型
	V19            int64      `ch:"v19" json:"v19"`            // Int64 类型
	V20            int64      `ch:"v20" json:"v20"`            // Int64 类型
	V21            int64      `ch:"v21" json:"v21"`            // Int64 类型
	V22            int64      `ch:"v22" json:"v22"`            // Int64 类型
	V23            int64      `ch:"v23" json:"v23"`            // Int64 类型
	V24            int64      `ch:"v24" json:"v24"`            // Int64 类型
	V25            int64      `ch:"v25" json:"v25"`            // Int64 类型
	V26            int64      `ch:"v26" json:"v26"`            // Int64 类型
	V27            int64      `ch:"v27" json:"v27"`            // Int64 类型
	V28            int64      `ch:"v28" json:"v28"`            // Int64 类型
	V29            int64      `ch:"v29" json:"v29"`            // Int64 类型
	V30            int64      `ch:"v30" json:"v30"`            // Int64 类型
	V31            int64      `ch:"v31" json:"v31"`            // Int64 类型
	V32            int64      `ch:"v32" json:"v32"`            // Int64 类型
	V33            int64      `ch:"v33" json:"v33"`            // Int64 类型
	V34            int64      `ch:"v34" json:"v34"`            // Int64 类型
	V35            int64      `ch:"v35" json:"v35"`            // Int64 类型
	V36            int64      `ch:"v36" json:"v36"`            // Int64 类型
	V37            int64      `ch:"v37" json:"v37"`            // Int64 类型
	V38            int64      `ch:"v38" json:"v38"`            // Int64 类型
	V39            int64      `ch:"v39" json:"v39"`            // Int64 类型
	V40            int64      `ch:"v40" json:"v40"`            // Int64 类型
	Info1          string     `ch:"info1" json:"info1"`          // String 类型
	Info2          string     `ch:"info2" json:"info2"`          // String 类型
	Info3          string     `ch:"info3" json:"info3"`          // String 类型
	Info4          string     `ch:"info4" json:"info4"`          // String 类型
	Info5          string     `ch:"info5" json:"info5"`          // String 类型
	Info6          string     `ch:"info6" json:"info6"`          // String 类型
	Info7          string     `ch:"info7" json:"info7"`          // String 类型
	Info8          string     `ch:"info8" json:"info8"`          // String 类型
	Info9          string     `ch:"info9" json:"info9"`          // String 类型
	Info10         string     `ch:"info10" json:"info10"`         // String 类型
	UD1            string     `ch:"ud1" json:"ud1"`            // String 类型
	UD2            string     `ch:"ud2" json:"ud2"`            // String 类型
	UD3            string     `ch:"ud3" json:"ud3"`            // String 类型
	UD4            string     `ch:"ud4" json:"ud4"`            // String 类型
	UD5            string     `ch:"ud5" json:"ud5"`            // String 类型
	UD6            string     `ch:"ud6" json:"ud6"`            // String 类型
	UD7            string     `ch:"ud7" json:"ud7"`            // String 类型
	UD8            string     `ch:"ud8" json:"ud8"`            // String 类型
	UD9            string     `ch:"ud9" json:"ud9"`            // String 类型
	UD10           string     `ch:"ud10" json:"ud10"`           // String 类型
	UD11           string     `ch:"ud11" json:"ud11"`           // String 类型	
	UD12           string     `ch:"ud12" json:"ud12"`           // String 类型
	UD13           string     `ch:"ud13" json:"ud13"`           // String 类型
	UD14           string     `ch:"ud14" json:"ud14"`           // String 类型
	UD15           string     `ch:"ud15" json:"ud15"`           // String 类型
	UD16           string     `ch:"ud16" json:"ud16"`           // String 类型
	UD17           string     `ch:"ud17" json:"ud17"`           // String 类型
	UD18           string     `ch:"ud18" json:"ud18"`           // String 类型
	UD19           string     `ch:"ud19" json:"ud19"`           // String 类型
	UD20           string     `ch:"ud20" json:"ud20"`           // String 类型
	UV1            int64      `ch:"uv1" json:"uv1"`            // Int64 类型
	UV2            int64      `ch:"uv2" json:"uv2"`            // Int64 类型
	UV3            int64      `ch:"uv3" json:"uv3"`            // Int64 类型
	UV4            int64      `ch:"uv4" json:"uv4"`            // Int64 类型
	UV5            int64      `ch:"uv5" json:"uv5"`            // Int64 类型
	UV6            int64      `ch:"uv6" json:"uv6"`            // Int64 类型
	UV7            int64      `ch:"uv7" json:"uv7"`            // Int64 类型
	UV8            int64      `ch:"uv8" json:"uv8"`            // Int64 类型
	UV9            int64      `ch:"uv9" json:"uv9"`            // Int64 类型
	UV10           int64      `ch:"uv10" json:"uv10"`           // Int64 类型
	SD1            string     `ch:"sd1" json:"sd1"`            // String 类型
	SD2            string     `ch:"sd2" json:"sd2"`            // String 类型
	SD3            string     `ch:"sd3" json:"sd3"`            // String 类型
	SD4            string     `ch:"sd4" json:"sd4"`            // String 类型
	SD5            string     `ch:"sd5" json:"sd5"`            // String 类型
	SD6            string     `ch:"sd6" json:"sd6"`            // String 类型
	SD7            string     `ch:"sd7" json:"sd7"`            // String 类型
	SD8            string     `ch:"sd8" json:"sd8"`            // String 类型
	SD9            string     `ch:"sd9" json:"sd9"`            // String 类型
	SD10           string     `ch:"sd10" json:"sd10"`           // String 类型
	SD11           string     `ch:"sd11" json:"sd11"`           // String 类型
	SD12           string     `ch:"sd12" json:"sd12"`           // String 类型
	SD13           string     `ch:"sd13" json:"sd13"`           // String 类型
	SD14           string     `ch:"sd14" json:"sd14"`           // String 类型
	SD15           string     `ch:"sd15" json:"sd15"`           // String 类型
	SD16           string     `ch:"sd16" json:"sd16"`           // String 类型
	SD17           string     `ch:"sd17" json:"sd17"`           // String 类型
	SD18           string     `ch:"sd18" json:"sd18"`           // String 类型
	SD19           string     `ch:"sd19" json:"sd19"`           // String 类型
	SD20           string     `ch:"sd20" json:"sd20"`           // String 类型
	SV1            int64      `ch:"sv1" json:"sv1"`            // Int64 类型
	SV2            int64      `ch:"sv2" json:"sv2"`            // Int64 类型
	SV3            int64      `ch:"sv3" json:"sv3"`            // Int64 类型
	SV4            int64      `ch:"sv4" json:"sv4"`            // Int64 类型
	SV5            int64      `ch:"sv5" json:"sv5"`            // Int64 类型
	SV6            int64      `ch:"sv6" json:"sv6"`            // Int64 类型
	SV7            int64      `ch:"sv7" json:"sv7"`            // Int64 类型
	SV8            int64      `ch:"sv8" json:"sv8"`            // Int64 类型
	SV9            int64      `ch:"sv9" json:"sv9"`            // Int64 类型
	SV10           int64      `ch:"sv10" json:"sv10"`           // Int64 类型
}