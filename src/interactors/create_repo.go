package interactors

import (
	"os"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

func init() {

}

func localNotesDirectory() billy.Filesystem {
	userHomeDir, _ := os.UserHomeDir()

	notesDir := osfs.New(userHomeDir + "/.caderno-notes")

	storageToUse := filesystem.NewStorage(
		notesDir,
		cache.NewObjectLRUDefault())

	storageToUse.Init()

	return notesDir
}

func createNoteDirStorage() {
	localNotesDirectory().MkdirAll("my-notes-repo", 0755)
}

// CreateGit -> Create Git Reoo
func CreateGit() {
	createNoteDirStorage()
}
