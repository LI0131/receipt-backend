package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/li0131/receipt-backend/internal/config"
	"github.com/otiai10/gosseract/v2"
)

func UploadFile(cfg config.FileServerConfig, w http.ResponseWriter, r *http.Request) {
	client := gosseract.NewClient()
	defer client.Close()

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		NotFound(w, r)
	}
	defer file.Close()

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	newFile, err := os.Create(cfg.FileLocation + "/" + handler.Filename)
	if err != nil {
		InternalServerError(w)
	}
	defer newFile.Close()

	log.Printf("Created File...")

	// copy contents of file into newfile
	if _, err := io.Copy(newFile, file); err != nil {
		InternalServerError(w)
	}

	log.Printf("Copied Contents File...")

	// scan the image
	client.SetImage(newFile.Name())
	text, err := client.Text()
	if err != nil {
		InternalServerError(w)
	}

	// return that we have successfully uploaded our file!
	Success(w, text)
}
