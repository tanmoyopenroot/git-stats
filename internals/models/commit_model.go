package models

// CommitModel ... Git Commit Model
type CommitModel struct {
	Tree   string
	Parent string
	Author
	Commiter
	Message string
	Type    bool
}
