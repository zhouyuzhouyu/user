package code

var UserMsg = map[int]string{
	SUCCESS:     "success",
	BadRequest:  "Fail",
	TimeExpired: "检测到您的手机时间和当前时间不一致，请到设置里调整为“自动设置”",
	NETWORKERR:  "网络繁忙, 请稍后重试",
	AUTHERR:     "权限认证失败，请稍后重试",

	ExchangeTradeInfoEmpty: "订单信息为空或已退订",

	PARAMSERR: "不合法的请求",

	NotMobileEmpty: "手机号不能为空",

	NotMobileErr: "手机号错误,请检查后重新输入",

	PayError: "支付调用失败，请稍后重试或联系客服",

	PointInfoGetErr: "积分信息获取失败",

	UnknownErr: "未知错误",

	PointGainLimitErr: "获取积分已达上限",

	ExchangeRequestErr: "请求商城失败，请稍后再试",

	ExchangeUserInfoErr: "用户信息查询失败",

	BankUnAuthVerifyCodeErr: "验证码错误",

	BankUnAuthTimesLimitErr: "每24小时仅可解绑三次银行卡哦",

	BankCardNotUsedErr: "暂不支持此银行卡，请查看支持银行",

	BankCardNotConfErr: "暂不支持此银行卡，请查看支持银行",

	UserTradeLimit: "订购失败，如有疑问，请联系客服",

	UserLoginLimit: "登录失败，如有疑问，请联系客服",

	SysAgeLess18Err: "18周岁以下暂无法实名",

	UserLocked: "您的账号已被锁定。可以用忘记密码找回，同步解锁，或联系客服解锁",

	BankIncorrectCode: "银行卡号与所属银行不匹配，请核对后提交",

	TaskReadErr: "设置阅读任务失败",

	TaskDoneErr: "任务领取失败",

	WarningSetErr: "设置失败，请稍后重试",

	WarningArriveUpperLimitErr: "单个商品设置已达上限，待触发后才可继续添加哦",

	WarningSetRepeatErr: "该价格已被您设置过啦，正等待触发~",

	WarningSetMeetTheConditionsErr: "最新价已满足预警条件，请重新设置",

	BankFourElementsErr: "四要素验证失败，请检查后重试",

	BankRepeatErr: "银行卡已添加，请换张卡添加",

	WithdrawRiskFail:    "用户账户有风险，请联系在线客服",
	WithdrawBalanceFail: "可用余额不足",

	PositionOpenButton: "下单异常，请联系客服反馈",
	RechargeNeedAuth:   "下单异常，请实名认证", //实名认证
	PwdLimit:           "密码输入错误次数已经到5次了",
	PwdWrong:           "密码不能是3个或3个以上的连续、重复数字 (如123456、666888、123123)",
}
