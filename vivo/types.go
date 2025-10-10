package vivo

type AuthTokenReq struct {
	// 用户申请推送业务时生成的appId
	AppId int `json:"appId"`
	// 用户申请推送业务时获得的appKey
	AppKey string `json:"appKey"`
	// 用户申请推送业务时获得的appSecret
	AppSecret string `json:"appSecret"`
	// Unix13位毫秒时间戳 做签名用
	Timestamp string `json:"timestamp"`
}

type AuthTokenResp struct {
	Result    int    `json:"result"`    // 0 成功，非0失败
	Desc      string `json:"desc"`      // 文字描述接口调用情况
	AuthToken string `json:"authToken"` // 默认有效一天
}

// 通知类型
type PushNotifyType = byte

const (
	// 无
	PushNotifyTypeNone PushNotifyType = iota + 1
	// 响铃
	PushNotifyTypeRing
	// 振动
	PushNotifyTypeVibrate
	// 响铃和振动
	PushNotifyTypeRingAndVibrate
)

// 点击跳转类型
type PushSkipType = byte

const (
	// 打开APP首页
	PushSkipTypeOpenFirstPage PushSkipType = iota + 1
	// 打开链接
	PushSkipTypeOpenUrl
	// 打开app内指定页面
	PushSkipTypeOpenCustomPage PushSkipType = 4
)

type ClassificationType = byte

const (
	// 运营类消息
	ClassificationTypeOperation PushSkipType = iota
	// 系统类消息
	ClassificationTypeSystem
)

type PushMode = byte

const (
	// 正式推送
	PushModeProduct PushMode = iota
	// 测试推送
	PushModeSandbox
)

type PushCategory = string

const (
	// 系统消息场景 begin
	// 即时消息
	PushCategoryIM PushCategory = "IM"
	// 账号与资产
	PushCategoryAccount PushCategory = "ACCOUNT"
	// 日程待办
	PushCategoryTodo PushCategory = "TODO"
	// 设备信息
	PushCategoryDeviceReminder PushCategory = "DEVICE_REMINDER"
	// 订单与物流
	PushCategoryOrder PushCategory = "ORDER"
	// 订阅提醒
	PushCategorySubscription PushCategory = "SUBSCRIPTION"
	// 系统消息场景 end

	// 运营消息场景 begin
	// 内容推荐
	PushCategoryContent PushCategory = "CONTENT"
	// 运营活动
	PushCategoryMarketing PushCategory = "MARKETING"
	// 社交动态
	PushCategorySocial PushCategory = "SOCIAL"
	// 运营消息场景 end

)

type PushExtra struct {
	//  回执ID，新回执对应的回执地址id，如果同时设置callback.id和callback，会根据callback.id进行回执
	CallBackId bool `json:"callback.id"`
	//  第三方自定义回执参数，最大长度192个字符
	CallBackParam string `json:"callback.param"`
}
type TimedDisplay struct {
	// required 表示应设备离线等原因，消息延迟下发到端侧后，超过定时展示时间范围是否展示。设置为true，表示展示。设置为false，表示不展示。
	OvertimeDisplay bool `json:"overtimeDisplay"`
	// required 定时展示开始时间，毫秒时间戳。showStartTime需在请求调用时间十分钟之后，且showStartTime和showEndTime需在请求调用时间之后的24小时内，showStartTime和showEndTime时间间隔需小于120分钟，两时间需在同一天（东八区）, 运营消息展示时间需在7:00-23:00之间。
	ShowStartTime string `json:"showStartTime"`
	// required 定时展示结束时间，毫秒时间戳。showStartTime和showEndTime需在请求调用时间之后的24小时内，showStartTime和showEndTime时间间隔需小于120分钟，两时间需在同一天（东八区）, 运营消息展示时间需在7:00-23:00之间。
	ShowEndTime string `json:"showEndTime"`
}

type PushMessageReq struct {
	// required 用户申请推送业务时生成的appId，用于与获取authToken时传递的appId校验，一致才可以推送
	AppId int `json:"appId"`
	// required
	NotifyType PushNotifyType `json:"notifyType"`
	// required 通知标题（用于通知栏消息） 最大40个字符（中英文字符及特殊符号（如emoji）均视为一个字符计算）
	Title string `json:"title"`
	// required 通知内容（用于通知栏消息） 最大100个字符（中英文字符及特殊符号（如emoji）均视为一个字符计算）
	Content string `json:"content"`
	// 消息缓存时间，单位是秒。在用户设备没有网络时，消息在vivo推送服务器进行缓存，在消息缓存时间内用户设备重新连接网络，消息会下发，超过缓存时间后消息会丢弃。取值至少60秒，最长1天，超过1天默认设置为1天。当值为空时，系统消息默认1天，运营消息默认8小时。
	TimeToLive int `json:"timeToLive,omitempty"`
	// required
	SkipType PushSkipType `json:"skipType"`
	// 跳转内容，跳转类型为2或4时，跳转内容最大2048个字符。
	SkipContent string `json:"skipContent,omitempty"`
	// 网络方式 -1：不限，1：wifi下发送，不填默认为-1
	NetworkType int `json:"networkType,omitempty"`
	// 消息类型 0：运营类消息，1：系统类消息。不填默认为1
	Classification ClassificationType `json:"classification"`
	// 客户端自定义键值对 key和Value键值对总长度不能超过1024字符。app可以按照客户端SDK接入文档获取该键值对
	ClientCustomMap map[string]string `json:"clientCustomMap,omitempty"`
	// 高级特性（详见目录：一.公共——4.高级特性 extra）
	Extra     PushExtra `json:"extra,omitempty"`
	RequestId string    `json:"requestId"`
	// 推送模式 0：正式推送；1：测试推送，不填默认为0（测试推送，只能给web界面录入的测试用户推送；审核中应用，只能用测试推送）
	PushMode PushMode `json:"pushMode"`
	// 二级分类
	Category PushCategory `json:"category,omitempty"`

	// 是否在线直推，设置为true表示是在线直推，false表示非直推。在线直推功能推送时在设备在线下发一次，设备离线直接丢弃。
	SendOnline bool `json:"sendOnline"`
	// 是否前台通知展示，设置为false表示应用在前台则不展示通知消息，true表示无论应用是否在前台都展示通知。不填默认为true
	ForegroundShow bool         `json:"foregroundShow,omitempty"`
	TimedDisplay   TimedDisplay `json:"timedDisplay,omitempty"`
}

func NewPushMessageReq() *PushMessageReq {
	return &PushMessageReq{
		NotifyType:     PushNotifyTypeNone,
		TimeToLive:     86400,
		SkipType:       PushSkipTypeOpenFirstPage,
		NetworkType:    -1,
		Classification: ClassificationTypeSystem,
		PushMode:       PushModeProduct,
		ForegroundShow: true,
	}
}

type PushSingleMessageReq struct {
	PushMessageReq
	// required 应用订阅vivo推送服务器得到的id
	RegId string `json:"regId"`
	// 关联终端设备登录用户标识，最大长度为64
	ProfileId string `json:"profileId,omitempty"`
}

func NewPushSingleMessageReq() *PushSingleMessageReq {
	return &PushSingleMessageReq{
		PushMessageReq: *NewPushMessageReq(),
	}
}

type PushSingleMessageResp struct {
	// 0 表示成功，非0失败
	Result int `json:"result"`
	// 文字描述接口调用情况
	Desc string `json:"desc"`
	// 非法用户信息
	InvalidUsers interface{} `json:"invalidUsers"`
	// 任务ID
	TaskId string `json:"taskId"` // 任务ID
}
type PushMessageResp struct {
	// 0 表示成功，非0失败
	Result int `json:"result"`
	// 文字描述接口调用情况
	Desc string `json:"desc"`
	// 请求ID
	RequestId string `json:"requestId"`
	// 非法用户信息
	InvalidUsers interface{} `json:"invalidUsers"`
	// 任务ID
	TaskId string `json:"taskId"` // 任务ID
}
type SaveMessageToCloudReq = PushMessageReq

type SaveMessageToCloudResp struct {
	// 接口调用是否成功的状态码 0成功，非0失败
	Result int `json:"result"`
	// 文字描述接口调用情况
	Desc string `json:"desc"`
	// 消息ID
	TaskId string `json:"taskId"`
}

type PushMultiMessage struct {
	// required 用户申请推送业务时生成的appId，用于与获取authToken时传递的appId校验，一致才可以推送
	AppId int `json:"appId"`
	// required regId列表 个数大于等于2，小于等于1000
	RegIds []string `json:"regIds"`
	// required 消息ID，取saveListPayload返回的taskId
	TaskId string `json:"taskId"`
	// required 请求唯一标识，最大64字符
	RequestId string `json:"requestId"`
}

type PushMultiMessageResp struct {
	// 0 表示成功，非0失败
	Result int `json:"result"`
	// 文字描述接口调用情况
	Desc string `json:"desc"`
	// 非法用户信息
	InvalidUsers interface{} `json:"invalidUsers"`
}
