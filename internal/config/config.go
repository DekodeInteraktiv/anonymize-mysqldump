package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DekodeInteraktiv/anonymize-mysqldump/internal/embed"
)

type Config struct {
	Name        string
	ProcessName string
	Version     string
	Commit      string
	Date        string
	WD          string
	Patterns    []ConfigPattern `json:"patterns"`
}

type ConfigPattern struct {
	TableName      string         `json:"tableName"`
	TableNameRegex string         `json:"tableNameRegex"`
	Purge          bool           `json:"purge"`
	Fields         []PatternField `json:"fields"`
}

type PatternField struct {
	Field       string                   `json:"field"`
	Position    int                      `json:"position"`
	Type        string                   `json:"type"`
	Constraints []PatternFieldConstraint `json:"constraints"`
}

type PatternFieldConstraint struct {
	Field    string `json:"field"`
	Position int    `json:"position"`
	Value    string `json:"value"`
}

// New creates a new Config from flags and environment variables
func New(version, commit, date string) *Config {
	c := &Config{
		Name:        "Anomymize MySQLDump",
		ProcessName: filepath.Base(os.Args[0]),
		Version:     version,
		Commit:      commit,
		Date:        date,
	}

	// Get Working Dir
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	c.WD = wd

	return c
}

// ParseConfig parses a default or user provided config file.
func (c *Config) ParseConfig(filepath string, presets []string) {
	var jsonConfig []byte
	var defaultConfig []byte
	var jsonParser *json.Decoder
	var jsonReader *strings.Reader
	var err error

	// Load default config.
	defaultConfig, err = embed.Content.ReadFile("files/default.json")
	if err != nil {
		log.Fatalln("Failed reading default config.")
	}

	jsonConfig = []byte(defaultConfig)

	if filepath != "" {
		jsonConfig, err = ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatalf("Failed reading config file: %s", err)
		}
	}

	jsonReader = strings.NewReader(string(jsonConfig))
	jsonParser = json.NewDecoder(jsonReader)
	err = jsonParser.Decode(c)
	if err != nil {
		log.Fatalln("JSON file not valid!")
	}

	if len(presets) != 0 {
		var presetConfig []byte
		for _, preset := range presets {
			presetConfig, err = embed.Content.ReadFile("files/config." + preset + ".json")
			if err != nil {
				log.Fatalf("No preset config found for: %s.", preset)
			}

			tempConfig := &Config{}

			jsonReader = strings.NewReader(string(presetConfig))
			jsonParser = json.NewDecoder(jsonReader)
			err = jsonParser.Decode(tempConfig)
			if err != nil {
				log.Fatalln("JSON file not valid!")
			}

			c.Patterns = append(c.Patterns, tempConfig.Patterns...)
		}
	}
}
