package mi

type PassThroughType int8

var (
	PassThroughNotify    PassThroughType = 0 // 通知栏消息
	PassThroughPenetrate PassThroughType = 1 // 透传消息
)

type NotifyTypeType int8

var (
	NotifyTypeAll     NotifyTypeType = -1
	NotifyTypeSound   NotifyTypeType = 1 // 使用默认提示音提示
	NotifyTypeVibrate NotifyTypeType = 2 // 使用默认震动提示
	NotifyTypeLights  NotifyTypeType = 4 // 使用默认led灯光提示
)

type NotifyEffectType string

var (
	NotifyLaunchApp NotifyEffectType = "1" // 打开当前app对应的Launcher Activity
	NotifyActivity  NotifyEffectType = "2" // 打开当前app内的任意一个Activity
	NotifyWeb       NotifyEffectType = "3" // 打开网页
)

type PushMessageExeraCallback struct {
	//  可选项，表示回执类型
	CallBackType int `json:"type,omitempty"`
	//  可选项，自定义回执参数，最大长度256字节
	CallBackParam string `json:"param,omitempty"`
}

type PushMessageExera struct {
	// 可选项，自定义通知栏消息铃声。extra.sound_uri的值设置为铃声的URI。（请参见https://dev.mi.com/xiaomihyperos/documentation/detail?pId=1558#_16）注：铃声文件放在Android app的raw目录下
	SoundUri string `json:"sound_uri,omitempty"`
	// 可选项，开启通知消息在状态栏滚动显示
	Ticker string `json:"ticker,omitempty"`
	// 可选项，开启/关闭app在前台时的通知弹出。当extra.notify_foreground值为”1″时，app会弹出通知栏消息；当extra.notify_foreground值为”0″时，app不会弹出通知栏消息。注：默认情况下会弹出通知栏消息。
	NotifyForeground string `json:"notify_foreground,omitempty"`

	NotifyEffect NotifyEffectType `json:"notify_effect,omitempty"`
	// 可选项，打开当前app的任一组件
	IntentUri string `json:"intent_uri,omitempty"`

	// 可选项，打开某一个网页
	WebUri string `json:"web_uri,omitempty"`
	// 可选项，控制是否需要进行平缓发送。默认不支持。value代表平滑推送的速度。注：服务端支持最低3000/s的qps。也就是说，如果将平滑推送设置为1000，那么真实的推送速度是3000/s
	FlowControl int `json:"flow_control,omitempty"`
	// 可选项，使用推送批次（JobKey）功能聚合消息。客户端会按照jobkey去重，即相同jobkey的消息只展示第一条，其他的消息会被忽略。合法的jobkey由数字（[0-9]），大小写字母（[a-zA-Z]），下划线（_）和中划线（-）组成，长度不大于20个字符，且不能以下划线(_)开头
	Jobkey string `json:"jobkey,omitempty"`
	// 可选项，可以接收消息的设备的语言范围，用逗号分隔
	Locale string `json:"locale,omitempty"`
	// 可选项，无法收到消息的设备的语言范围，逗号分隔
	LocaleNotIn string `json:"locale_not_in,omitempty"`
	// 可以接收消息的app版本号，用逗号分割
	AppVersion string `json:"app_version,omitempty"`
	// 无法接收消息的app版本号，用逗号分割
	AppVersionNotIn string `json:"app_version_not_in,omitempty"`
	// 可选项，指定在特定的网络环境下才能接收到消息。目前仅支持指定”wifi”
	Connpt string `json:"connpt,omitempty"`
	// 可选项，extra.only_send_once的值设置为1，表示该消息仅在设备在线时发送一次，不缓存离线消息进行多次下发
	OnlySendOnce string `json:"only_send_once,omitempty"`

	Callback PushMessageExeraCallback `json:"callback,omitempty"`
}

type PushMessageChannelExera struct {
	PushMessageExera
	// 必填，通知类别的ID。
	ChannelId string `json:"channel_id"`
	// 可选，通知类别的名称。
	ChannelName string `json:"channel_name,omitempty"`
	// 可选，通知类别的描述。
	ChannelDescription string `json:"channel_description,omitempty"`
}

type PushMessage struct {
	// 消息的内容。（注意：需要对payload字符串做urlencode处理）
	Payload string `json:"payload"`
	// App的包名。备注：V2版本支持一个包名，V3版本支持多包名（中间用逗号分割
	RestrictedPackageName string `json:"restricted_package_name"`
	// 通知栏展示的通知的标题，不允许全是空白字符，长度小于50， 一个中英文字符均计算为1 (通知栏消息必填)
	Title string `json:"title"`
	// 通知栏展示的通知的描述，不允许全是空白字符，长度小于128，一个中英文字符均计算为1 (通知栏消息必填)
	Description string `json:"description"`
	// 可选项，消息有效期，单位：毫秒（ms）。当用户设备未联网时，消息默认缓存时间为：公信消息最长1天，私信消息最长10天，超过缓存时间消息会丢弃。
	TimeToLive int64 `json:"time_to_live,omitempty"`
	// 可选项。定时发送消息。用自1970年1月1日以来00:00:00.0 UTC时间表示（以毫秒为单位的时间）。注：仅支持七天内的定时消息。
	TimeToSend int64 `json:"time_to_send,omitempty"`
	// 可选项。默认情况下，通知栏只显示一条推送消息。如果通知栏要显示多条推送消息，需要针对不同的消息设置不同的notify_id（相同notify_id的通知栏消息会覆盖之前的），且要求notify_id为取值在0~2147483647的整数
	NotifyId uint `json:"notify_id,omitempty"`

	Extra any `json:"extra,omitempty"`
}

type PushMessageReq struct {
	*PushMessage
	// 根据registration_id，发送消息到指定设备上。可以提供多个registration_id，发送给一组设备，不同的registration_id之间用“,”分割。参数仅适用于“/message/regid”HTTP API
	RegistrationId string `json:"registration_id,omitempty"`
	// 根据alias，发送消息到指定的设备。可以提供多个alias，发送给一组设备，不同的alias之间用“,”分割。参数仅适用于“/message/alias”HTTP API。
	Alias string `json:"alias,omitempty"`
	// 根据user_account，发送消息给设置了该user_account的所有设备。可以提供多个user_account，user_account之间用“,”分割。参数仅适用于“/message/user_account”HTTP API。
	UserAccount string `json:"user_account,omitempty"`
	// 根据topic，发送消息给订阅了该topic的所有设备。参数仅适用于“/message/topic”HTTP API。
	Topic string `json:"topic,omitempty"`
	// topic列表，使用;$;分割。注: topics参数需要和topic_op参数配合使用，另外topic的数量不能超过5。参数仅适用于“/message/multi_topic”HTTP API。
	Topics string `json:"topics,omitempty"`
	// topic之间的操作关系。支持以下三种：UNION并集INTERSECTION交集EXCEPT差集例如：topics的列表元素是[A, B, C, D]，则并集结果是A∪B∪C∪D，交集的结果是A ∩B ∩C ∩D，差集的结果是A-B-C-D。参数仅适用于“/message/multi_topic”HTTP API
	TopicOp string `json:"topic_op,omitempty"`
}

type PushMessageResp struct {
	Result      string `json:"result"`
	TraceId     string `json:"trace_id"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	Info        string `json:"info"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
}

type HttpError struct {
	error
	Code    int    `json:"code"`
	Message string `json:"message"`
}
