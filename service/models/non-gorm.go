package models

type GithubFileContent struct {
	Path    string `json:"path"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

type RegexRule struct {
	RuleName  string `json:"rule_name"`
	RuleValue string `json:"rule_value"`
}
