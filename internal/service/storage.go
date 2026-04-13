package service

import ("fmt"; "io"; "os"; "path/filepath"; "time")

type StorageService struct{ baseDir string }

func NewStorageService(dir string) *StorageService {
	_ = os.MkdirAll(dir, 0755)
	return &StorageService{baseDir: dir}
}

func (s *StorageService) Save(filename string, src io.Reader) (string, error) {
	dir := filepath.Join(s.baseDir, time.Now().Format("2006/01"))
	_ = os.MkdirAll(dir, 0755)
	name := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(filename))
	f, err := os.Create(filepath.Join(dir, name))
	if err != nil { return "", err }
	defer f.Close()
	_, err = io.Copy(f, src)
	return "/uploads/" + name, err
}

func (s *StorageService) Delete(filename string) error {
	return os.Remove(filepath.Join(s.baseDir, filepath.Base(filename)))
}
