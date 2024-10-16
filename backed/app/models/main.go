package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type SelectMapInfoResponse struct {
	Status   string    `json:"status"`
	Info     string    `json:"info"`
	Infocode string    `json:"infocode"`
	Count    string    `json:"count"`
	Geocodes []Geocode `json:"geocodes"`
}

type Geocode struct {
	FormattedAddress string   `json:"formatted_address"`
	Country          string   `json:"country"`
	Province         string   `json:"province"`
	City             string   `json:"city"`
	Township         []string `json:"township"`
	Neighborhood     NameType `json:"neighborhood"`
	Building         NameType `json:"building"`
	Adcode           string   `json:"adcode"`
	Location         string   `json:"location"`
	Level            string   `json:"level"`
}

type NameType struct {
	Name []string `json:"name"`
	Type []string `json:"type"`
}

// LogisticsInfo 物流信息表
type LogisticsInfo struct {
	gorm.Model
	PackageId           string `json:"package_id"`
	Email               string `json:"email"`
	SendPerson          string `json:"send_person"`           // "发货人"
	SendPhoneNumber     string `json:"send_phone_number"`     // "发货人手机号"
	SendDetailedAddress string `json:"send_detailed_address"` // "发货人详细地址"
	SendCompanyName     string `json:"send_company_name"`     // "发货人公司名称公司名称"
	ReceivePerson       string `json:"receive_person"`        // "收货人"
	ReceivePhoneNumber  string `json:"receive_phone_number"`  // "收货人手机号"
	ReceiveAddress      string `json:"receive_address"`       // "收货人地址"
	ReceiveCompanyName  string `json:"receive_company_name"`  // "收货人公司名称"
	TradeNo             string `json:"trade_no"`              // 支付订单号
	Status              int64  `json:"status"`                // 状态 0 待支付 1 待发货 2 待收货 3 已完成 4 退款成功
}

// MessageInfo 消息
type MessageInfo struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

func (m *MessageInfo) Marshal() string {
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

// Unmarshal 反序列化
func (m *MessageInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}
