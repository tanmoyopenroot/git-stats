package models

// CommitModel ... Git Commit Model
type CommitModel struct {
	Commit   string
	Tree     string
	Parent   string
	Author   *AuthorModel
	Commiter *CommiterModel
	Message  string
	Type     string
}
