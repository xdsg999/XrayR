package pmpanel

import "encoding/json"

// NodeInfoResponse is the response of node
type NodeInfoResponse struct {
	Group            int     `json:"node_group"`
	Class            int     `json:"node_class"`
	SpeedLimit       int     `json:"node_speedlimit"`
	Method           int     `json:"method"`
	TrafficRate      int     `json:"trafficrate"`
	MuOnly           int     `json:"mu_only"`
	Sort             int     `json:"sort"`
	RawServerString  string  `json:"server"`
	Type             string  `json:"type"`
}

// UserResponse is the response of user
type UserResponse struct {
	ID             int     `json:"id"`
	EMail          string  `json:"email"`
	Passwd         string  `json:"passwd"`
	Port           int     `json:"port"`
	Method         string  `json:"method"`
	SpeedLimit     float64 `json:"speedlimt"`
	DeviceLimit    int     `json:"devicelimit"`
	Protocol       string  `json:"protocol"`
	ProtocolParam  string  `json:"protocol_param"`
	Obfs           string  `json:"obfs"`
	ObfsParam      string  `json:"obfs_param"`
	ForbiddenIP    string  `json:"forbidden_ip"`
	Forbiddenport  string  `json:"forbidden_port"`
	UUID           string  `json:"uuid"`
	MultiUser      string  `json:"is_multi_user"`
}

// Response is the common response
type Response struct {
	Ret  uint             `json:"ret"`
	Data json.RawMessage  `json:"date"`
}

// PostData is the data structure of post data
type PostData struct {
	Data interface{} `json:"data"`
}

// systemload is the data structure of systemload
type Systemload struct {
	Uptime string `json:"uptime"`
	Load   string `json:"load"`
}

// onlineuser is the data structure of online uer
type OnlineUser struct {
	UID int       `json:"user_id"`
	IP string     `josn:"ip"`  
}

// usertraffic is the data structure of traffic
type UserTraffic struct {
	UID       int    `json:"user_id"`
	Upload    int64  `json:"u"`
	Download  int64  `json:"d"`
}

type RuleItem struct {
	ID         int      `json:"id"`
	Content    string   `json:"regex"`
}

type IllegalItem struct {
	ID  int  `json:"list_id"`
	UID int  `json:"user_id"`
}