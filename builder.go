package godis

type Builder interface {
	build(data interface{}) interface{}
}

var (
	STRING_BUILDER       = newStringBuilder()
	STRING_ARRAY_BUILDER = newStringArrayBuilder()
)

type stringBuilder struct {
}

func newStringBuilder() *stringBuilder {
	return &stringBuilder{}
}

func (b *stringBuilder) build(data interface{}) interface{} {
	if data == nil {
		return ""
	}
	return string(data.([]byte))
}

type stringArrayBuilder struct {
}

func newStringArrayBuilder() *stringArrayBuilder {
	return &stringArrayBuilder{}
}

func (b *stringArrayBuilder) build(data interface{}) interface{} {
	if data == nil {
		return []string{}
	}
	arr := make([]string, 0)
	for _, b := range data.([]interface{}) {
		if b == nil {
			arr = append(arr, "")
		} else {
			arr = append(arr, string(b.([]byte)))
		}
	}
	return arr
}
