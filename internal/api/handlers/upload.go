package handlers

import ("encoding/json"; "fmt"; "io"; "net/http"; "os"; "path/filepath"; "time")

type UploadHandler struct{ UploadDir string }
const maxUploadSize = 10 << 20 // 10MB

func (h *UploadHandler) File(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil { http.Error(w, "file too large", 413); return }
	file, header, err := r.FormFile("file")
	if err != nil { http.Error(w, err.Error(), 400); return }
	defer file.Close()
	dir := filepath.Join(h.UploadDir, time.Now().Format("2006/01/02"))
	_ = os.MkdirAll(dir, 0755)
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(header.Filename))
	dst, err := os.Create(filepath.Join(dir, filename))
	if err != nil { http.Error(w, err.Error(), 500); return }
	defer dst.Close()
	size, _ := io.Copy(dst, file)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url": "/uploads/" + filename,
		"filename": filename,
		"size": size,
		"content_type": header.Header.Get("Content-Type"),
	})
}
