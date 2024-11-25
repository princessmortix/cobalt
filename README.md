<div align="center">
<img src="./winres/icon.png" width="64" height="64">
<h1>cobalt-cli </h1>
<!-- Badges -->

[![GitHub Release](https://img.shields.io/github/v/release/princessmortix/cobalt?display_name=tag&style=for-the-badge&color=success)](#downloading)
[![Static Badge](https://img.shields.io/badge/cobalt_discord-join-blue?style=for-the-badge&logo=discord)](https://discord.gg/pQPt8HBUPu)
[![Static Badge](https://img.shields.io/badge/supported-services-0077b6?style=for-the-badge&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCA2NDAgNTEyIj48IS0tIUZvbnQgQXdlc29tZSBGcmVlIDYuNS4xIGJ5IEBmb250YXdlc29tZSAtIGh0dHBzOi8vZm9udGF3ZXNvbWUuY29tIExpY2Vuc2UgLSBodHRwczovL2ZvbnRhd2Vzb21lLmNvbS9saWNlbnNlL2ZyZWUgQ29weXJpZ2h0IDIwMjQgRm9udGljb25zLCBJbmMuLS0%2BPHBhdGggZD0iTTU3OS44IDI2Ny43YzU2LjUtNTYuNSA1Ni41LTE0OCAwLTIwNC41Yy01MC01MC0xMjguOC01Ni41LTE4Ni4zLTE1LjRsLTEuNiAxLjFjLTE0LjQgMTAuMy0xNy43IDMwLjMtNy40IDQ0LjZzMzAuMyAxNy43IDQ0LjYgNy40bDEuNi0xLjFjMzIuMS0yMi45IDc2LTE5LjMgMTAzLjggOC42YzMxLjUgMzEuNSAzMS41IDgyLjUgMCAxMTRMNDIyLjMgMzM0LjhjLTMxLjUgMzEuNS04Mi41IDMxLjUtMTE0IDBjLTI3LjktMjcuOS0zMS41LTcxLjgtOC42LTEwMy44bDEuMS0xLjZjMTAuMy0xNC40IDYuOS0zNC40LTcuNC00NC42cy0zNC40LTYuOS00NC42IDcuNGwtMS4xIDEuNkMyMDYuNSAyNTEuMiAyMTMgMzMwIDI2MyAzODBjNTYuNSA1Ni41IDE0OCA1Ni41IDIwNC41IDBMNTc5LjggMjY3Ljd6TTYwLjIgMjQ0LjNjLTU2LjUgNTYuNS01Ni41IDE0OCAwIDIwNC41YzUwIDUwIDEyOC44IDU2LjUgMTg2LjMgMTUuNGwxLjYtMS4xYzE0LjQtMTAuMyAxNy43LTMwLjMgNy40LTQ0LjZzLTMwLjMtMTcuNy00NC42LTcuNGwtMS42IDEuMWMtMzIuMSAyMi45LTc2IDE5LjMtMTAzLjgtOC42Qzc0IDM3MiA3NCAzMjEgMTA1LjUgMjg5LjVMMjE3LjcgMTc3LjJjMzEuNS0zMS41IDgyLjUtMzEuNSAxMTQgMGMyNy45IDI3LjkgMzEuNSA3MS44IDguNiAxMDMuOWwtMS4xIDEuNmMtMTAuMyAxNC40LTYuOSAzNC40IDcuNCA0NC42czM0LjQgNi45IDQ0LjYtNy40bDEuMS0xLjZDNDMzLjUgMjYwLjggNDI3IDE4MiAzNzcgMTMyYy01Ni41LTU2LjUtMTQ4LTU2LjUtMjA0LjUgMEw2MC4yIDI0NC4zeiIvPjwvc3ZnPg%3D%3D&logoColor=0077b6)](https://github.com/wukko/cobalt?tab=readme-ov-file#supported-services)
![GitHub License](https://img.shields.io/github/license/princessmortix/cobalt?style=for-the-badge&logo=unlicense)
</div>

Unofficial [cobalt](https://cobalt.tools) command line client made in go. cobalt-cli uses [gobalt library](https://github.com/lostdusty/gobalt) for communication between your machine <-> cobalt servers.

- [Features](#features)
- [Download](#downloading)
- [Roadmap](#roadmap)
- [Usage](#usage)
- [Compiling](#compiling)

## Features
- Get directly link from the service cdn (if possible);
- More than 15 services supported;
- Option to check status of cobalt servers;
- Use custom cobalt instances (see https://instances.hyper.lol);
- Download the file directly to your computer.

## Download
| **Platform/OS** | **Download link**  |
|-----------------|--------------------|
| Windows         | [**x64**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-amd64.zip) / [x86](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-386.zip) / [arm](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-arm.zip)    |
| Linux           | [**x64**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-amd64.tar.gz) / [x86](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-386.tar.gz) / [arm64](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-arm64.tar.gz)    |
| Mac           | [Intel](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-darwin-amd64.tar.gz) / [**M1, M2, M3**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-darwin-arm64.tar.gz) |

You can also check the [releases page](https://github.com/princessmortix/cobalt/releases/latest) to download the latest version according to your platform.

Alternatively, if you have Go installed, you can use `go install github.com/lostdusty/cobalt@latest` to install to your machine.

## Roadmap
Planned features for cobalt-cli:

- [x] Option to save file to the current/custom folder, likely `-s` flag;
  - [x] Display progress bar to track download progress (when supported by cobalt).
  - [ ] Hability to use custom downloader program (wget, curl, got, etc);
- [ ] Translations.
- [X] Benchmarking.
 

## Usage
cobalt-cli is similar to yt-dlp, just use `cobalt [url]`. If you use `cobalt help`, it will just show the help message.

By default cobalt-cli saves the request link to the current directory, use the `-s` flag to change to another directory, like: `cobalt https://www.youtube.com/watch?v=n1a7o44WxNo -s ..\Videos`

### Help
```
usage: cobalt-cli [-h|--help] [url "<value>"] [-c|--video-codec (av1|vp9|h264)]
                  [-q|--video-quality (144|240|360|480|720|1080|1440|2160)]
                  [-f|--audio-format (opus|ogg|wav|mp3|best)]
                  [-Q|--audio-quality (64|128|192|256|320)]
                  [-p|--filename-pattern (basic|pretty|nerdy|classic)]
                  [-m|--mode (auto|audio|mute)] [-x|--proxy]
                  [-d|--disable-metadata] [-t|--tiktok-h265]
                  [-T|--tiktok-full-audio] [-g|--gif] [-s|--save "<value>"]
                  [-a|--api "<value>"] [-i|--instances] [-v|--verbose]
                  [-k|--key "<value>"] [-b|--benchmark] [-P|--print]

                  save what you want, directly from the terminal, no unwanted
                  distractions involved. powered by cobalt's api

Arguments:

  -h  --help               Print help information
      <url>                url to save
  -c  --video-codec        Video codec to be used. Applies only to youtube
                           downloads. AV1: 8K/HDR, lower support | VP9: 4K/HDR,
                           best quality | H264: 1080p, works everywhere.
                           Default: h264
  -q  --video-quality      Quality of the video, applies only to youtube
                           downloads. Default: 1080
  -f  --audio-format       Audio format/codec to be used. "best" doesn't
                           re-encodes audio. Default: best
  -Q  --audio-quality      Audio quality in kbps. Default: 320
  -p  --filename-pattern   File name pattern. Classic:
                           youtube_yPYZpwSpKmA_1920x1080_h264.mp4 | audio:
                           youtube_yPYZpwSpKmA_audio.mp3 // Basic: Video Title
                           (1080p, h264).mp4 | audio: Audio Title - Audio
                           Author.mp3 // Pretty: Video Title (1080p, h264,
                           youtube).mp4 | audio: Audio Title - Audio Author
                           (soundcloud).mp3 // Nerdy: Video Title (1080p, h264,
                           youtube, yPYZpwSpKmA).mp4 | audio: Audio Title -
                           Audio Author (soundcloud, 1242868615).mp3. Default:
                           pretty
  -m  --mode               Mode to download the video. Auto: video with audio |
                           Audio: only audio | Mute: video without audio.
                           Default: auto
  -x  --proxy              Tunnel the download through cobalt's servers,
                           bypassing potential restrictions and protecting your
                           identity and privacy. Default: false
  -d  --disable-metadata   Disable metadata in the downloaded file. Default:
                           false
  -t  --tiktok-h265        Use H265 codec for TikTok downloads. Default: false
  -T  --tiktok-full-audio  Download TikTok videos with the original sound used
                           in a TikTok video. Default: false
  -g  --gif                Convert Twitter videos to GIFs. Default: false
  -s  --save               What folder to save the file to. If not provided,
                           will use the current directory. Default:
                           D:\Docs\GitHub\cobalt
  -a  --api                Which API to use. Default is hyperdefined cobalt's
                           API. If you are hosting a custom API, or want to use
                           a different server, you can use it here. Default:
                           https://cobalt-backend.canine.tools
  -i  --instances          Show community instances and exit. Default: false
  -v  --verbose            Enable verbose logging. Default: false
  -k  --key                API key by the instance owner. You may need to
                           provide one to use download. Can be set with
                           COBALT_API_KEY environment variable. If not
                           provided, will load from keychain. Default:
  -b  --benchmark          Run a benchmark to test the download speed and
                           integrity. Default: false
  -P  --print              Print the download link only, do not download the
                           file. Default: false
```

### Instances
The command changed, now to view other instances, use `cobalt -i`


## Compiling
Make sure you have the lastest go compiler. [Download it here](https://go.dev/dl).

Easy as:
1. Clone this repository.
2. On the root of this repository, run `go mod tidy`. This will download this project dependencies.
3. To run the application, use `go run .`. To compile, run `go build`.

To add additional Windows metadata, you'll need:
- [go-winres](https://github.com/tc-hib/go-winres), follow the install instructions there

Then run `go-winres make` on the root of this repository, it will create two .syso files.

After that, building with `go build` will automatically embed these files on the Windows executable.

## Other projects
Check out too:
- [tobalt, cobalt in typescript](https://github.com/tskau/tobalt)
- [tcobalt, cobalt cli in rust](https://github.com/khyerdev/tcobalt)
- [pybalt, cobalt cli & api in python](https://github.com/nichind/pybalt)
- [gobalt(2), another lib for cobalt in go](https://github.com/andresperezl/gobalt)


# About & Thanks
- [cobalt](https://github.com/imputnet/cobalt) made by [wukko](https://github.com/wukko) && [jj](https://github.com/dumbmoron), cool people;
- [argparse](https://github.com/akamensky/argparse), for handling args;
- Icon made by [me](https://lostdusty.com.br);
- You, for using my application!
