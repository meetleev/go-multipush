package oppo

type StyleType = byte

const (
	// 标准样式
	StyleTypeNormal StyleType = iota + 1
	// 长文本样式（ColorOS版本>5.0可用，通知栏第一条消息可展示全部内容，非第一条消息只展示一行内容）
	StyleTypeLongText
	// （ColorOS版本>5.0可用，通知栏第一条消息展示大图，非第一条消息不显示大图，推送方式仅支持广播，且不支持定速功能
	StyleTypeBigImage
)

type ClickActionType = byte

const (
	// 启动应用
	ClickActionTypeLaunchApp ClickActionType = iota
	// 跳转指定应用内页（action标签名）
	ClickActionTypeJumpAction
	// 跳转网页
	ClickActionTypeJumpSite
	// 跳转指定应用内页（全路径类名）；【非必填，默认值为0】
	ClickActionTypeJumpActivity
	// 跳转Intent scheme URL
	ClickActionTypeJumpScheme
)

type ShowTimeType = byte

const (
	// 即时展示
	ShowTimeTypeImmediate ShowTimeType = iota
	// 定时展示，配置该参数后定时展示开始时间（show_start_time）及定时展示的结束时间（show_end_time）为必填
	ShowTimeTypeCustomTime
)

type PushCategory = string

const (
	// 通讯与服务 begin
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

	// 内容与营销 begin
	// 新闻资讯
	PushCategoryNews PushCategory = "NEWS"
	// 内容推荐
	PushCategoryContent PushCategory = "CONTENT"
	// 运营活动
	PushCategoryMarketing PushCategory = "MARKETING"
	// 社交动态
	PushCategorySocial PushCategory = "SOCIAL"
	// 运营消息场景 end

)

// 通知栏消息提醒等级
type PushNotifyLevel = byte

const (
	// 通知栏
	PushNotifyLevelBar PushNotifyLevel = 1
	// 通知栏+锁屏
	PushNotifyLevelBarAndLockScreen PushNotifyLevel = 2
	// 通知栏+锁屏+横幅+震动+铃声
	PushNotifyLevelBarAndLockScreenAndVibrateRing PushNotifyLevel = 16
)

// PushMessage https://open.oppomobile.com/documentation/page/info?id=11236
type PushMessage struct {
	/* App开发者自定义消息Id，主要用于消息去重。对于广播消息，相同app_message_id只会保存一条；
	 * 对于单推消息，相同app_message_id的消息只会对同一个目标推送一次。
	 */
	AppMseId string `json:"app_message_id,omitempty"`
	// 通知栏样式, default: [StyleTypeNormal]
	Style StyleType `json:"style,omitempty"`
	// 设置在通知栏展示的通知栏标题, 【字数串长度限制在50个字符内，中英文字符及特殊符号（如emoji）均视为一个字符】
	Title string `json:"title"`
	// 子标题，设置在通知栏展示的通知栏标题, 【字符串长度限制在10个字符以内，中英文字符及特殊符号（如emoji）均视为一个字符计算】
	SubTitle string `json:"sub_title,omitempty"`
	/*
	* 设置在通知栏展示的通知的正文内容
	* 1）当选择标准样式（style 设置为 1）时，内容字符串长度限制在50以内；
	* 2）当选择长文本样式（style设置 为 2）时，内容字符串长度限制在128以内；
	* 3）当选择大图样式（style 设置为 3）时，内容字符串长度限制在50以内。
	*【字符串长度计算说明：中英文字符及特殊符号（如emoji）均视作一个字符计算】
	 */
	Content string `json:"content"`
	// 点击通知栏后触发的动作类型。default: [ClickActionTypeLaunchApp]
	ClickActionType ClickActionType `json:"click_action_type,omitempty"`
	// 当设置click_action_type为1或者4时，需要配置本参数。
	ClickActionActivity string `json:"click_action_activity,omitempty"`
	// 跳转URL，当跳转的形式为URL时，click_action_type参数需要设置为2或5，同时设置本参数。本参数接受最大长度2000以内的URL。
	ClickActionUrl string `json:"click_action_url,omitempty"`
	/*
	* 跳转动作参数。
	* 打开应用内页或网页时传递给应用或网页的附加参数【JSON格式】，字符串长度不超过4000。当跳转类型是URL类型时，参数会以URL参数直接拼接在URL后面。
	* 示例：{“key1”:“value1”,“key2”:“value2”}
	 */
	ActionParameters string `json:"action_parameters,omitempty"`
	// 通知栏展示类型。展示类型如下
	ShowTimeType ShowTimeType `json:"show_time_type,omitempty"`
	// 定时展示的开始时间, 13位的unix时间戳。
	ShowStartTime int64 `json:"show_start_time,omitempty"`
	// 定时展示的结束时间, 13位的unix时间戳。
	ShowEndTime int64 `json:"show_end_time,omitempty"`
	// 是否是离线消息 default: true
	OffLine bool `json:"off_line,omitempty"`
	// 离线消息的存活时间，单位是秒。存活时间最大允许设置为10天，参数超过10天以10天传入。 default: 3600
	OffLineTtl int `json:"off_line_ttl,omitempty"`
	// 时区，默认值：（GMT+08:00）北京，香港，新加坡
	TimeZone string `json:"time_zone,omitempty"`

	// 回执功能
	CallBackUrl string `json:"call_back_url,omitempty"`
	// 开发者指定的自定义回执参数。
	CallBackParameter string `json:"call_back_parameter,omitempty"`
	/*
	* 指定下发的通道ID。 通知栏通道（NotificationChannel），从Android9开始，Android设备发送通知栏消息必须要指定通道ID，
	*（如果是快应用，必须带置顶的通道Id:OPPO PUSH推送）default: OPPO PUSH 提供的默认通道ID
	 */
	ChannelId string `json:"channel_id,omitempty"`
	/*
	* 限时展示时间(单位：秒)。
	* 消息在通知栏展示后开始计时，展示时长超过展示事件后，消息会从通知栏中消失。
	* 限时展示的时间范围：公信0-12小时，私信0-24小时。
	 */
	ShowTtl int `json:"show_ttl,omitempty"`

	Category PushCategory `json:"category,omitempty"`
	// 通知栏消息提醒等级
	NotifyLevel PushNotifyLevel `json:"notify_level,omitempty"`
	// 私信模板id。下发对应私信模板时必须携带。不支持自拟。
	PrivateMsgTemplateId string `json:"private_msg_template_id,omitempty"`
	/* 标题模板填充参数。
	* 例：私信模板id标题模板为：欢迎来到$ {city} $ ，$ {city} $ 欢迎您。
	* 此参数内容为：{“city”:“北京”}
	 */
	PrivateTitleParameters map[string]string `json:"private_title_parameters,omitempty"`
	/*
	* 内容模板填充参数。
	* 例：私信模板id对应的内容模板为：欢迎$ {userName} $ 来到$ {city} $
	* 参数内容为：{“userName”:“汤姆”，“city”:“深圳市”}
	 */
	PrivateContentParameters map[string]string `json:"private_content_parameters,omitempty"`
}

func NewPushMessage() *PushMessage {
	return &PushMessage{
		Style:           StyleTypeNormal,
		ClickActionType: ClickActionTypeLaunchApp,
		OffLine:         true,
		OffLineTtl:      3600,
		TimeZone:        "GMT+08:00",
	}
}

type PushSingleMessageReq struct {
	*PushMessage
	TargetType int16 `json:"target_type"`
	// required registration_id or 别名
	TargetValue string `json:"target_value"`
	/*
	* 消息到达客户端后是否校验registration_id。 true表示推送目标与客户端registration_id进行比较，
	* 如果一致则继续展示，不一致则就丢弃；false表示不校验
	 */
	VerifyRegistrationId bool `json:"verify_registration_id,omitempty"`
}

func NewPushSingleMessageReqWithToken(token string) *PushSingleMessageReq {
	return &PushSingleMessageReq{
		TargetType:           2,
		PushMessage:          NewPushMessage(),
		TargetValue:          token,
		VerifyRegistrationId: false}
}
