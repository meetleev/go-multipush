package honor

type StyleType = byte

const (
	// 标准样式
	StyleTypeNormal StyleType = iota
	// 大文本样式
	StyleTypeBigText
)

type ButtonActionType = byte

const (
	// 启动应用
	ButtonActionTypeLaunchApp ButtonActionType = iota
	// 打开应用自定义页面
	ButtonActionTypeJumpCustom
	// 跳转网页
	ButtonActionTypeJumpSite
)

type ClickActionType = byte

const (

	// 打开应用自定义页面
	ClickActionTypeOpenCustom ClickActionType = iota + 1
	// 点击后打开特定URL
	ClickActionTypeOpenUrl
	// 启动应用
	ClickActionTypeLaunchApp
)

type AndroidIntentType = byte

const (
	// 设置通过intent打开应用自定义页面
	AndroidIntentTypeIntent AndroidIntentType = iota
	// 打开应用自定义页面
	AndroidIntentTypeAction
)

type PushCategory = string

const (
	// 资讯营销类消息
	PushCategoryLow PushCategory = "LOW"
	// 服务与通讯类消息
	PushCategoryNormal PushCategory = "NORMAL"
)

// https://developer.honor.com/cn/docs/11002/reference/downlink-message#%E8%8E%B7%E5%8F%96%E9%89%B4%E6%9D%83%E6%8E%A5%E5%8F%A3
type PushNotification struct {
	// 通知栏消息的标题。发送通知栏消息时，此处title和android.notification .title两者最少需要设置一个。
	Title string `json:"title"`
	// 通知栏消息的内容。发送通知栏消息时，此处body和android.notification .body两者最少需要设置一个。
	Body string `json:"body"`
	// 用户自定义的通知栏消息通知小图URL，该字段仅允许在服务通讯类消息中使用；如果不设置，则不展示通知小图。URL使用的协议必须是HTTPS协议
	ImageUrl string `json:"image,omitempty"`
}

type PushAndroidButton struct {
	// 按钮名称，最大长度40。
	Name string `json:"name"`
	// 点击通知栏后触发的动作类型。default: [ButtonActionTypeLaunchApp]
	ActionType ButtonActionType `json:"actionType,omitempty"`
	// 打开自定义页面的方式, 当actionType为1时，该字段必填
	IntentType AndroidIntentType `json:"intentType,omitempty"`
	/*
	* 当actionType为1，此字段按照intentType字段设置应用页面的uri或者action，具体设置方式参见打开应用自定义页面。
	* 当actionType为2，此字段设置打开指定网页的URL，URL使用的协议必须是HTTPS协议
	 */
	IntentUri string `json:"intent,omitempty"`
	// 最大长度1024。当字段actionType为0或1时，该字段用于在点击按钮后给应用透传数据，选填，格式必须为key-value形式
	IntentData string `json:"data,omitempty"`
}

type PushAndroidClickAction struct {
	// 	消息点击行为类型
	ActionType ClickActionType `json:"type"`
	/*
	* 设置打开特定URL，本字段填写需要打开的URL，URL使用的协议必须是HTTPS协议，
	* 取值样例：https://example.com/image.png。当type为2时必选。如果是游戏类应用，不支持设置特定URL
	 */
	Url       string `json:"url,omitempty"`
	IntentUri string `json:"intent,omitempty"`
	// 设置通过action打开应用自定义页面时，本字段填写要打开的页面activity对应的action。当type为1（打开自定义页面）时，字段intent和action至少二选一
	IntentAction string `json:"action,omitempty"`
}

type PushBadge struct {
	// 应用角标累加数字（大于0小于100的整数），非应用角标实际显示数
	AddNum uint16 `json:"addNum,omitempty"`
	// 应用入口Activity类全路径。样例：com.example.test.MainActivity
	BadgeClass string `json:"badgeClass"`
	// 角标设置数字（大于等于0小于100的整数），应用角标实际显示数字
	SetNum uint16 `json:"setNum,omitempty"`
}

type PushAndroidNotification struct {
	// 每条消息在通知显示时的唯一标识。不携带时或者设置-1时，Push NC自动为给每条消息生成一个唯一标识；不同的通知栏消息可以拥有相同的notifyId，实现新的消息覆盖上一条消息功能。
	NotifyId string `json:"notifyId,omitempty"`
	// 通知栏样式, default: [StyleTypeNormal]
	Style StyleType `json:"style,omitempty"`
	// 通知栏消息的标题。发送通知栏消息时，此处title和android.notification .title两者最少需要设置一个。
	Title string `json:"title"`
	// 通知栏消息的内容。发送通知栏消息时，此处body和android.notification .body两者最少需要设置一个。
	Body string `json:"body"`
	// 用户自定义的通知栏消息通知小图URL，该字段仅允许在服务通讯类消息中使用；如果不设置，则不展示通知小图。URL使用的协议必须是HTTPS协议
	ImageUrl string `json:"image,omitempty"`

	ClickAction *PushAndroidClickAction `json:"clickAction"`

	// Android通知栏消息大文本标题，当style为1时必选，设置bigTitle后通知栏展示时，bigTitle设置后内容要与title一致
	BigTitle string `json:"bigTitle"`
	// Android通知栏消息大文本内容，当style为1时必选，设置bigBody后通知栏展示时，bigBody设置后内容要与body一致
	BigBody string `json:"bigBody"`

	// 设置通知栏消息的到达时间，如果您同时发送多条消息，Android通知栏中的消息根据这个值进行排序，同时将排序后的消息在通知栏上显示。该时间戳为UTC时间戳，样例：2014-10-02T15:01:23.045123456Z。
	When int64 `json:"when,omitempty"`
	// 通知栏消息动作按钮，最多设置3个
	Buttons []PushAndroidButton `json:"buttons,omitempty"`
	// 	Android通知消息角标控制
	Badge *PushBadge `json:"badge,omitempty"`
	// 消息标签，同一应用下使用同一个消息标签的消息会相互覆盖，只展示最新的一条。
	Tag string `json:"tag,omitempty"`
	// 消息分组，例如发送10条带有同样group字段的消息，手机上只会展示该组消息中最新的一条和当前该组接收到的消息总数目，不会展示10条消息。
	Group string `json:"off_line,omitempty"`

	Category PushCategory `json:"importance,omitempty"`
}

type PushAndroidOptions struct {
	/*
	* 消息缓存时间，单位是秒。在用户设备离线时，消息在Push服务器进行缓存，
	* 在消息缓存时间内用户设备上线，消息会下发，超过缓存时间后消息会丢弃，默认值为“86400s”（1天），最大值为“1296000s”（15天）。
	 */
	Ttl string `json:"ttl,omitempty"`
	// 批量任务消息标识，消息回执时会返回给应用服务器，应用服务器可以识别biTag对消息的下发情况进行统计分析。
	BiTag string `json:"biTag,omitempty"`
	// 	Android通知栏消息结构体
	Notification *PushAndroidNotification `json:"notification,omitempty"`
	// 0：普通消息（默认值）1：测试消息。每个应用每日可发送该测试消息1000条且不受每日单设备推送数量上限要求
	TargetUserType byte `json:"targetUserType,omitempty"`
}

type PushMessage struct {
	// 通知栏消息内容
	Notification *PushNotification `json:"notification,omitempty"`
	// 	自定义消息负载，通知栏消息支持JSON格式字符串，透传消息支持普通字符串或者JSON格式字符串
	Payload string   `json:"data,omitempty"`
	Token   []string `json:"token"`
	// 	Android消息推送控制参数
	AndroidOptions *PushAndroidOptions `json:"android,omitempty"`
}

type PushMessageResp struct {
	// 响应码
	Code int `json:"code"`
	// 响应信息
	Message int `json:"message,omitempty"`
	Data    struct {
		// 推送消息的结果
		SendResult bool `json:"sendResult"`
		// 请求ID
		RequestId string `json:"requestId,omitempty"`
		// 推送失败的pushTokens
		FailTokens []string `json:"failTokens,omitempty"`
		// 失效的pushTokens
		ExpireTokens []string `json:"expireTokens,omitempty"`
	} `json:"data,omitempty"`
}
