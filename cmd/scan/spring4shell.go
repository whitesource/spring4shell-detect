package scan

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/whitesource/spring4shell-detect/records"
)

type Vulnerability struct {
	CveId         string
	Summary       string
	FixResolution string
}

var fixes = map[string]map[string]string{
	"org.springframework": {
		"spring-beans": "Upgrade to version org.springframework:spring-beans:5.2.20.RELEASE,5.3.18",
	},
}

//go:embed cve/libs.json
var libs []byte

type Sha1ToLib map[string]records.VulnerableLib

type CveToSha1ToLib map[string]Sha1ToLib

var cve2Sha2Lib CveToSha1ToLib

func init() {
	err := json.Unmarshal(libs, &cve2Sha2Lib)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal libraries: %v", err))
	}
}
