#ifndef _mdspi_h_
#define _mdspi_h_

#include <stdint.h>


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
