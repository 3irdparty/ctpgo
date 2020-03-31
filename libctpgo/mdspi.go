package main

import "C"
import (
	"fmt"
)

//export CTPG_MD_OnFrontConnected
func CTPG_MD_OnFrontConnected(mdspiId uint32) {
	fmt.Println("CTPG_MD_OnFrontConnected: ", mdspiId)
}

//export CTPG_MD_OnFrontDisconnected
func CTPG_MD_OnFrontDisconnected(mdspiId uint32,
	nReason int) {
	fmt.Println("CTPG_MD_OnFrontDisconnected: ", mdspiId)
}

//export CTPG_MD_OnHeartBeatWarning
func CTPG_MD_OnHeartBeatWarning(mdspiId uint32,
	nTimeLapse int) {
	fmt.Println("CTPG_MD_OnHeartBeatWarning: ", mdspiId)
}

//export CTPG_MD_OnRspUserLogin
func CTPG_MD_OnRspUserLogin(mdspiId uint32,
	TradingDay string,
	LoginTime string,
	BrokerID string,
	UserID string,
	SystemName string,
	FrontID int,
	SessionID int,
	MaxOrderRef string,
	SHFETime string,
	DCETime string,
	CZCETime string,
	FFEXTime string,
	INETime string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	isLast bool) {
	fmt.Println("CTPG_MD_OnRspUserLogin: ", mdspiId)

	if ErrorID == 0 {

	} else {

	}
}

//export CTPG_MD_OnRspError
func CTPG_MD_OnRspError(mdspiId uint32,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspError: ", mdspiId)
}

//export CTPG_MD_OnRspUserLogout
func CTPG_MD_OnRspUserLogout(mdspiId uint32,
	BrokerID string,
	UserID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspUserLogout: ", mdspiId)
}

//export CTPG_MD_OnRspSubMarketData
func CTPG_MD_OnRspSubMarketData(mdspiId uint32,
	InstrumentID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspSubMarketData: ", mdspiId)
}

//export CTPG_MD_OnRspUnSubMarketData
func CTPG_MD_OnRspUnSubMarketData(mdspiId uint32,
	InstrumentID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspUnSubMarketData: ", mdspiId)
}

//export CTPG_MD_OnRspSubForQuoteRsp
func CTPG_MD_OnRspSubForQuoteRsp(mdspiId uint32,
	InstrumentID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspSubForQuoteRsp: ", mdspiId)
}

//export CTPG_MD_OnRspUnSubForQuoteRsp
func CTPG_MD_OnRspUnSubForQuoteRsp(mdspiId uint32,
	InstrumentID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_MD_OnRspUnSubForQuoteRsp: ", mdspiId)
}

//export CTPG_MD_OnRtnDepthMarketData
func CTPG_MD_OnRtnDepthMarketData(mdspiId uint32,
	TradingDay string,
	InstrumentID string,
	ExchangeID string,
	ExchangeInstID string,
	LastPrice float64,
	PreSettlementPrice float64,
	PreClosePrice float64,
	PreOpenInterest float64,
	OpenPrice float64,
	HighestPrice float64,
	LowestPrice float64,
	Volume int,
	Turnover float64,
	OpenInterest float64,
	ClosePrice float64,
	SettlementPrice float64,
	UpperLimitPrice float64,
	LowerLimitPrice float64,
	PreDelta float64,
	CurrDelta float64,
	UpdateTime byte,
	UpdateMillisec int,
	BidPrice1 float64,
	BidVolume1 int,
	AskPrice1 float64,
	AskVolume1 int,
	BidPrice2 float64,
	BidVolume2 int,
	AskPrice2 float64,
	AskVolume2 int,
	BidPrice3 float64,
	BidVolume3 int,
	AskPrice3 float64,
	AskVolume3 int,
	BidPrice4 float64,
	BidVolume4 int,
	AskPrice4 float64,
	AskVolume4 int,
	BidPrice5 float64,
	BidVolume5 int,
	AskPrice5 float64,
	AskVolume5 int,
	AveragePrice float64,
	ActionDay string) {
	fmt.Println("CTPG_MD_OnRtnDepthMarketData: ", mdspiId)
}

//export CTPG_MD_OnRtnForQuoteRsp
func CTPG_MD_OnRtnForQuoteRsp(mdspiId uint32,
	TradingDay string,
	InstrumentID string,
	ForQuoteSysID string,
	ForQuoteTime string,
	ActionDay string,
	ExchangeID string) {
	fmt.Println("CTPG_MD_OnRtnForQuoteRsp: ", mdspiId)
}
