<div align="center">
<img src="./winres/icon.png" width="64" height="64">
<h1>cobalt-cli </h1>
<!-- Badges -->

[![GitHub Release](https://img.shields.io/github/v/release/princessmortix/cobalt?display_name=tag&style=for-the-badge&color=success)](#downloading)
[![Static Badge](https://img.shields.io/badge/cobalt_discord-join-blue?style=for-the-badge&logo=discord)](https://discord.gg/pQPt8HBUPu)
[![Static Badge](https://img.shields.io/badge/supported-services-0077b6?style=for-the-badge&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCA2NDAgNTEyIj48IS0tIUZvbnQgQXdlc29tZSBGcmVlIDYuNS4xIGJ5IEBmb250YXdlc29tZSAtIGh0dHBzOi8vZm9udGF3ZXNvbWUuY29tIExpY2Vuc2UgLSBodHRwczovL2ZvbnRhd2Vzb21lLmNvbS9saWNlbnNlL2ZyZWUgQ29weXJpZ2h0IDIwMjQgRm9udGljb25zLCBJbmMuLS0%2BPHBhdGggZD0iTTU3OS44IDI2Ny43YzU2LjUtNTYuNSA1Ni41LTE0OCAwLTIwNC41Yy01MC01MC0xMjguOC01Ni41LTE4Ni4zLTE1LjRsLTEuNiAxLjFjLTE0LjQgMTAuMy0xNy43IDMwLjMtNy40IDQ0LjZzMzAuMyAxNy43IDQ0LjYgNy40bDEuNi0xLjFjMzIuMS0yMi45IDc2LTE5LjMgMTAzLjggOC42YzMxLjUgMzEuNSAzMS41IDgyLjUgMCAxMTRMNDIyLjMgMzM0LjhjLTMxLjUgMzEuNS04Mi41IDMxLjUtMTE0IDBjLTI3LjktMjcuOS0zMS41LTcxLjgtOC42LTEwMy44bDEuMS0xLjZjMTAuMy0xNC40IDYuOS0zNC40LTcuNC00NC42cy0zNC40LTYuOS00NC42IDcuNGwtMS4xIDEuNkMyMDYuNSAyNTEuMiAyMTMgMzMwIDI2MyAzODBjNTYuNSA1Ni41IDE0OCA1Ni41IDIwNC41IDBMNTc5LjggMjY3Ljd6TTYwLjIgMjQ0LjNjLTU2LjUgNTYuNS01Ni41IDE0OCAwIDIwNC41YzUwIDUwIDEyOC44IDU2LjUgMTg2LjMgMTUuNGwxLjYtMS4xYzE0LjQtMTAuMyAxNy43LTMwLjMgNy40LTQ0LjZzLTMwLjMtMTcuNy00NC42LTcuNGwtMS42IDEuMWMtMzIuMSAyMi45LTc2IDE5LjMtMTAzLjgtOC42Qzc0IDM3MiA3NCAzMjEgMTA1LjUgMjg5LjVMMjE3LjcgMTc3LjJjMzEuNS0zMS41IDgyLjUtMzEuNSAxMTQgMGMyNy45IDI3LjkgMzEuNSA3MS44IDguNiAxMDMuOWwtMS4xIDEuNmMtMTAuMyAxNC40LTYuOSAzNC40IDcuNCA0NC42czM0LjQgNi45IDQ0LjYtNy40bDEuMS0xLjZDNDMzLjUgMjYwLjggNDI3IDE4MiAzNzcgMTMyYy01Ni41LTU2LjUtMTQ4LTU2LjUtMjA0LjUgMEw2MC4yIDI0NC4zeiIvPjwvc3ZnPg%3D%3D&logoColor=0077b6)](https://github.com/wukko/cobalt?tab=readme-ov-file#supported-services)
![GitHub License](https://img.shields.io/github/license/princessmortix/cobalt?style=for-the-badge&logo=unlicense)
</div>

Unofficial [cobalt](https://cobalt.tools) command line client made in go. cobalt-cli uses [gobalt library](https://github.com/princessmortix/gobalt) for communication between your machine <-> cobalt servers.

- [Features](#features)
- [Download](#downloading)
- [Roadmap](#roadmap)
- [Usage](#usage)
- [Compiling](#compiling)

## Features
- Get directly link from the service cdn (if possible);
- More than 15 services supported;
- JSON output using the flag `-j` or `--json`;
- Option to check status of cobalt servers;
- Use custom cobalt instances (see https://instances.hyper.lol);
- Get dubbed youtube audio.

## Download
| **Platform/OS** | **Download link**  |
|-----------------|--------------------|
| Windows         | [**x64**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-amd64.zip) / [x86](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-386.zip) / [arm](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-windows-arm.zip)    |
| Linux           | [**x64**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-amd64.tar.gz) / [x86](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-386.tar.gz) / [arm64](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-linux-arm64.tar.gz)    |
| MacOS           | [Intel](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-darwin-amd64.tar.gz) / [**M1, M2, M3**](https://github.com/princessmortix/cobalt/releases/latest/download/cobalt-darwin-arm64.tar.gz) |

You can also check the [releases page](https://github.com/princessmortix/cobalt/releases/latest) to download the latest version according to your platform.

## Roadmap
Planned features for cobalt-cli:

- [x] Json output;
  - [ ] Time expiration for the returned url.
- [ ] Option to save file to the current/custom folder, likely `-s` flag;
  - [ ] Display progress bar to track download progress (when supported by cobalt).
- [ ] Hability to use custom downloader program (wget, curl, got, etc);
- [ ] Translations.
 

## Usage
cobalt-cli has two subcommands:
- download: downloads something using cobalt
- instances: lists all known cobalt instances

### Download
```
usage: cobalt download [-h|--help] [-u|--url "<value>"] [-c|--video-codec
              (av1|vp9|h264)] [-q|--video-quality
              (144|240|360|480|720|1080|1440|2160)] [-f|--audio-format
              (opus|ogg|wav|mp3|best)] [-p|--filename-pattern
              (basic|pretty|nerdy|classic)] [-a|--no-video] [-V|--vimeo-dash]
              [-t|--full-tiktok-audio] [-v|--no-audio] [-d|--dubbed-audio]
              [-m|--metadata] [-g|--gif] [-j|--json] [-s|--status] [-i|--api
              "<value>"] [-l|--language "<value>"] [-b|--browser]

              download something using cobalt

Arguments:

  -h  --help               Print help information
  -u  --url                The url to download using cobalt
  -c  --video-codec        Video codec to be used. Applies only to youtube
                           downloads. AV1: 8K/HDR, lower support | VP9: 4K/HDR,
                           best quality | H264: 1080p, works everywhere.
                           Default: h264
  -q  --video-quality      Quality of the video. Default: 1080
  -f  --audio-format       Audio format/codec to be used. Using the default the
                           audio won't be re-encoded. Default: best
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
  -a  --no-video           Extract audio only. Default: false
  -V  --vimeo-dash         Downloads Vimeo videos using dash instead of
                           progressive. Default: false
  -t  --full-tiktok-audio  Enables download of original sound used in a tiktok
                           video. Default: false
  -v  --no-audio           Downloads only the video, without audio, when
                           possible. Default: false
  -d  --dubbed-audio       Downloads youtube audio dubbed, if present. Change
                           the language using -l <ISO 639-1 format>. Default:
                           false
  -m  --metadata           Disables file metadata. Default: false
  -g  --gif                Disables conversion of twitter gifs to a .gif file.
                           Default: true
  -j  --json               Output to stdin as json. Default: false
  -s  --status             Will only check status of the select cobalt server,
                           print and exit. All other options will be ignored,
                           except -j. Default: false
  -i  --api                Change the cobalt api endpoint to be used. See
                           others instances in https://instances.hyper.lol.
                           Default: https://co.wuk.sh
  -l  --language           Downloads dubbed youtube audio according to the
                           language set following the ISO 639-1 format. Only
                           takes effect if -d was passed as an argument.
                           Default: en
  -b  --browser            Opens the response link in default browser, if
                           successful. Default: false
```

### Instances
```
usage: cobalt instances [-h|--help] [-j|--json]

              get the list of cobalt instances

Arguments:

  -h  --help  Print help information
  -j  --json  Output to stdout as json
```

## JSON Output
See the documentation for the json output of cobalt-cli.
### Download
All json output from the download subcommands follows this format:
```json
{
  "error": bool,
  "message": "string",
  "urls": ["string1", "string2", ...]
}
```
Where:
| **name** | **type** | **info** | **example** |
|---|---|---|---|
| error | bool | true if something went wrong | "error":true |
| message | string | return error messages, otherwise "ok" | "message":"cobalt error: i couldn't connect to the service api. maybe it's down, or cobalt got blocked" |
| urls | []string | array of urls returned by the service, query scaped | "urls":["https%3A%2F%2Fus3-co.wuk.sh%2Fapi%2Fstream%3Ft%3D6kS3Xr97CAoqvPlBYX0r8%26e%3D1713849463113%26h%3DNHPfrLZ-BJejEnH2orowNy0zzTlVXSYw77RBhzIf0MU%26s%3DxvmTt9DTNl4wLslkfYfCUv6UDIPOTv9iZutl7ENM_dc%26i%3DuX7INLhsbzzofNxZaw6o7g"] |

### Instances
Returns almost the original json from [https://instances.hyper.lol/](https://instances.hyper.lol/instances.json), except we add two extra keys: error and message, just like above.

Example JSON:
```json
[{"error":false,"message":"success!"},[{"version":"7.12.6","commit":"50a98c8","branch":"current","name":"us3","url":"co.wuk.sh","cors":1,"startTime":"1713626380117","FrontendUrl":"cobalt.tools","ApiOnline":true,"FrontEndOnline":true},{"version":"7.12.6","commit":"50a98c8","branch":"current","name":"us-east","url":"cobalt.canine.tools","cors":1,"startTime":"1713837765475","FrontendUrl":"cobalt.canine.tools","ApiOnline":true,"FrontEndOnline":true},{"version":"7.12.6","commit":"50a98c8","branch":"current","name":"us-mw","url":"coapi.selfstacked.com","cors":1,"startTime":"1713626820678","FrontendUrl":"co.selfstacked.com","ApiOnline":true,"FrontEndOnline":true}]]
```

Error example:
```json
{"error":true,"message":"Get \"https://instances.hyper.lol/instances.json\": dial tcp: lookup instances.hyper.lol: no such host"}
```

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


# About & Thanks
- [cobalt](https://github.com/wukko/cobalt) made by [wukko](https://github.com/wukko) && [jj](https://github.com/dumbmoron), super cool people;
- [argparse](https://github.com/akamensky/argparse), for handling args;
- Icon made by [me](https://lostdusty.com.br);
- You, for using my application!
