package logic

import (
	"github.com/sonnht1409/scanning/service/models"
)

var RegexRules = []models.RegexRule{
	{
		RuleName:  "PublicKeyCheck",
		RuleValue: `\s+public_key\s+=`,
	},
	{
		RuleName:  "PrivateKeyCheck",
		RuleValue: `\s+private_key\s+=`,
	},
}

// func (s ServiceLogic) CheckFileContent(content string) (bool, []int) {

// 	lines := []int{}
// 	contentLines := strings.Split(content, "\n")
// 	for _, line := range contentLines {
// 		for _, rule := range RegexRules {
// 			if
// 		}
// 	}
// 	return true, lines
// }
