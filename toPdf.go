package photoToPdf

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"math"
	"os"
)
import "path/filepath"
import "github.com/signintech/gopdf"

var (
	src = "./assets"
	as  = "./result.pdf"
)

func main() {
	fmt.Println("Reading files from (", src, ") and saving the result as (", as,")")
	fmt.Println("-----------------------")
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4 })
	jpgs, _ := filepath.Glob(src + "/*.jpg")
	pngs, _ := filepath.Glob(src + "/*.png")
	jpegs, _ := filepath.Glob(src + "/*.jpeg")
	files := append(jpgs, append(jpegs, pngs...)...)
	for i := 0; i < len(files); i++ {
		fmt.Println(i+1, "adding ", files[i])
		x := float64(0)
		if x < 0 {
			continue
		}
		pdf.AddPage()
		goodSize(files[i])
		pdf.Image(files[i], 0, 0, nil)
	}
	fmt.Println("saving to ", as, " ...")
	pdf.WritePdf(as)
	fmt.Println("-----------------------")
	fmt.Println("Done, have fun ;)")
}



const DefaultMaxWidth float64 = 595 * 1.6
const DefaultMaxHeight float64 = 842 * 1.6

// Рассчитываем размер изображения после масштабирования
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	ratio := math.Min(DefaultMaxWidth/float64(srcWidth), DefaultMaxHeight/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

func goodSize(file string) {
	photo, _ := os.Open(file)
	defer photo.Close()

	img, _, err := image.Decode(photo)
	if err != nil {
		return
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	fmt.Println("width = ", width, " height = ", height)
	fmt.Println("w = ", w, " h = ", h)

	// Вызов библиотеки изменения размера изображения
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// файл для сохранения
	imgfile, _ := os.Create(file)
	defer imgfile.Close()

	// Сохраняем файл в формате PNG
	//err = png.Encode(imgfile, m)
	err = jpeg.Encode(imgfile, m, nil)
	if err != nil {
		return
	}
}
