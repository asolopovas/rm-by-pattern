package rmByPattern

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Patterns        []string `yaml:"patterns"`
	Extensions      []string `yaml:"extentions"`
	ExcludePatterns []string `yaml:"excludePatterns"`
}

func GetYamlConfig(path string) (*Config, error) {
	result := Config{}
	yamlConfig, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(yamlConfig, &result)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return &result, nil
}

// func countMatches(text string, pattern string) (int, error) {
// 	re, err := regexp.Compile(pattern)
// 	if err != nil {
// 		return 0, err // Return 0 and the compilation error
// 	}

// 	matches := re.FindAllString(text, -1) // Find all matches
// 	return len(matches), nil              // Return the count of matches and no error
// }

func RmFiles(dirPath string, patterns []string, testMode bool) error {
	mainRegex := regexp.MustCompile(`\d+x\d+`)

	var regexPatterns []*regexp.Regexp
	for _, pattern := range patterns {
		regexPattern, err := regexp.Compile(pattern + "x\\d+\\.")
		if err != nil {
			return err
		}
		regexPatterns = append(regexPatterns, regexPattern)
	}

	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// get filename without extension
		filename := info.Name()

		if mainRegex.MatchString(filename) {
			toBeDeleted := false
			for _, pattern := range regexPatterns {
				if pattern.MatchString(filename) {
					toBeDeleted = true
					break
				}
			}

			if toBeDeleted {
				if testMode {
					fmt.Printf("Would remove file: %s\n", path)
					return nil
				} else if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func IdentifyPatterns(dirPath string, re *regexp.Regexp) ([]string, []string, error) {
	patternMap := make(map[string]bool)
	extensionMap := make(map[string]bool)

	fmt.Println("Scanning: " + dirPath)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			matches := re.FindStringSubmatch(info.Name())
			if len(matches) > 0 {
				parts := strings.Split(matches[0], "x")
				if len(parts) > 0 {
					patternMap[parts[0]] = true
				}
			}
			ext := filepath.Ext(info.Name())
			if ext != "" {
				extensionMap[ext[1:]] = true
			}
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	var patterns, extensions []string
	for pattern := range patternMap {
		patterns = append(patterns, pattern)
	}
	for extension := range extensionMap {
		extensions = append(extensions, extension)
	}

	sort.Slice(patterns, func(i, j int) bool {
		numI, _ := strconv.Atoi(patterns[i])
		numJ, _ := strconv.Atoi(patterns[j])
		return numI < numJ
	})

	sort.Strings(extensions)

	return patterns, extensions, nil
}

func CreateYamlConfig(dirPath string, patterns []string, extentions []string) error {
	config := Config{
		Patterns:   patterns,
		Extensions: extentions,
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	configFile := filepath.Join(dirPath, "rm-by-pattern.yaml")

	fmt.Printf("Creating file: %s\n", configFile)
	return os.WriteFile(configFile, data, 0644)
}

func printVersion() {
	// Read the content of the ./version file
	versionBytes, err := os.ReadFile("./version")
	if err != nil {
		log.Printf("Failed to read version file: %v\n", err)
		return
	}
	fmt.Println(string(versionBytes))
}
