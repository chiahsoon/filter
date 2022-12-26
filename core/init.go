package core

type DataFilter interface {
	FilterUsingStringFields(data []byte, fields []string) ([]byte, error)
}
