package base

type BaseClientContract interface {
	Get(path string, query map[string]string, returnObject interface{}) error
}
