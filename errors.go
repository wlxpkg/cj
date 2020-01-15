package zwyd

import "errors"

type Errors struct {
	Code    int
	Message string
}

var Errs = map[string]Errors{
	"ERR_UNKNOW_ERROR":     Errors{0, "未知系统错误"},
	"ERR_NOERROR":          Errors{1, ""},
	"SUCCESS":              Errors{1, ""},
	"ERR_TOO_MANY_REQUEST": Errors{429, "操作过于频繁, 请稍候再试"},

	"ERR_PARAM":         Errors{1000, "参数错误"},
	"ERR_UNLOGIN":       Errors{1001, "请先注册登录"},
	"ERR_INVALID_TOKEN": Errors{1002, "无效的token"},
	"ERR_UNAUTHORIZED":  Errors{1004, "您没有权限访问该数据"},
	"ERR_DATA_DECODE":   Errors{1005, "数据解析失败"},
	"ERR_HTTP_BASEURL":  Errors{1006, "请设置 baseurl"},
	"ERR_TCP_TIMEOUT":   Errors{1504, "TCP接口响应超时"},
	"ERR_HTTP_TIMEOUT":  Errors{1505, "HTTP接口响应超时"},

	"ERR_MYSQL":              Errors{2000, "MySQL错误"},
	"ERR_MYSQL_INSTALL_FAIL": Errors{2001, "数据插入失败"},
	"ERR_MYSQL_DELETE_FAIL":  Errors{2002, "数据删除失败"},
	"ERR_MYSQLPOOL_FAIL":     Errors{2004, "mysql连接池丢失"},
	"ERR_REDIS":              Errors{2100, "Reids错误"},
	"ERR_REDISPOOL_FAIL":     Errors{2104, "redis连接池丢失"},
	"ERR_IDGEN_FAIL":         Errors{2404, "id生成失败"},

	"ERR_ALBUM_GROUP_NAME_EMPTY":        Errors{15000, "请输入团队名称"},
	"ERR_ALBUM_GROUP_CHECK_PASSCODE":    Errors{15001, "未找到团队，请核实团队暗号"},
	"ERR_ALBUM_GROUP_HAS_JOIN":          Errors{15002, "您已加入该团队"},
	"ERR_ALBUM_GROUP_CREATE_OVER_LIMIT": Errors{15003, "您创建的团队已达上限"},
	"ERR_ALBUM_GROUP_JOIN_OVER_LIMIT":   Errors{15004, "您加入的团队已达上限"},
	"ERR_ALBUM_GROUP_PASSCODE_EXIST":    Errors{15005, "该暗号已存在"},
	"ERR_ALBUM_GROUP_NOT_EXIST":         Errors{15006, "未找到该团队"},
	"ERR_ALBUM_GROUP_NOT_JOIN":          Errors{15007, "您未加入该团队"},

	"ERR_MATERIAL_CANNOT_PUSH":      Errors{16000, "非发布者无法发布"},
	"ERR_MATERIAL_NOT_EXIST":        Errors{16001, "素材不存在"},
	"ERR_MATERIAL_CANNOT_DELETE":    Errors{16002, "无权删除"},
	"ERR_MATERIAL_CANNOT_TOP":       Errors{16003, "无权置顶"},
	"ERR_MATERIAL_PUBLISH_FAILED":   Errors{16004, "发布失败"},
	"ERR_MATERIAL_PETAL_NOT_ENOUGH": Errors{16005, "花瓣不足，发布失败"},

	"ERR_PAY_FUBEI_FAIL": Errors{20000, "支付请求失败"},

	"ERR_VIDEO_NOT_EXIST":        Errors{25000, "视频不存在"},
	"ERR_VIDEO_EXCEEDED_MAXIMUM": Errors{25001, "翻页总条数超过最大限制"},
	"ERR_GET_FAIL":               Errors{25002, "获取视频异常"},

	"ERR_ORDER_NOT_EXIST":         Errors{21000, "订单不存在"},
	"ERR_ORDER_ALREADY_PAY":       Errors{21001, "订单已经支付"},
	"ERR_ORDER_ALREADY_REFUNDING": Errors{21002, "订单正在退款中"},
	"ERR_ORDER_ALREADY_REFUND":    Errors{21003, "订单已经退款"},
	"ERR_ORDER_ALREADY_CLOSE":     Errors{21004, "订单已经关闭"},
	"ERR_ORDER_TIME_OUT":          Errors{21005, "订单已经超时，请重新下单"},
	"ERR_ORDER_APPLY_AGENT_PRICE": Errors{21006, "订单商品未达到申请代理的价格的数量"},
	"ERR_ORDER_AGENT_LEAST_PRICE": Errors{21007, "订单不能低于起购价"},

	"ERR_LINK_APPLY_AGENT_INVALID":       Errors{21101, "申请代理的链接失败"},
	"ERR_LINK_APPLY_AGENT_ALREADY_AGENT": Errors{21102, "申请代理的用户已经是代理用户"},
	"ERR_LINK_GROUP_BOOKING_INVALID":     Errors{21103, "拼团已结束"},
	"ERR_LINK_GROUP_BUYING_INVALID":      Errors{21104, "团购链接已失效"},
	"GROUP_BOOKING_COMPLETE_FAILED":      Errors{21105, "数量不达标，无法提前完成拼团"},
	"ERR_GROUP_BOOKING_STATUS":           Errors{21106, "拼团状态必须为拼团中"},
	"ERR_PRE_PRICE_NO_EXIST":             Errors{21107, "预付款价格获取失败"},
	"ERR_PRE_PRICE_BOOKING_PRICE_ERROR":  Errors{21108, "拼团价格需高于预付款价格"},
	"ERR_BOOKING_PRICE_PRE_PRICE_ERROR":  Errors{21109, "预付款需低于拼团价格价格"},
	"ERR_LINK_RETAIL_INVALID":            Errors{21110, "客服链接失效"},
	"ERR_GROUP_NAME_EXIST":               Errors{21111, "团队名称已存在"},
	"ERR_GROUP_BOOKING_TAIL_ALREADY_PAY": Errors{21112, "您已经支付过尾款了"},
	"ERR_GROUP_LINK_EXISTED":             Errors{21113, "当前团队已经有团购进行中, 请勿重复创建"},
	"ERR_GROUP_BUY_EXCEED_NUM":           Errors{21114, "当前团队已超过购买数量"},
	"ERR_LINK_NO_BY_YOUSELF":             Errors{21115, "不能通过自己的申请链接成为代理"},

	"ERR_WALLET_PAY_FAIL":        Errors{22000, "钱包支付失败"},
	"ERR_WALLET_PAY_EXP":         Errors{22001, "钱包数额不正确"},
	"ERR_GET_PRICE":              Errors{22003, "获取商品价格失败"},
	"ERR_PAY_TYPE":               Errors{22002, "支付方式不正确"},
	"ERR_PAY_NOT_SUPPORT_WALLET": Errors{22002, "该笔订单不支持钱包支付"},

	"ERR_GOODS_EXIST":       Errors{23000, "商品已存在"},
	"ERR_GOODS_NOT_EXIST":   Errors{23001, "商品不存在"},
	"ERR_GOODS_OFF_SHELVES": Errors{23002, "商品已下架"},
	"ERR_GOODS_LOCK":        Errors{23003, "商品锁定失败"},
	"ERR_TIME_EXPIRE":       Errors{23004, "时间设置不正确"},
}

func Excp(errString string) error {
	return errors.New(errString)
}
