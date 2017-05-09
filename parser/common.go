package parser

type ConfRule struct {
	Type  string `json:"type" bson:"type"`
	Name  string `json:"name" bson:"name"`
	Xpath string `json:"xpath" bson:"xpath"`
	Regex string `json:"regex" bson:"regex"`
}

type ParseConf map[string][]ConfRule

type ParsedItem map[string]interface{}

type ParsedTask struct {
	TaskType string `json:"task_type" bson:"task_type"`
	Url      string `json:"url" bson:"url"`
	Data     string `json:"data" bson:"data"`
}
