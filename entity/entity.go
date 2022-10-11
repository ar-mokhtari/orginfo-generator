package entity

type (
	Domain struct {
		LowerName  string
		UpperName  string
		SnakeName  string
		AllLowName string
		Fields     []Field
	}
	Field struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		JsonName string `json:"json_name"`
	}
)
