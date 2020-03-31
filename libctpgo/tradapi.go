package main

import "C"
import (
	"fmt"
)

//export CTPG_TRADE_OnFrontConnected
func CTPG_TRADE_OnFrontConnected(tradespiId uint32) {
	fmt.Println("CTPG_TRADE_OnFrontConnected: ", tradespiId)
}

//export CTPG_TRADE_OnFrontDisconnected
func CTPG_TRADE_OnFrontDisconnected(tradespiId uint32,
	nReason int) {
	fmt.Println("CTPG_TRADE_OnFrontDisconnected: ", tradespiId)
}

//export CTPG_TRADE_OnHeartBeatWarning
func CTPG_TRADE_OnHeartBeatWarning(tradespiId uint32,
	nTimeLapse int) {
	fmt.Println("CTPG_TRADE_OnHeartBeatWarning: ", tradespiId)
}

//export CTPG_TRADE_OnRspAuthenticate
func CTPG_TRADE_OnRspAuthenticate(tradespiId uint32,
	BrokerID string,
	UserID string,
	UserProductInfo string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspAuthenticate: ", tradespiId)
}

//export CTPG_TRADE_OnRspUserLogin
func CTPG_TRADE_OnRspUserLogin(tradespiId uint32,
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
	fmt.Println("CTPG_TRADE_OnRspUserLogin: ", tradespiId)

	if ErrorID == 0 {

	} else {

	}
}

//export CTPG_TRADE_OnRspUserLogout
func CTPG_TRADE_OnRspUserLogout(tradespiId uint32,
	BrokerID string,
	UserID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspUserLogout: ", tradespiId)
}

//export CTPG_TRADE_OnRspUserPasswordUpdate
func CTPG_TRADE_OnRspUserPasswordUpdate(tradespiId uint32,
	BrokerID string,
	UserID string,
	OldPassword string,
	NewPassword string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspUserPasswordUpdate: ", tradespiId)
}

//export CTPG_TRADE_OnRspTradingAccountPasswordUpdate
func CTPG_TRADE_OnRspTradingAccountPasswordUpdate(tradespiId uint32,
	BrokerID string,
	AccountID string,
	OldPassword string,
	NewPassword string,
	CurrencyID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspTradingAccountPasswordUpdate: ", tradespiId)
}

//export CTPG_TRADE_OnRspOrderInsert
func CTPG_TRADE_OnRspOrderInsert(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	InstrumentID string,
	OrderRef string,
	UserID string,
	OrderPriceType string,
	Direction byte,
	CombOffsetFlag string,
	CombHedgeFlag string,
	LimitPrice float64,
	VolumeTotalOriginal int,
	TimeCondition byte,
	GTDDate string,
	VolumeCondition byte,
	MinVolume int,
	ContingentCondition byte,
	StopPrice float64,
	ForceCloseReason byte,
	IsAutoSuspend int,
	BusinessUnit string,
	RequestID int,
	UserForceClose int,
	IsSwapOrder int,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspOrderInsert: ", tradespiId)
}

//export CTPG_TRADE_OnRspParkedOrderInsert
func CTPG_TRADE_OnRspParkedOrderInsert(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	InstrumentID string,
	OrderRef string,
	UserID string,
	OrderPriceType string,
	Direction byte,
	CombOffsetFlag string,
	CombHedgeFlag string,
	LimitPrice float64,
	VolumeTotalOriginal int,
	TimeCondition byte,
	GTDDate string,
	VolumeCondition byte,
	MinVolume int,
	ContingentCondition byte,
	StopPrice float64,
	ForceCloseReason byte,
	IsAutoSuspend int,
	BusinessUnit string,
	RequestID int,
	UserForceClose int,
	ExchangeID string,
	ParkedOrderID string,
	UserType byte,
	Status byte,
	rspErrorID int,
	rspErrorMsg string,
	IsSwapOrder int,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspParkedOrderInsert: ", tradespiId)
}

//export CTPG_TRADE_OnRspOrderAction
func CTPG_TRADE_OnRspOrderAction(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	OrderActionRef int,
	OrderRef string,
	RequestID int,
	FrontID int,
	SessionID int,
	ExchangeID string,
	OrderSysID string,
	ActionFlag byte,
	LimitPrice float64,
	VolumeChange int,
	UserID string,
	InstrumentID string,
	ParkedOrderActionID string,
	UserType byte,
	Status byte,
	rspErrorID int,
	rspErrorMsg string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspOrderAction: ", tradespiId)
}

//export CTPG_TRADE_OnRspQueryMaxOrderVolume
func CTPG_TRADE_OnRspQueryMaxOrderVolume(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	InstrumentID string,
	Direction byte,
	OffsetFlag byte,
	HedgeFlag byte,
	MaxVolume int,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQueryMaxOrderVolume: ", tradespiId)
}

//export CTPG_TRADE_OnRspSettlementInfoConfirm
func CTPG_TRADE_OnRspSettlementInfoConfirm(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	ConfirmDate string,
	ConfirmTime string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspSettlementInfoConfirm: ", tradespiId)
}

//export CTPG_TRADE_OnRspRemoveParkedOrder
func CTPG_TRADE_OnRspRemoveParkedOrder(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	ParkedOrderID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspRemoveParkedOrder: ", tradespiId)
}

//export CTPG_TRADE_OnRspRemoveParkedOrderAction
func CTPG_TRADE_OnRspRemoveParkedOrderAction(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	ParkedOrderActionID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspRemoveParkedOrderAction: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryOrder
func CTPG_TRADE_OnRspQryOrder(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	InstrumentID string,
	OrderRef string,
	UserID string,
	OrderPriceType string,
	Direction byte,
	CombOffsetFlag string,
	CombHedgeFlag string,
	LimitPrice float64,
	VolumeTotalOriginal int,
	TimeCondition byte,
	GTDDate string,
	VolumeCondition byte,
	MinVolume int,
	ContingentCondition byte,
	StopPrice float64,
	ForceCloseReason byte,
	IsAutoSuspend int,
	BusinessUnit string,
	RequestID int,
	OrderLocalID string,
	ExchangeID string,
	ParticipantID string,
	ClientID string,
	ExchangeInstID string,
	TraderID string,
	InstallID int,
	OrderSubmitStatus byte,
	NotifySequence int,
	TradingDay string,
	SettlementID int,
	OrderSysID string,
	OrderSource byte,
	OrderStatus byte,
	OrderType byte,
	VolumeTraded int,
	VolumeTotal int,
	InsertDate string,
	InsertTime string,
	ActiveTime string,
	SuspendTime string,
	UpdateTime string,
	CancelTime string,
	ActiveTraderID string,
	ClearingPartID string,
	SequenceNo int,
	FrontID int,
	SessionID int,
	UserProductInfo string,
	StatusMsg string,
	UserForceClose int,
	ActiveUserID string,
	BrokerOrderSeq int,
	RelativeOrderSysID string,
	ZCETotalTradedVolume int,
	IsSwapOrder int,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryOrder: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryTrade
func CTPG_TRADE_OnRspQryTrade(tradespiId uint32,
	BrokerID string,
	InvestorID string,
	InstrumentID string,
	OrderRef string,
	UserID string,
	ExchangeID string,
	TradeID string,
	Direction byte,
	OrderSysID string,
	ParticipantID string,
	ClientID string,
	TradingRole byte,
	ExchangeInstID string,
	OffsetFlag byte,
	HedgeFlag byte,
	Price float64,
	Volume int,
	TradeDate string,
	TradeTime string,
	TradeType byte,
	PriceSource byte,
	TraderID string,
	OrderLocalID string,
	ClearingPartID string,
	BusinessUnit string,
	SequenceNo int,
	TradingDay string,
	SettlementID int,
	BrokerOrderSeq int,
	TradeSource byte,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryTrade: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryInvestorPosition
func CTPG_TRADE_OnRspQryInvestorPosition(tradespiId uint32,
	InstrumentID string,
	BrokerID string,
	InvestorID string,
	PosiDirection byte,
	HedgeFlag byte,
	PositionDate byte,
	YdPosition int,
	Position int,
	LongFrozen int,
	ShortFrozen int,
	LongFrozenAmount float64,
	ShortFrozenAmount float64,
	OpenVolume int,
	CloseVolume int,
	OpenAmount float64,
	CloseAmount float64,
	PositionCost float64,
	PreMargin float64,
	UseMargin float64,
	FrozenMargin float64,
	FrozenCash float64,
	FrozenCommission float64,
	CashIn float64,
	Commission float64,
	CloseProfit float64,
	PositionProfit float64,
	PreSettlementPrice float64,
	SettlementPrice float64,
	TradingDay string,
	SettlementID int,
	OpenCost float64,
	ExchangeMargin float64,
	CombPosition int,
	CombLongFrozen int,
	CombShortFrozen int,
	CloseProfitByDate float64,
	CloseProfitByTrade float64,
	TodayPosition int,
	MarginRateByMoney float64,
	MarginRateByVolume float64,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryInvestorPosition: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryTradingAccount
func CTPG_TRADE_OnRspQryTradingAccount(tradespiId uint32,
	BrokerID string,
	AccountID string,
	PreMortgage float64,
	PreCredit float64,
	PreDeposit float64,
	PreBalance float64,
	PreMargin float64,
	InterestBase float64,
	Interest float64,
	Deposit float64,
	Withdraw float64,
	FrozenMargin float64,
	FrozenCash float64,
	FrozenCommission float64,
	CurrMargin float64,
	CashIn float64,
	Commission float64,
	CloseProfit float64,
	PositionProfit float64,
	Balance float64,
	Available float64,
	WithdrawQuota float64,
	Reserve float64,
	TradingDay string,
	SettlementID int,
	Credit float64,
	Mortgage float64,
	ExchangeMargin float64,
	DeliveryMargin float64,
	ExchangeDeliveryMargin float64,
	ReserveBalance float64,
	CurrencyID string,
	PreFundMortgageIn float64,
	PreFundMortgageOut float64,
	FundMortgageIn float64,
	FundMortgageOut float64,
	FundMortgageAvailable float64,
	MortgageableFund float64,
	SpecProductMargin float64,
	SpecProductFrozenMargin float64,
	SpecProductCommission float64,
	SpecProductFrozenCommission float64,
	SpecProductPositionProfit float64,
	SpecProductCloseProfit float64,
	SpecProductPositionProfitByAlg float64,
	SpecProductExchangeMargin float64,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryTradingAccount: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryInvestor
func CTPG_TRADE_OnRspQryInvestor(tradespiId uint32,
	InvestorID string,
	BrokerID string,
	InvestorGroupID string,
	InvestorName string,
	IdentifiedCardType byte,
	IdentifiedCardNo string,
	IsActive int,
	Telephone string,
	Address string,
	OpenDate string,
	Mobile string,
	CommModelID string,
	MarginModelID string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryInvestor: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryTradingCode
func CTPG_TRADE_OnRspQryTradingCode(tradespiId uint32,
	InvestorID string,
	BrokerID string,
	ExchangeID string,
	ClientID string,
	IsActive int,
	ClientIDType byte,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryTradingCode: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryInstrumentMarginRate
func CTPG_TRADE_OnRspQryInstrumentMarginRate(tradespiId uint32,
	InstrumentID string,
	InvestorRange byte,
	BrokerID string,
	InvestorID string,
	HedgeFlag byte,
	LongMarginRatioByMoney float64,
	LongMarginRatioByVolume float64,
	ShortMarginRatioByMoney float64,
	ShortMarginRatioByVolume float64,
	IsRelative int,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryInstrumentMarginRate: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryInstrumentCommissionRate
func CTPG_TRADE_OnRspQryInstrumentCommissionRate(tradespiId uint32,
	InstrumentID string,
	InvestorRange byte,
	BrokerID string,
	InvestorID string,
	OpenRatioByMoney float64,
	OpenRatioByVolume float64,
	CloseRatioByMoney float64,
	CloseRatioByVolume float64,
	CloseTodayRatioByMoney float64,
	CloseTodayRatioByVolume float64,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryInstrumentCommissionRate: ", tradespiId)

}

//export CTPG_TRADE_OnRspQryExchange
func CTPG_TRADE_OnRspQryExchange(tradespiId uint32,
	ExchangeID string,
	ExchangeName string,
	ExchangeProperty byte,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryExchange: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryProduct
func CTPG_TRADE_OnRspQryProduct(tradespiId uint32,
	ProductID string,
	ProductName string,
	ExchangeID string,
	ProductClass byte,
	VolumeMultiple int,
	PriceTick float64,
	MaxMarketOrderVolume int,
	MinMarketOrderVolume int,
	MaxLimitOrderVolume int,
	MinLimitOrderVolume int,
	PositionType byte,
	PositionDateType byte,
	CloseDealType byte,
	TradeCurrencyID string,
	MortgageFundUseRange byte,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryProduct: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryInstrument
func CTPG_TRADE_OnRspQryInstrument(tradespiId uint32,
	InstrumentID string,
	ExchangeID string,
	InstrumentName string,
	ExchangeInstID string,
	ProductID string,
	ProductClass byte,
	DeliveryYear int,
	DeliveryMonth int,
	MaxMarketOrderVolume int,
	MinMarketOrderVolume int,
	MaxLimitOrderVolume int,
	MinLimitOrderVolume int,
	VolumeMultiple int,
	PriceTick float64,
	CreateDate string,
	OpenDate string,
	ExpireDate string,
	StartDelivDate string,
	EndDelivDate string,
	InstLifePhase byte,
	IsTrading int,
	PositionType byte,
	PositionDateType byte,
	LongMarginRatio float64,
	ShortMarginRatio float64,
	MaxMarginSideAlgorithm byte,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryInstrument: ", tradespiId)
}

//export CTPG_TRADE_OnRspQryDepthMarketData
func CTPG_TRADE_OnRspQryDepthMarketData(tradespiId uint32,
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
	UpdateTime string,
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
	ActionDay string,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspQryDepthMarketData: ", tradespiId)
}

//export CTPG_TRADE_OnRspError
func CTPG_TRADE_OnRspError(tradespiId uint32,
	ErrorID int,
	ErrorMsg string,
	nRequestID int,
	bIsLast bool) {
	fmt.Println("CTPG_TRADE_OnRspError: ", tradespiId)
}

/*
	请求查询投资者结算结果响应
	virtual void OnRspQrySettlementInfo(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询转帐银行响应
	virtual void OnRspQryTransferBank(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询投资者持仓明细响应
	virtual void OnRspQryInvestorPositionDetail(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询客户通知响应
	virtual void OnRspQryNotice(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询结算信息确认响应
	virtual void OnRspQrySettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询投资者持仓明细响应
	virtual void OnRspQryInvestorPositionCombineDetail(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	查询保证金监管系统经纪公司资金账户密钥响应
	virtual void OnRspQryCFMMCTradingAccountKey(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询仓单折抵信息响应
	virtual void OnRspQryEWarrantOffset(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询投资者品种/跨品种保证金响应
	virtual void OnRspQryInvestorProductGroupMargin(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询交易所保证金率响应
	virtual void OnRspQryExchangeMarginRate(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询交易所调整保证金率响应
	virtual void OnRspQryExchangeMarginRateAdjust(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询汇率响应
	virtual void OnRspQryExchangeRate(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询二级代理操作员银期权限响应
	virtual void OnRspQrySecAgentACIDMap(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询转帐流水响应
	virtual void OnRspQryTransferSerial(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询银期签约关系响应
	virtual void OnRspQryAccountregister(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	错误应答
	virtual void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	报单通知
	virtual void OnRtnOrder(CThostFtdcOrderField *pOrder) {};

	成交通知
	virtual void OnRtnTrade(CThostFtdcTradeField *pTrade) {};

	报单录入错误回报
	virtual void OnErrRtnOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo) {};

	报单操作错误回报
	virtual void OnErrRtnOrderAction(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo) {};

	合约交易状态通知
	virtual void OnRtnInstrumentStatus(CThostFtdcInstrumentStatusField *pInstrumentStatus) {};

	交易通知
	virtual void OnRtnTradingNotice(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo) {};

	提示条件单校验错误
	virtual void OnRtnErrorConditionalOrder(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder) {};

	请求查询签约银行响应
	virtual void OnRspQryContractBank(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询预埋单响应
	virtual void OnRspQryParkedOrder(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询预埋撤单响应
	virtual void OnRspQryParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询交易通知响应
	virtual void OnRspQryTradingNotice(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询经纪公司交易参数响应
	virtual void OnRspQryBrokerTradingParams(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	请求查询经纪公司交易算法响应
	virtual void OnRspQryBrokerTradingAlgos(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	银行发起银行资金转期货通知
	virtual void OnRtnFromBankToFutureByBank(CThostFtdcRspTransferField *pRspTransfer) {};

	银行发起期货资金转银行通知
	virtual void OnRtnFromFutureToBankByBank(CThostFtdcRspTransferField *pRspTransfer) {};

	银行发起冲正银行转期货通知
	virtual void OnRtnRepealFromBankToFutureByBank(CThostFtdcRspRepealField *pRspRepeal) {};

	银行发起冲正期货转银行通知
	virtual void OnRtnRepealFromFutureToBankByBank(CThostFtdcRspRepealField *pRspRepeal) {};

	期货发起银行资金转期货通知
	virtual void OnRtnFromBankToFutureByFuture(CThostFtdcRspTransferField *pRspTransfer) {};

	期货发起期货资金转银行通知
	virtual void OnRtnFromFutureToBankByFuture(CThostFtdcRspTransferField *pRspTransfer) {};

	系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	virtual void OnRtnRepealFromBankToFutureByFutureManual(CThostFtdcRspRepealField *pRspRepeal) {};

	系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	virtual void OnRtnRepealFromFutureToBankByFutureManual(CThostFtdcRspRepealField *pRspRepeal) {};

	期货发起查询银行余额通知
	virtual void OnRtnQueryBankBalanceByFuture(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount) {};

	期货发起银行资金转期货错误回报
	virtual void OnErrRtnBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo) {};

	期货发起期货资金转银行错误回报
	virtual void OnErrRtnFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo) {};

	系统运行时期货端手工发起冲正银行转期货错误回报
	virtual void OnErrRtnRepealBankToFutureByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo) {};

	系统运行时期货端手工发起冲正期货转银行错误回报
	virtual void OnErrRtnRepealFutureToBankByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo) {};

	期货发起查询银行余额错误回报
	virtual void OnErrRtnQueryBankBalanceByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo) {};

	期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	virtual void OnRtnRepealFromBankToFutureByFuture(CThostFtdcRspRepealField *pRspRepeal) {};

	期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	virtual void OnRtnRepealFromFutureToBankByFuture(CThostFtdcRspRepealField *pRspRepeal) {};

	期货发起银行资金转期货应答
	virtual void OnRspFromBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	期货发起期货资金转银行应答
	virtual void OnRspFromFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	期货发起查询银行余额应答
	virtual void OnRspQueryBankAccountMoneyByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};

	银行发起银期开户通知
	virtual void OnRtnOpenAccountByBank(CThostFtdcOpenAccountField *pOpenAccount) {};

	银行发起银期销户通知
	virtual void OnRtnCancelAccountByBank(CThostFtdcCancelAccountField *pCancelAccount) {};

	银行发起变更银行账号通知
	virtual void OnRtnChangeAccountByBank(CThostFtdcChangeAccountField *pChangeAccount) {};
*/
