package filelisting

import (
	"GoDemoProj/errorhandling/filelistingserver/consts"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len(consts.ListPath):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
