package bitmex

import (
	"bitmex-status-cli/io"
	"bitmex-status-cli/network"
	"encoding/json"
	"time"
)

const APIURL = "https://www.bitmex.com"

func GetPositions() (positions Positions) {

	endpoint := "/api/v1/position"
	keys := io.ReadKeyFile("bitmex.apikey")

	result := network.PerformAPICall(keys, "GET", endpoint, APIURL)
	bytevalue := []byte(result)
	json.Unmarshal(bytevalue, &positions)
	return

}

func GetOrders() (orders Orders) {

	endpoint := "/api/v1/order?filter=%7B%22open%22%3A%20true%7D"
	keys := io.ReadKeyFile("bitmex.apikey")

	result := network.PerformAPICall(keys, "GET", endpoint, APIURL)
	bytevalue := []byte(result)
	json.Unmarshal(bytevalue, &orders)
	return

}

type Positions []struct {
	Account              int         `json:"account"`
	Symbol               string      `json:"symbol"`
	Currency             string      `json:"currency"`
	Underlying           string      `json:"underlying"`
	QuoteCurrency        string      `json:"quoteCurrency"`
	Commission           float64     `json:"commission"`
	InitMarginReq        float64     `json:"initMarginReq"`
	MaintMarginReq       float64     `json:"maintMarginReq"`
	RiskLimit            int64       `json:"riskLimit"`
	Leverage             int         `json:"leverage"`
	CrossMargin          bool        `json:"crossMargin"`
	DeleveragePercentile int         `json:"deleveragePercentile"`
	RebalancedPnl        int         `json:"rebalancedPnl"`
	PrevRealisedPnl      int         `json:"prevRealisedPnl"`
	PrevUnrealisedPnl    int         `json:"prevUnrealisedPnl"`
	PrevClosePrice       float64     `json:"prevClosePrice"`
	OpeningTimestamp     time.Time   `json:"openingTimestamp"`
	OpeningQty           int         `json:"openingQty"`
	OpeningCost          int         `json:"openingCost"`
	OpeningComm          int         `json:"openingComm"`
	OpenOrderBuyQty      int         `json:"openOrderBuyQty"`
	OpenOrderBuyCost     int         `json:"openOrderBuyCost"`
	OpenOrderBuyPremium  int         `json:"openOrderBuyPremium"`
	OpenOrderSellQty     int         `json:"openOrderSellQty"`
	OpenOrderSellCost    int         `json:"openOrderSellCost"`
	OpenOrderSellPremium int         `json:"openOrderSellPremium"`
	ExecBuyQty           int         `json:"execBuyQty"`
	ExecBuyCost          int         `json:"execBuyCost"`
	ExecSellQty          int         `json:"execSellQty"`
	ExecSellCost         int         `json:"execSellCost"`
	ExecQty              int         `json:"execQty"`
	ExecCost             int         `json:"execCost"`
	ExecComm             int         `json:"execComm"`
	CurrentTimestamp     time.Time   `json:"currentTimestamp"`
	CurrentQty           int         `json:"currentQty"`
	CurrentCost          int         `json:"currentCost"`
	CurrentComm          int         `json:"currentComm"`
	RealisedCost         int         `json:"realisedCost"`
	UnrealisedCost       int         `json:"unrealisedCost"`
	GrossOpenCost        int         `json:"grossOpenCost"`
	GrossOpenPremium     int         `json:"grossOpenPremium"`
	GrossExecCost        int         `json:"grossExecCost"`
	IsOpen               bool        `json:"isOpen"`
	MarkPrice            float64     `json:"markPrice"`
	MarkValue            int         `json:"markValue"`
	RiskValue            int         `json:"riskValue"`
	HomeNotional         float64     `json:"homeNotional"`
	ForeignNotional      int         `json:"foreignNotional"`
	PosState             string      `json:"posState"`
	PosCost              int         `json:"posCost"`
	PosCost2             int         `json:"posCost2"`
	PosCross             int         `json:"posCross"`
	PosInit              int         `json:"posInit"`
	PosComm              int         `json:"posComm"`
	PosLoss              int         `json:"posLoss"`
	PosMargin            int         `json:"posMargin"`
	PosMaint             int         `json:"posMaint"`
	PosAllowance         int         `json:"posAllowance"`
	TaxableMargin        int         `json:"taxableMargin"`
	InitMargin           int         `json:"initMargin"`
	MaintMargin          int         `json:"maintMargin"`
	SessionMargin        int         `json:"sessionMargin"`
	TargetExcessMargin   int         `json:"targetExcessMargin"`
	VarMargin            int         `json:"varMargin"`
	RealisedGrossPnl     int         `json:"realisedGrossPnl"`
	RealisedTax          int         `json:"realisedTax"`
	RealisedPnl          int         `json:"realisedPnl"`
	UnrealisedGrossPnl   int         `json:"unrealisedGrossPnl"`
	LongBankrupt         int         `json:"longBankrupt"`
	ShortBankrupt        int         `json:"shortBankrupt"`
	TaxBase              int         `json:"taxBase"`
	IndicativeTaxRate    int         `json:"indicativeTaxRate"`
	IndicativeTax        int         `json:"indicativeTax"`
	UnrealisedTax        int         `json:"unrealisedTax"`
	UnrealisedPnl        int         `json:"unrealisedPnl"`
	UnrealisedPnlPcnt    float64     `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt    float64     `json:"unrealisedRoePcnt"`
	SimpleQty            interface{} `json:"simpleQty"`
	SimpleCost           interface{} `json:"simpleCost"`
	SimpleValue          interface{} `json:"simpleValue"`
	SimplePnl            interface{} `json:"simplePnl"`
	SimplePnlPcnt        interface{} `json:"simplePnlPcnt"`
	AvgCostPrice         int         `json:"avgCostPrice"`
	AvgEntryPrice        int         `json:"avgEntryPrice"`
	BreakEvenPrice       float64     `json:"breakEvenPrice"`
	MarginCallPrice      int         `json:"marginCallPrice"`
	LiquidationPrice     int         `json:"liquidationPrice"`
	BankruptPrice        int         `json:"bankruptPrice"`
	Timestamp            time.Time   `json:"timestamp"`
	LastPrice            float64     `json:"lastPrice"`
	LastValue            int         `json:"lastValue"`
}

type Orders []struct {
	OrderID               string    `json:"orderID"`
	ClOrdID               string    `json:"clOrdID"`
	ClOrdLinkID           string    `json:"clOrdLinkID"`
	Account               int       `json:"account"`
	Symbol                string    `json:"symbol"`
	Side                  string    `json:"side"`
	SimpleOrderQty        int       `json:"simpleOrderQty"`
	OrderQty              int       `json:"orderQty"`
	Price                 float64   `json:"price"`
	DisplayQty            int       `json:"displayQty"`
	StopPx                int       `json:"stopPx"`
	PegOffsetValue        int       `json:"pegOffsetValue"`
	PegPriceType          string    `json:"pegPriceType"`
	Currency              string    `json:"currency"`
	SettlCurrency         string    `json:"settlCurrency"`
	OrdType               string    `json:"ordType"`
	TimeInForce           string    `json:"timeInForce"`
	ExecInst              string    `json:"execInst"`
	ContingencyType       string    `json:"contingencyType"`
	ExDestination         string    `json:"exDestination"`
	OrdStatus             string    `json:"ordStatus"`
	Triggered             string    `json:"triggered"`
	WorkingIndicator      bool      `json:"workingIndicator"`
	OrdRejReason          string    `json:"ordRejReason"`
	SimpleLeavesQty       int       `json:"simpleLeavesQty"`
	LeavesQty             int       `json:"leavesQty"`
	SimpleCumQty          int       `json:"simpleCumQty"`
	CumQty                int       `json:"cumQty"`
	AvgPx                 int       `json:"avgPx"`
	MultiLegReportingType string    `json:"multiLegReportingType"`
	Text                  string    `json:"text"`
	TransactTime          time.Time `json:"transactTime"`
	Timestamp             time.Time `json:"timestamp"`
}
