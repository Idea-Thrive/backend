package config

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/logger"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

const (
	// PREFIX const.
	PREFIX = "backend"
)

// Config struct.
type Config struct {
	HTTP http.Config
	Log  logger.Config
}

// Load function.
func Load(path string) Config {
	var cfg Config

	knf := koanf.New(".")

	// load default configuration
	if err := knf.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default config: %v", err)
	}

	// load configuration from file
	if err := knf.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml: %v", err)
	}

	// load environment variables
	cbr := func(key string, value string) (string, interface{}) {
		finalKey := strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(key, PREFIX)), "__", ".")

		if strings.Contains(value, ",") {
			// remove all the whitespace from value
			// split the value using comma
			finalValue := strings.Split(removeWhitespace(value), ",")

			return finalKey, finalValue
		}

		return finalKey, value
	}
	if err := knf.Load(env.ProviderWithValue(PREFIX, ".", cbr), nil); err != nil {
		log.Printf("error loading environment variables: %v", err)
	}

	if err := knf.Unmarshal("", &cfg); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	indent, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatalf("error marshal config: %v", err)
	}

	indent = pretty.Color(indent, nil)
	cfgStrTemplate := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(cfgStrTemplate, string(indent))

	return cfg
}

// removeWhitespace remove all the whitespaces from the input.
func removeWhitespace(in string) string {
	compile := regexp.MustCompile(`\s+`)

	return compile.ReplaceAllString(in, "")
}
