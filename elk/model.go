package elk

type LogReq struct {
	Organization string `json:"organization,omitempty"` //组织
	AppId        string `json:"appId,omitempty"`        //应用id
	LogTime      string `json:"logTime,omitempty"`      //时间
	UserId       string `json:"userId,omitempty"`       //用户id
	UserName     string `json:"userName,omitempty"`     //用户名
	Msg          string `json:"msg,omitempty"`          //日志信息
	EventType    string `json:"eventType,omitempty"`    //事件类型
	LogLevel     string `json:"logLevel,omitempty"`     //日志级别
	Comportment  string `json:"comportment,omitempty"`  //行为
	IPUrl        string `json:"ipUrl,omitempty"`        //请求url
	IPAddress    string `json:"ipAddress,omitempty"`    //ip地址
	TenantId     string `json:"tenantId,omitempty"`     //租户id
	Business     string `json:"business,omitempty"`     //业务类型
	Port         string `json:"port,omitempty"`         //端口
	Host         string `json:"host,omitempty"`         //主机名
}
type SearchLogReq struct {
	Time        string `json:"time,omitempty"`        //时间
	Keyword     string `json:"keyword,omitempty"`     //关键字
	AppId       string `json:"appId,omitempty"`       //应用id
	LogTime     string `json:"logTime,omitempty"`     //时间
	UserId      string `json:"userId,omitempty"`      //用户id
	UserName    string `json:"userName,omitempty"`    //用户名
	Msg         string `json:"msg,omitempty"`         //日志信息
	EventType   string `json:"eventType,omitempty"`   //事件类型
	LogLevel    string `json:"logLevel,omitempty"`    //日志级别
	Comportment string `json:"comportment,omitempty"` //行为
	IPUrl       string `json:"ipUrl,omitempty"`       //请求url
	IPAddress   string `json:"ipAddress,omitempty"`   //ip地址
	TenantId    string `json:"tenantId,omitempty"`    //租户id
	PageNum     int    `json:"pageNum,omitempty"`     //页码
	PageSize    int    `json:"pageSize,omitempty"`    //每页数量
	Business    string `json:"business,omitempty"`    //业务类型
	Port        string `json:"port,omitempty"`        //端口
	Host        string `json:"host,omitempty"`        //主机名
}
