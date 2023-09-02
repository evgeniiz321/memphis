package memphis_test

import (
	"os"
	"testing"

	mem "github.com/evgeniiz321/memphis"
)

func TestRename(t *testing.T) {
	fs := mem.FromOS(t.TempDir())
	bfs := fs.AsBillyFS(0, 0)
	f, err := bfs.Create("f1")
	if err != nil {
		t.Fatal(err)
	}
	prev_contents, err := bfs.ReadDir("")
	if err != nil {
		t.Fatal(err)
	}
	if prev_contents[0].Name() != "f1" {
		t.Fatal(err)
	}

	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	err = bfs.Rename("f1", "f2")
	if err != nil {
		t.Fatal(err)
	}
	new_contents, err := bfs.ReadDir("")
	if err != nil {
		t.Fatal(err)
	}
	if new_contents[0].Name() != "f2" {
		t.Fatal("read dir after rename returns old data")
	}
}

func TestSymlinkToNonExistentFile(t *testing.T) {
	fs := mem.FromOS(t.TempDir())
	bfs := fs.AsBillyFS(0, 0)
	err := bfs.Symlink("f1", "f2")
	if err != nil {
		t.Fatal(err)
	}

	err = bfs.Chmod("f2", os.FileMode(0777))
	if err != nil {
		t.Fatal(err)
	}
}
