package huawei

type PushType = byte

const (
	// Alert消息（通知消息、授权订阅消息）
	PushTypeAlert PushType = iota
	// 卡片刷新消息
	PushTypeFormUpdate
	// 通知扩展消息
	PushTypeExtension
	// 后台消息
	PushTypeBackground PushType = 6
	// 实况窗消息
	PushTypeLiveView PushType = 7
	// 应用内通话消息
	PushTypeVoIPCall PushType = 10
)

type PushCategory = string

// PLAY_VOICE（语音播报）消息仅可发送push-type为2的通知扩展消息。2：通知扩展消息
const (
	// 即时聊天
	PushCategoryIM PushCategory = "IM"
	// 语音通话邀请、视频通话邀请
	PushCategoryVOIP PushCategory = "VOIP"
	// 未接通话消息提醒
	PushCategoryMissCall PushCategory = "MISS_CALL"
	// 订阅
	PushCategorySubscrition PushCategory = "SUBSCRIPTION"
	// 出行
	PushCategoryTravel PushCategory = "TRAVEL"
	// 健康
	PushCategoryHealth PushCategory = "HEALTH"
	// 工作事项提醒
	PushCategoryWork PushCategory = "WORK"
	// 账号动态
	PushCategoryAccount PushCategory = "ACCOUNT"
	// 订单&物流
	PushCategoryExpress PushCategory = "EXPRESS"
	// 财务
	PushCategoryFinance PushCategory = "FINANCE"
	// 设备提醒
	PushCategoryDeviceReminder PushCategory = "DEVICE_REMINDER"
	// 邮件
	PushCategoryMail PushCategory = "MAIL"
	// 语音播报
	PushCategoryPlayVoice PushCategory = "PLAY_VOICE"

	// MARKETING：新闻、内容推荐、社交动态、产品促销、财经动态、生活资讯、调研、功能推荐、运营活动（仅对内容进行标识，不会加快消息发送），统称为资讯营销类消息
	PushCategoryMarking PushCategory = "MARKETING"
)

type PushStyle = byte

const (
	// 普通通知
	PushStyleNormal PushStyle = 0
	// 多行文本样式
	PushStyleMultiLine PushStyle = 3
)

// 消息点击后的行为
type PushActionType = byte

const (
	// 打开应用首页
	PushActionTypeLaunchApp PushActionType = 0
	// 打开应用自定义页面
	PushActionTypeOpenCustom PushActionType = 1
	// 清除通知
	PushActionTypeClearNotication PushActionType = 3
	// 打开拨号界面
	PushActionTypeOpenCallView PushActionType = 5
)

type PushClickAction struct {
	ActionType PushActionType `json:"actionType"`
	// 应用内置页面ability对应的action。当actionType为1时，字段uri和action至少填写一个
	Action string `json:"action,omitempty"`
	/*
	* 应用内置页面ability对应的uri，uri对象内部结构请参见skills标签。
	* 当actionType为1时，字段uri和action至少填写一个。
	* 当存在多个Ability时，分别填写不同Ability的action和uri，优先使用action查找对应的应用内置页面
	 */
	Uri string `json:"uri,omitempty"`
	// 点击时传递给应用的数据, actionType为5时，data必填。固定携带{"tel": "xxx"} value为电话号码，长度最大为20
	Data map[string]any `json:"data,omitempty"`
}

type PushBadge struct {
	// 应用角标累加数字（大于0小于100的整数），非应用角标实际显示数
	AddNum uint16 `json:"addNum,omitempty"`
	// 角标设置数字（大于等于0小于100的整数），应用角标实际显示数字
	SetNum uint16 `json:"setNum,omitempty"`
}

type PushNotification struct {
	Category PushCategory `json:"category"`
	// 通知消息标题
	Title string `json:"title"`
	// 通知消息内容
	Body string `json:"body"`
	// 通知右侧大图标URL，URL使用的协议必须是HTTPS协议, 支持图片格式为PNG、JPG、JPEG、BMP，图片长*宽建议小于128*128像素，若超过49152像素，则图片不展示
	ImageUrl string    `json:"image,omitempty"`
	Style    PushStyle `json:"style,omitempty"`
	// 每条消息在通知显示时的唯一标识。不携带或者设置-1时，推送服务自动为每条消息生成一个唯一标识；不同的通知消息可以拥有相同的notifyId，实现新消息覆盖旧消息功能
	NotifyId uint32 `json:"notifyId,omitempty"`
	// 应用消息的唯一标识，不携带时默认无appMessageId。长度范围为[1,64]，支持大小写字母、数字、+、/、=、-、_和空白字符
	AppMessageId string `json:"appMessageId,omitempty"`
	// 应用内账号id匿名标识，最大长度为64。
	ProfileId string `json:"profileId,omitempty"`
	// 多行文本样式的内容，当style为3时，本字段必填，最多支持3条内容，每条最大长度1024且无法完全展示时以“...”截断
	InboxContent []string `json:"inboxContent,omitempty"`
	// 点击消息动作
	ClickAction *PushClickAction `json:"clickAction"`
	/*
	* 通知消息角标控制参数，详情请参见Badge结构体，不设置时应用不显示角标数字，若当前已存在角标，则角标数字不变化
	 */
	Badge *PushBadge `json:"badge,omitempty"`
	/*
	* 自定义消息通知铃声。此处设置的铃声文件必须放在应用的/resources/rawfile路径下。例如设置为alert.mp3，对应应用本地的/resources/rawfile/alert.mp3文件。支持的文件格式包括MP3、WAV、MPEG等，如果不设置，则用默认系统铃声。
	* 当请求不携带soundDuration字段时，建议铃声时长不超过30秒，若超过30秒则截断处理；当请求携带soundDuration字段时，详情请参见soundDuration字段说明
	 */
	Sound string `json:"sound,omitempty"`
	/*
	* 自定义消息通知铃声时长。需要配合sound字段使用，只有当请求同时携带sound字段，soundDuration字段才会生效。仅支持数字，单位为秒，取值范围 [1, 60]。
	* sound字段传入的自定义消息通知铃声会播放至soundDuration字段值后停止，若自定义消息通知铃声对应的时长不足soundDuration字段值则会循环播放，在达到soundDuration字段值后停止
	 */
	SoundDuration uint16 `json:"soundDuration,omitempty"`
	// 应用在前台时是否展示通知消息。默认为true，表示前后台都展示
	ForegroundShow bool `json:"foregroundShow,omitempty"`
}
type AlertPayload struct {
	Notification *PushNotification `json:"notification"`
}

type PushPayload struct{}
type PushTarget struct {
	Token []string `json:"token"`
}
type PushOptions struct {

	// 测试消息标识
	TestMessage bool `json:"testMessage,omitempty"`
	// 消息缓存时间，单位是秒。在用户设备离线时，消息在Push服务器进行缓存，在消息缓存时间内用户设备上线，消息会下发，超过缓存时间后消息会丢弃，默认值为86400秒（1天），最大值为1296000秒（15天）。
	Ttl uint `json:"ttl,omitempty"`
	// 批量任务消息标识，消息回执时会返回给应用服务器，长度最大64字节。
	BiTag string `json:"biTag,omitempty"`
	// 输入一个唯一的回执ID指定本次下行消息的回执地址及配置，该回执ID可以在配置回执参数中查看
	ReceiptId string `json:"receiptId,omitempty"`
	/*
	* 用户设备离线时，Push服务器对离线消息缓存机制的控制方式，用户设备上线后缓存消息会再次下发，取值如下：
	* -1：对该取值的所有离线消息都缓存（默认值）
	* 0~100：离线消息缓存分组标识，对离线消息进行分组缓存，每个应用每一组最多缓存一条离线消息
	* 如果您发送了10条消息，其中前5条的collapseKey为1，后5条的collapseKey为2，那么待用户上线后collapseKey为1和2的分别下发最新的一条消息给最终用户。
	* 注意
	* collapseKey字段只对push-type为0或2的消息生效。
	* 0：通知消息
	* 2：通知扩展消息
	 */
	CollapseKey int `json:"collapseKey,omitempty"`

	// 后台消息模式，仅对push-type为6的消息生效
	BackgroundMode int `json:"backgroundMode,omitempty"`
}
type PushMessage struct {
	Payload any          `json:"payload"`
	Target  *PushTarget  `json:"target"`
	Options *PushOptions `json:"pushOptions,omitempty"`
}

type PushResp struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	RequestId string `json:"requestId"`
}
