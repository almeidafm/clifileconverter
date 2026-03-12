# clifileconverter

![Go](https://img.shields.io/badge/go-1.25-blue)
![GitHub issues](https://img.shields.io/github/issues/almeidafm/clifileconverter)

`clifileconverter` is a command-line utility for converting common file types into other compatible formats. It supports batch conversion of images, audio, and video files directly from the terminal.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Supported Formats](#supported-formats)
- [Examples](#examples)
- [Uninstall](#uninstall)
- [License](#license)

## Features

- Supports image, audio, and video file conversion.
- Batch processing for converting multiple files in a single command
- Minimal command-line interface designed for simplicity and scripting

## Requirements

- Go 1.25 or newer
- [FFmpeg](https://ffmpeg.org/download.html)

FFmpeg must be installed and available in the system `PATH`.

## Installation

Install with the following command:

```bash
curl -fsSL https://raw.githubusercontent.com/almeidafm/clifileconverter/main/install.sh | sh
```

## Usage

Basic command format:

```bash
clifileconverter [files...] --to FORMAT
```

`FORMAT` defines the target format.

The tool generates a new file with the same base name and the new extension.

## Supported Formats

Image:

- jpg / jpeg
- png
- webp

Audio:

- mp3
- wav
- flac

Video:

- mp4
- mkv
- webm
- mov

## Examples

The program reads input files from the current working directory and writes the converted files to the same directory.

Each output file keeps the original base name and only changes the extension.

Convert an image:

```bash
clifileconverter image.png --to jpg
```

Convert multiple images:

```bash
clifileconverter img1.png img2.png img3.png --to webp
```

Convert an audio:

```bash
clifileconverter audio.wav --to mp3
```

Convert multiple audios:

```bash
clifileconverter audio.wav audio.flac --to mp3
```

Convert a video:

```bash
clifileconverter video.mov --to mp4
```

Convert multiple videos:

```bash
clifileconverter video1.mkv video2.mp4 video.mp4 --to webm
```

## Uninstall

For uninstalling run the command:

```bash
curl -fsSL https://raw.githubusercontent.com/almeidafm/clifileconverter/main/uninstall.sh | sh
```
## License

MIT License