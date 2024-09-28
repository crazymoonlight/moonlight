package elk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

func SearchElkLog(l SearchLogReq) (*elastic.SearchResult, error) {
	if l.TenantId == "" || l.AppId == "" || l.EventType == "" || l.Time == "" {
		return nil, errors.New("参数不能为空")
	}
	url := "http://" + l.Host + ":" + l.Port
	// 创建 Elasticsearch 客户端
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Printf("Error creating the client:" + err.Error())
		panic(err.Error())
	}
	// 构建查询条件
	query := elastic.NewBoolQuery()
	buildQuery(query, l)
	//拼分区
	index := l.AppId + "-" + l.TenantId + "-" + l.EventType + "-" + l.Time
	// 检查连接
	//info, code, err := client.Ping("http://120.78.84.119:9200").Do(context.Background())
	//if err != nil {
	//	log.Error("Error pinging Elasticsearch:" + err.Error())
	//	panic(err.Error())
	//}
	resp, err := client.Search().
		//分页
		From((l.PageNum) * l.PageSize).
		Size(l.PageSize).
		Index(index).
		Query(query).
		// 执行查询
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return resp, nil
}
func buildQuery(query *elastic.BoolQuery, l SearchLogReq) {
	if l.Keyword != "" {
		// 使用 MultiMatchQuery 进行模糊匹配
		multiMatchQuery := elastic.NewMultiMatchQuery(l.Keyword,
			"userName", "comportment", "msg") // 指定要匹配的字段
		query.Must(multiMatchQuery) // 使用 Should 逻辑运算符
	}
	if l.UserId != "" {
		query.Must(elastic.NewTermQuery("userId", l.UserId)) // 精确匹配 user_id
	}
	if l.LogLevel != "" {
		query.Must(elastic.NewTermQuery("logLevel", l.LogLevel)) // 精确匹配 logLevel
	}
	if l.EventType != "" {
		query.Must(elastic.NewTermQuery("eventType", l.EventType)) // 精确匹配 eventType
	}
	if l.IPAddress != "" {
		query.Must(elastic.NewMatchQuery("ipAddress", l.IPAddress)) // 模糊匹配 ipAddress
	}
	if l.IPUrl != "" {
		query.Must(elastic.NewMatchQuery("ipUrl", l.IPUrl)) // 模糊匹配 ipUrl
	}
}
func PostElkLog(l LogReq) (string, error) {
	url := "http://" + l.Host + ":" + l.Port
	marshal, err := json.Marshal(l)
	if err != nil {
		log.Printf(err.Error())
		panic(err.Error())
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	return "上传成功", nil
}
