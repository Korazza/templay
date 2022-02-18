package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/Korazza/templay/config"
)

func ParseFile(src, dst string, vars map[string]interface{}) error {
	var (
		err     error
		dstfd   *os.File
		srcinfo os.FileInfo
		templay *template.Template
	)

	templay = template.Must(template.ParseFiles(src)).Option("missingkey=error")
	buf := &bytes.Buffer{}

	err = templay.Execute(buf, vars)
	if err != nil {
		return err
	}

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = dstfd.Write(buf.Bytes()); err != nil {
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	return os.Chmod(dst, srcinfo.Mode())
}

func ParseDirectory(src, dst string, vars map[string]interface{}) error {
	var (
		err     error
		fds     []os.FileInfo
		srcinfo os.FileInfo
	)

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}

	var errors []string

	for _, fd := range fds {
		if !fd.IsDir() && !strings.HasSuffix(fd.Name(), config.TEMPLAY_EXTENSION) {
			return fmt.Errorf("templay files need \"%s\" extension", config.TEMPLAY_EXTENSION)
		}
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, strings.TrimSuffix(fd.Name(), config.TEMPLAY_EXTENSION))

		if fd.IsDir() {
			if err = ParseDirectory(srcfp, dstfp, vars); err != nil {
				errors = append(errors, err.Error())
			}
		} else {
			if err = ParseFile(srcfp, dstfp, vars); err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		errMsg := strings.Join(errors, fmt.Sprintf("\n%2s", ""))
		return fmt.Errorf("\n%2s%s", "", errMsg)
	}
	return nil
}
