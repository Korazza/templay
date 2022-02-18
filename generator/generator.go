package generator

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template/parse"
)

func listTemplFields(t *template.Template) []string {
	return listNodeFields(t.Tree.Root)
}

func listNodeFields(node parse.Node) []string {
	var res []string
	if node.Type() == parse.NodeAction {
		field := strings.ReplaceAll(node.String(), " | _html_template_htmlescaper", "")
		res = append(res, field[3:len(field)-2])
	}
	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = append(res, listNodeFields(n)...)
		}
	}
	return res
}

func validateTemplay(t *template.Template, vars map[string]interface{}) error {
	tvars := listTemplFields(t)
	var valid int
	var validvars []int
	for _, k := range tvars {
		if _, ok := vars[k]; ok {
			validvars = append(validvars, valid)
			valid++
		}
	}
	if valid != len(tvars) {
		for i := range validvars {
			tvars[i] = tvars[len(tvars)-1]
			tvars = tvars[:len(tvars)-1]
		}
		var errors []string
		for _, v := range tvars {
			errors = append(errors, fmt.Sprintf("%3s %s", "-", v))
		}
		return fmt.Errorf("templay \"%s\" requires more variables: \n%v", t.Name(), strings.Join(errors, "\n"))
	}
	return nil
}

func ParseFile(src, dst string, vars map[string]interface{}) error {
	var (
		err     error
		srcfd   *os.File
		dstfd   *os.File
		srcinfo os.FileInfo
		templay *template.Template
	)

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	templay, err = template.ParseFiles(src)
	if err != nil {
		return err
	}

	if err = validateTemplay(templay, vars); err != nil {
		return err
	}

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	err = templay.Execute(dstfd, vars)
	if err != nil {
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
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

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
		errMsg := strings.Join(errors, "\n\n")
		return fmt.Errorf("\n%s", errMsg)
	}
	return nil
}
