package main

import (
	"crypto/rand"
	"flag"
	"github.com/cryptix/wav"
	"os"
)

func main() {

	bits := flag.Int("bits", 32, "Sample rate")
	rate := flag.Int("rate", 44100, "Bitrate")
	filename := flag.String("filename", "white_noise.wav", "Output file name")
	duration := flag.Int("duration", 60, "Duration in seconds")
	flag.Parse()

	outputFile, err := os.Create(*filename)
	checkErr(err)
	defer outputFile.Close()

	meta := wav.File{
		Channels:        1,
		SampleRate:      uint32(*rate),
		SignificantBits: uint16(*bits),
	}

	writer, err := meta.NewWriter(outputFile)
	checkErr(err)
	defer writer.Close()

	for i := 0; i < (*duration)*(*rate); i += 1 {
		var sample []byte = make([]byte, *bits/8)

		_, err = rand.Read(sample)
		checkErr(err)

		err := writer.WriteSample(sample)
		checkErr(err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
