package api

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/thecodingmachine/gotenberg/internal/pkg/printer"
)

var officeExts = []string{
	".txt",
	".rtf",
	".doc",
	".docx",
	".odt",
	".xls",
	".xlsx",
	".ods",
	".ppt",
	".pptx",
	".odp",
}

func convertOffice(c echo.Context) error {
	r, err := newResource(c)
	if err != nil {
		return err
	}
	defer r.removeAll()
	ctx, cancel := newContext(r)
	if cancel != nil {
		defer cancel()
	}
	fpaths, err := r.filePaths(officeExts)
	if err != nil {
		return err
	}
	if len(fpaths) == 0 {
		return errors.New("no suitable office documents to convert")
	}
	p := &printer.Office{Context: ctx, FilePaths: fpaths}
	return print(c, p, r)
}
