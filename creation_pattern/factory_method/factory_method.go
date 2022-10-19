package factory_method

type IRuleConfigParser interface {
	parse(data []byte) string
}

type jsonConfigParser struct{}

func (j *jsonConfigParser) parse(data []byte) string {
	return "json parse"
}

type yamlConfigParser struct{}

func (y *yamlConfigParser) parse(data []byte) string {
	return "yaml parse"
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonConfigParserFactory struct{}

func (j *jsonConfigParserFactory) CreateParser() IRuleConfigParser {
	return &jsonConfigParser{}
}

type yamlConfigParserFactory struct{}

func (y *yamlConfigParserFactory) CreateParser() IRuleConfigParser {
	return &yamlConfigParser{}
}

func NewIRuleConfigParserFactory(model string) IRuleConfigParserFactory {
	switch model {
	case "json":
		return &jsonConfigParserFactory{}
	case "yaml":
		return &yamlConfigParserFactory{}
	}
	return nil
}

