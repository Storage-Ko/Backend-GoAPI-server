package file

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Backend-GoAPI-server/utils"
	"github.com/gorilla/mux"
)

type fileType struct {
	Filename string `json:"filename"`
}

func UploadsHandler(rw http.ResponseWriter, r *http.Request) {
	// Read file from Request
	uploadFile, header, err := r.FormFile("file")
	utils.HandleErr(err)
	if err != nil {
		utils.BadRequestException(rw)
		return
	}

	// Generate Dir for save file
	dirname := "./public"
	os.MkdirAll(dirname, 0777)

	// Save file
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	io.Copy(file, uploadFile)

	data := fileType{
		Filename: header.Filename,
	}
	utils.MarshalAndRW(201, data, rw)
}

func LoadsFile(rw http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	http.ServeFile(rw, r, "./public/"+path["path"])
	return
}
