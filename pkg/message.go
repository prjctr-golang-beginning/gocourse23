package pkg

type Message struct {
	DeletedID     int                    `json:"id"`
	RecordID      interface{}            `json:"record_id"`
	GroupName     string                 `json:"group_name"`
	GoodsID       interface{}            `json:"record_from_id"`
	RelationID    interface{}            `json:"record_to_id"`
	Name          string                 `json:"name"`
	Variant       string                 `json:"variant"`
	FieldsData    map[string]interface{} `json:"fields_data"`
	ChangedFields map[string]interface{} `json:"changed_fields"`
	SourceData    SourceData             `json:"source_data"`
}

type SourceData struct {
	Request string `json:"request"`
}

var Example = `{
	"id": 921,
	"record_id": "IUOgiuyt(&^gi76gI76f76",
	"group_name": "Products",
	"record_from_id": 217,
	"record_to_id": 12,
	"name": "Tourizm",
	"variant": "tourism.jpg",
	"fields_data": "",
	"changed_fields": "id, name",
	"source_data": {
		"request": "BO^*gf87^F*75TVU^r5Dc6rTVu^T5vu^5IU&"
	}
}`
