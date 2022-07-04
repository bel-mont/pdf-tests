package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"io"
	"net/http"
	"os"
)

func main() {
	var err error

	// Download a PDF
	fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"
	if err = DownloadFile("example-pdf.pdf", fileUrl); err != nil {
		panic(err)
	}

	fmt.Println("1")
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4

	pdf.AddPage()

	err = pdf.AddTTFFont("daysone", "gomarice_mukasi_mukasi.ttf")
	if err != nil {
		panic(err)
	}
	fmt.Println("2")
	err = pdf.SetFont("daysone", "", 20)
	if err != nil {
		panic(err)
	}

	// Color the page
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(124, 252, 0) //setup fill color
	pdf.RectFromUpperLeftWithStyle(50, 100, 400, 600, "FD")
	pdf.SetFillColor(0, 0, 0)
	fmt.Println("3")
	pdf.SetXY(50, 50)
	pdf.Cell(nil, "ment サンプルテキスト漢字も書ける？")

	// Import page 1
	tpl1 := pdf.ImportPage("example-pdf.pdf", 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 50, 100, 400, 0)

	err = pdf.WritePdf("example.pdf")

	fmt.Println(err)

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
