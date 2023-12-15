package pattern

import "fmt"

/* Паттерн стратегия используется для быстрого инкапсулированного переключения какого-то контекста. В данном случае мы используем
его для переключения способов компрессии изображения, что может послужить быстрому и эффективному написанию кода.

В реальных примерах данный паттерн можно использовать как средство инкапсуляции объектов, которое будет скрывать ненужную и
громоздкую реализацию капота структуры.

К плюсам можно отнести полезность при разработке средних и крупных проектов, в особенности при работе с командой, так
как от них можно скрыть ненужную реализацию, что обеспечит лучшую читаемость кода.

К минусам можно отнести излишность в некоторых случаях. В особенности, при написании небольших скриптовых алгоритмов.
*/

type CompressionStrategy interface {
	Compress(data []byte) []byte
}

type JpegCompression struct{}

func (j *JpegCompression) Compress(data []byte) []byte {
	fmt.Println("Applying JPEG compression")
	return append(data, []byte(" (JPEG compressed)")...)
}

type PngCompression struct{}

func (p *PngCompression) Compress(data []byte) []byte {
	fmt.Println("Applying PNG compression")
	return append(data, []byte(" (PNG compressed)")...)
}

type ImageProcessor struct {
	compressionStrategy CompressionStrategy
}

func (ip *ImageProcessor) SetCompressionStrategy(strategy CompressionStrategy) {
	ip.compressionStrategy = strategy
}

func (ip *ImageProcessor) CompressImage(data []byte) []byte {
	if ip.compressionStrategy == nil {
		fmt.Println("Please set a compression strategy first")
		return data
	}
	return ip.compressionStrategy.Compress(data)
}

func StrategyMain() {
	imageProcessor := &ImageProcessor{}

	imageProcessor.SetCompressionStrategy(&JpegCompression{})
	result := imageProcessor.CompressImage([]byte("Image data"))
	fmt.Println(string(result))

	imageProcessor.SetCompressionStrategy(&PngCompression{})
	result = imageProcessor.CompressImage([]byte("Image data"))
	fmt.Println(string(result))
}
