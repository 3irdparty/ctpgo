#ifndef _mdspi_h_
#define _mdspi_h_

#include <stdint.h>


#include "ThostFtdcMdApi.h"

class CTPGO_MdSpi :public CThostFtdcMdSpi {
public:
	CTPGO_MdSpi(CThostFtdcMdApi *pUserApi, int idx);
	~CTPGO_MdSpi();


	void OnFrontConnected();

	///当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	///@param nReason 错误原因
	///        0x1001 网络读失败
	///        0x1002 网络写失败
	///        0x2001 接收心跳超时
	///        0x2002 发送心跳失败
	///        0x2003 收到错误报文
	virtual void OnFrontDisconnected(int nReason);

	///心跳超时警告。当长时间未收到报文时，该方法被调用。
	///@param nTimeLapse 距离上次接收报文的时间
	void OnHeartBeatWarning(int nTimeLapse);


	///登录请求响应
	void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	

	
	/* 行情应答  */
	void  OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData);

	void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	///登出请求响应
	void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	///错误应答
	/*void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
		printf("CtpMdSpi::OnRspError()\n");
	}*/

	///订阅行情应答
	void OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	///取消订阅行情应答
	void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	///订阅询价应答
	void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	///取消订阅询价应答
	void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);

	//深度行情通知
	void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData);

	///询价通知
	void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp);

private:
	CThostFtdcMdApi *m_pUserApi;
	uint32_t m_id;
};

#ifdef __cplusplus
extern "C" {
#endif

///发给做市商的询价请求
struct CTPGO_CThostFtdcForQuoteRspField
{
	///交易日
	char	TradingDay[9];
	///合约代码
	char	InstrumentID[31];
	///询价编号
	char	ForQuoteSysID[21];
	///询价时间
	char	ForQuoteTime[9];
	///业务日期
	char	ActionDay[9];
	///交易所代码
	char	ExchangeID[9];
};

#ifdef __cplusplus
}
#endif

#endif // _mdspi_h_
