package simplecache

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// SimlpleCache struct
type SimlpleCache struct {
	Folder   string
	Dir      string
	fullpath string
}

// NewSimpleCache new a SimlpleCache
func NewSimpleCache(Dir, Folder string) *SimlpleCache {
	if Dir == "" {
		Dir = getOSRamdiskPath()
	}

	fullpath := filepath.Join(Dir, Folder)

	if err := os.Mkdir(fullpath, 0700); os.IsNotExist(err) {
		fullpath = filepath.Join(os.TempDir(), Folder)
		os.Mkdir(fullpath, 0700)
	}

	return &SimlpleCache{
		Dir:      Dir,
		Folder:   Folder,
		fullpath: fullpath,
	}
}

// Get get cache
func (s *SimlpleCache) Get(name string) ([]byte, error) {
	var err error
	if file, err := os.Open(filepath.Join(s.fullpath, name)); err == nil {
		defer file.Close()
		return ioutil.ReadAll(file)
	}
	return nil, err
}

// Set data
func (s *SimlpleCache) Set(name string, data []byte) error {
	var err error
	if file, err := os.Create(filepath.Join(s.fullpath, name)); err == nil {
		defer file.Close()
		file.Write(data)
	}
	return err
}

func getOSRamdiskPath() string {
	switch runtime.GOOS {
	case "linux":
		return "/run/shm/"
	default:
		return os.TempDir()
	}
}
