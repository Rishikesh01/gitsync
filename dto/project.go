package dto

type Project struct {
	ProjectName string `json:"project_name,omitempty"`
	GithubLink  string `json:"github_link,omitempty"`
	IsPrivate   bool   `json:"is_private,omitempty"`
	ParentDir   string `json:"parent_dir,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
}
