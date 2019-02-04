package request

type Header struct {
	Name      string      `json:"name"`
	Value     string      `json:"value"`
	ValueFrom interface{} `json:"valueFrom"`
}
