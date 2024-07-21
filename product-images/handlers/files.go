package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/daffaromero/gomicro/product-images/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type Files struct {
	log   hclog.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}
func (f *Files) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, w, r)
}

func (f *Files) invalidURI(uri string, w http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(w, "Invalid file path, should be in the format: /[id]/[filename]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, w http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
	}
}
