package simple_factory

type IRuleConfigParser interface {
	parse(data []byte) string
}

type jsonConfigParser struct{}

func (j *jsonConfigParser) parse(data []byte) string {
	return "json parser"
}

type yamlConfigParser struct{}

func (y *yamlConfigParser) parse(data []byte) string {
	return "yaml parser"
}

func NewIRuleConfigParser(model string) IRuleConfigParser {
	switch model {
	case "json":
		return &jsonConfigParser{}
	case "yaml":
		return &yamlConfigParser{}
	}
	return nil
}
