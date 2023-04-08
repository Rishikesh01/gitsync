package dto

type Project struct {
	ProjectName string `json:"project_name,omitempty" yaml:"project_name"`
	GithubLink  string `json:"github_link,omitempty" yaml:"github_link"`
	IsPrivate   bool   `json:"is_private,omitempty" yaml:"is_private"`
	ParentDir   string `json:"parent_dir,omitempty" yaml:"parent_dir"`
	IsActive    bool   `json:"is_active,omitempty" yaml:"is_active"`
}
