package main

//#include <stdlib.h>
import (
	"bytes"
	"C"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"os/exec"
	"fmt"
)


//export convert
func convert(html *C.char){
	go_html := []byte(C.GoString(html))
	go_pdfFile := "./output_pdfs/md.pdf"
	// fmt.Println(go_html)

	pdfg, _:= wkhtmltopdf.NewPDFGenerator()
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(go_html))
	page.NoBackground.Set(true)
	page.DisableExternalLinks.Set(false)
	pdfg.AddPage(page)
	pdfg.Dpi.Set(350)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	pdfg.Create()
	pdfg.WriteFile(go_pdfFile)
	cmd := exec.Command("python3", "pdf_merge.py")
	out, _ := cmd.CombinedOutput()
	// if err != nil {
	//     fmt.Println(err)
	// }
	fmt.Println(string(out))

}

func main() {}