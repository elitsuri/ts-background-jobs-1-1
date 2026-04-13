package storage

import ("fmt"; "io"; "os"; "path/filepath"; "time")

type LocalStorage struct{ baseDir string }

func NewLocalStorage(dir string) *LocalStorage {
	_ = os.MkdirAll(dir, 0755)
	return &LocalStorage{baseDir: dir}
}

func (s *LocalStorage) Save(name string, r io.Reader) (string, error) {
	day := time.Now().Format("2006/01/02")
	dir := filepath.Join(s.baseDir, day)
	_ = os.MkdirAll(dir, 0755)
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(name))
	f, err := os.Create(filepath.Join(dir, filename))
	if err != nil { return "", err }
	defer f.Close()
	_, err = io.Copy(f, r)
	return filename, err
}

func (s *LocalStorage) Delete(name string) error {
	return os.Remove(filepath.Join(s.baseDir, filepath.Base(name)))
}

func (s *LocalStorage) URL(name string) string {
	return "/uploads/" + filepath.Base(name)
}
