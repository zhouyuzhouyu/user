package code

// public
const (
	SUCCESS            = 200 //请求成功
	BadRequest         = 400 // req param mileage
	TimeExpired        = 401
	PositionOpenButton = 2000000 //防沉迷下单限制挂屏
	RechargeNeedAuth   = 2000001 //实名认证
	PwdLimit           = 4000000
	PwdWrong           = 4000001
)

// api 200xxx
const (
	PARAMSERR  = 200400 //参数错误
	NETWORKERR = 200401 //网络错误
	AUTHERR    = 200402 //权限认证失败
)
const (
	SysAgeLess18Err = 200310 //未满18岁

)
const (
	TaskDoneErr = 201400 //任务领取失败
	TaskReadErr = 201401 //阅读失败
)

// 预警
const (
	WarningSetErr                  = 201520 //设置失败
	WarningArriveUpperLimitErr     = 201521 //到达上限
	WarningSetRepeatErr            = 201522 //价格重复
	WarningSetMeetTheConditionsErr = 201523 //满足条件
)

// 提现
const (
	WithdrawRiskFail    = 201530 //提现失败，用户账户风险(请联系在线客服)
	WithdrawBalanceFail = 201531 //可用余额不足
)

// 银行卡
const (
	BankCardNotUsedErr      = 201540 //暂不支持此银行卡，请查看支持银行
	BankCardNotConfErr      = 201541 //暂不支持银行卡，未找到配置信息
	BankFourElementsErr     = 201542 //四要素验证失败
	BankRepeatErr           = 201543 //重复添加银行卡
	BankIncorrectCode       = 201544 //银行卡与输入的不匹配
	BankUnAuthVerifyCodeErr = 201545 //解绑银行卡验证码错误
	BankUnAuthTimesLimitErr = 201546 //每24小时仅可解绑三次银行卡
)

// Not
const (
	NotMobileEmpty = 201600 //手机不能为空
	NotMobileErr   = 201601 //手机号错误
)

// 用户状态 2005xx
const (
	UserLocked     = 200500 //用户已经被锁定
	UserLoginLimit = 200501 //限制登录
	UserTradeLimit = 200502 //限制交易

)

//grpc code
//point 2010xx
//pay 2011xx

// point 2010
const (
	PointInfoGetErr   = 201000 //获取用户已有积分失败
	PointGainLimitErr = 201001 //用户获取积分已达上限

)

// pay 2011
const (
	PayError = 201101 //支付调用失败

)

// Exchange 2030xx
const (
	ExchangeRequestErr     = 203000 //商城请求失败
	ExchangeNotRespErr     = 203001 //交易所无返回数据
	ExchangeUserInfoErr    = 203002 //用户信息查询失败
	ExchangeTradeInfoEmpty = 203003 //订单信息为空或已平仓
	UnknownErr             = 203004 //未知错误，用户初始化grpc错误内容
)

//grpc end

func Text(code int) string {
	return UserMsg[code]
}
