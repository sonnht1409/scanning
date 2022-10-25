package logic

import (
	"regexp"
	"strings"

	"github.com/sonnht1409/scanning/service/models"
)

var PUBLIC_RULE = models.RegexRule{
	RuleName:  "PublicKeyCheck",
	RuleValue: `\s+public_key\s+=`,
}

var PRIVATE_RULE = models.RegexRule{
	RuleName:  "PrivateKeyCheck",
	RuleValue: `\s+private_key\s+=`,
}

func (s ServiceLogic) CheckRule(content string, rule models.RegexRule) (bool, []int) {
	if match, err := regexp.MatchString(rule.RuleValue, content); !match || err != nil {
		return false, []int{}
	}

	lines := strings.Split(content, "\n")
	matchedLines := []int{}
	for idx, line := range lines {
		if match, err := regexp.MatchString(rule.RuleValue, line); !match || err != nil {
			continue
		}
		matchedLines = append(matchedLines, idx+1)
	}
	return true, matchedLines
}
