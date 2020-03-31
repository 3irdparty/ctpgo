package main

//##include "mdspi.h"
//#cgo LDFLAGS: thostmduserapi.so
//#cgo LDFLAGS: thosttraderapi.so
//#cgo LDFLAGS: libctpgo.so
import "C"
import (
	"fmt"
)

func OnFrontConnected(mdspiId uint32) {
	//printf("CtpMdSpi::OnFrontConnected()\n");

	//for (int i = 0; i < 10; i++) {
	//TThostFtdcBrokerIDType brokerID = "9999";
	// 交易用户代码
	/*TThostFtdcUserIDType userID;
	sprintf_s(userID, "%06d", 35877 + m_idx);;
	TThostFtdcUserIDType passwd = "888888";
	printf("--- userID: %s\n", userID);
	CThostFtdcReqUserLoginField req;
	memset(&req, 0, sizeof(req));


	strcpy_s(req.BrokerID, brokerID);
	strcpy_s(req.UserID, userID);
	strcpy_s(req.Password, passwd);
	int ret = m_pUserApi->ReqUserLogin(&req, 0);
	switch (ret) {
	case 0:
		break;
	case 1:
		break;
	case 2:
		break;
	case 3:
		break;
	default:
		break;
	}*/
	//Sleep(1000);
	//printf("login ret: %d\n", ret);
	//}
}

///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
///@param nReason 错误原因
///        0x1001 网络读失败
///        0x1002 网络写失败
///        0x2001 接收心跳超时
///        0x2002 发送心跳失败
///        0x2003 收到错误报文
func OnFrontDisconnected(mdspiId uint32, int nReason) {
	//printf("CtpMdSpi::OnFrontDisconnected()\n");
}

///心跳超时警告。当长时间未收到报文时，该方法被调用。
///@param nTimeLapse 距离上次接收报文的时间
func OnHeartBeatWarning(mdspiId uint32, int nTimeLapse) {
	//printf("CtpMdSpi::OnHeartBeatWarning()\n");
}

///登录请求响应
func OnRspUserLogin(mdspiId uint32, userLoginRsp *C.CTPGO_CThostFtdcRspUserLoginField, rspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, isLast bool) {
	//printf("CtpMdSpi::OnRspUserLogin() Broker: %s, UserID: %c\n", pRspUserLogin->BrokerID, pRspUserLogin->UserID[0]);

	if rspInfo.ErrorID == 0 {
		/*printf("Login Success!\n");

		char * Instrumnet[] = { "IF0809", "IF0812" };
		m_pUserApi->SubscribeMarketData(Instrumnet, 2);*/
		/* 或退订行情 */
		//m_pUserApi->UnSubscribeMarketData(Instrumnet, 2);
	} else {
		//printf("Login Failed!\n");
	}
}

/* 行情应答  */
func OnRtnDepthMarketData(mdspiId uint32, pDepthMarketData *C.CTPGO_CThostFtdcDepthMarketDataField) {
	// 输出报单录入结果
	//printf("pDepthMarketData->LastPrice= %f\n", pDepthMarketData->LastPrice);

	//收到深度行情
	//SetEvent(g_hEvent);

}

func OnRspError(mdspiId uint32, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	/*printf("OnRspError:\n");
	printf("ErrorCode=[%d],  ErrorMsg=[%s]\n", pRspInfo->ErrorID, pRspInfo->ErrorMsg);
	printf("RequestID=[%d], Chain=[%d]\n", nRequestID, bIsLast);*/
}

///登出请求响应
func OnRspUserLogout(mdspiId uint32, pUserLogout *C.CTPGO_CThostFtdcUserLogoutField, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//printf("CtpMdSpi::OnRspUserLogout()\n");
}

///错误应答
func OnRspError(pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//printf("CtpMdSpi::OnRspError()\n");
}

///订阅行情应答
func OnRspSubMarketData(mdspiId uint32, pSpecificInstrument *C.CTPGO_CThostFtdcSpecificInstrumentField, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//printf("CtpMdSpi::OnRspSubMarketData(), nRequestID: %d, bIsLast: %d, pSpecificInstrument->InstrumentID: %s\n", nRequestID, int(bIsLast), pSpecificInstrument->InstrumentID);
	//printf("ErrorCode=[%d],  ErrorMsg=[%s]\n", pRspInfo->ErrorID, pRspInfo->ErrorMsg);
	//收到深度行情
	//SetEvent(g_hEvent);
}

///取消订阅行情应答
func OnRspUnSubMarketData(mdspiId uint32, pSpecificInstrument *C.CTPGO_CThostFtdcSpecificInstrumentField, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//printf("CtpMdSpi::OnRspUnSubMarketData()\n");
}

///订阅询价应答
func OnRspSubForQuoteRsp(mdspiId uint32, pSpecificInstrument *C.CTPGO_CThostFtdcSpecificInstrumentField, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	//printf("CtpMdSpi::OnRspSubForQuoteRsp()\n");
}

///取消订阅询价应答
func OnRspUnSubForQuoteRsp(mdspiId uint32, pSpecificInstrument *C.CTPGO_CThostFtdcSpecificInstrumentField, pRspInfo *C.CTPGO_CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
}

//深度行情通知
func OnRtnDepthMarketData(mdspiId uint32, pDepthMarketData *C.CTPGO_CThostFtdcDepthMarketDataField) {
	//printf("CtpMdSpi::OnRtnDepthMarketData()\n");
}

///询价通知
func OnRtnForQuoteRsp(mdspiId uint32, quoteRsp *C.CTPGO_CThostFtdcForQuoteRspField) {
	//printf("CtpMdSpi::OnRtnForQuoteRsp()\n");
}
