package xiaomipush

const (
	ProductionHost = "https://api.xmpush.xiaomi.com"
	SandboxHost    = "https://sandbox.xmpush.xiaomi.com" // iOS supported only
)

const (
	RegURL           = "/v3/message/regid"          // 向某个regid或一组regid列表推送某条消息
	MessageAllURL    = "/v3/message/all"            // 向所有设备推送某条消息
	MultiMessagesURL = "/v2/multi_messages/regids"  // 针对不同的regid推送不同的消息
	StatsURL         = "/v1/stats/message/counters" // 统计push
	MessageStatusURL = "/v1/trace/message/status"   // 获取指定ID的消息状态
)

// for future targeted push
var (
	BrandsMap = map[string]string{
		"品牌":    "MODEL",
		"小米":    "xiaomi",
		"三星":    "samsung",
		"华为":    "huawei",
		"中兴":    "zte",
		"中兴努比亚": "nubia",
		"酷派":    "coolpad",
		"联想":    "lenovo",
		"魅族":    "meizu",
		"HTC":   "htc",
		"OPPO":  "oppo",
		"VIVO":  "vivo",
		"摩托罗拉":  "motorola",
		"索尼":    "sony",
		"LG":    "lg",
		"金立":    "jinli",
		"天语":    "tianyu",
		"诺基亚":   "nokia",
		"美图秀秀":  "meitu",
		"谷歌":    "google",
		"TCL":   "tcl",
		"锤子手机":  "chuizi",
		"一加手机":  "1+",
		"中国移动":  "chinamobile",
		"昂达":    "angda",
		"邦华":    "banghua",
		"波导":    "bird",
		"长虹":    "changhong",
		"大可乐":   "dakele",
		"朵唯":    "doov",
		"海尔":    "haier",
		"海信":    "hisense",
		"康佳":    "konka",
		"酷比魔方":  "kubimofang",
		"米歌":    "mige",
		"欧博信":   "ouboxin",
		"欧新":    "ouxin",
		"飞利浦":   "philip",
		"维图":    "voto",
		"小辣椒":   "xiaolajiao",
		"夏新":    "xiaxin",
		"亿通":    "yitong",
		"语信":    "yuxin",
	}

	PriceMap = map[string]string{
		"0-999":     "0-999",
		"1000-1999": "1000-1999",
		"2000-3999": "2000-3999",
		"4000+":     "4000+",
	}
)