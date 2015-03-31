# white-noise-generator
A simple white noise generator written in Go.

### Introduction
White noise is a sound of all frequencies in the range of human hearing in equal intensity.

This generator uses random numbers of a CSPRNG to generate white noise.

### Build
```sh
git clone https://github.com/jannispinter/white-noise-generator.git
cd white-noise-generator
go get && go build```

### Usage
You can execute the program with it's default options:
```sh
./white-noise-generator```
A file named "white_noise.wav" will be written into your current working directory.

To change the default options, use these switches:
```Usage of ./white-noise-generator:
  -bits=32: Sample rate
  -duration=60: Duration in seconds
  -filename="white_noise.wav": Output file name
  -rate=44100: Bitrate
```

### See also
https://github.com/cryptix/wav - A library for Go to read and write wav files.
http://mynoise.net - A website to generate all kinds of sounds (including white noise)
https://en.wikipedia.org/wiki/White_noise - The Wikipedia article about white noise
