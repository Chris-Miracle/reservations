package models

// TemplateData is the data sent to templates
type TemplateData struct {
	StrinpMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}