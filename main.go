package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lostdusty/gobalt/v2"
	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tgoncuoglu/argparse"
)

var version = "2.0.1"
var useragent = fmt.Sprintf("cobalt-cli/%v (+https://github.com/lostdusty/cobalt; go/%v; %v/%v)", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)

func main() {
	cobaltParser := argparse.NewParser("cobalt-cli", "save what you want, directly from the terminal, no unwanted distractions involved. powered by cobalt's api")
	cobaltParser.ExitOnHelp(true)
	urlToDownload := cobaltParser.StringPositional("url", &argparse.Options{
		Required: false,
		Validate: func(args []string) error {
			if args[0] == "help" {
				return fmt.Errorf("\r%s", cobaltParser.Usage(nil))
			}
			if args[0] == "version" {
				return fmt.Errorf("\rcobalt-cli version %s\n%s", version, cobaltParser.Usage(nil))
			}
			_, err := url.Parse(args[0])
			if err != nil {
				return fmt.Errorf("invalid url, or parser failed to parse it: %s", err)
			}
			return nil
		},
		Help: "url to save",
	})

	youtubeVideoCodec := cobaltParser.Selector("c", "video-codec", []string{"av1", "vp9", "h264"}, &argparse.Options{
		Required: false,
		Help:     "Video codec to be used. Applies only to youtube downloads. AV1: 8K/HDR, lower support | VP9: 4K/HDR, best quality | H264: 1080p, works everywhere",
		Default:  "h264",
	})
	youtubeVideoQuality := cobaltParser.Selector("q", "video-quality", []string{"144", "240", "360", "480", "720", "1080", "1440", "2160"}, &argparse.Options{
		Required: false,
		Help:     "Quality of the video, applies only to youtube downloads",
		Default:  "1080",
	})
	audioCodec := cobaltParser.Selector("f", "audio-format", []string{"opus", "ogg", "wav", "mp3", "best"}, &argparse.Options{
		Required: false,
		Help:     "Audio format/codec to be used. \"best\" doesn't re-encodes audio",
		Default:  "best",
	})
	audioQuality := cobaltParser.Selector("Q", "audio-quality", []string{"64", "128", "192", "256", "320"}, &argparse.Options{
		Required: false,
		Help:     "Audio quality in kbps",
		Default:  "320",
	})
	fileNamePattern := cobaltParser.Selector("p", "filename-pattern", []string{"basic", "pretty", "nerdy", "classic"}, &argparse.Options{
		Required: false,
		Help:     "File name pattern. Classic: youtube_yPYZpwSpKmA_1920x1080_h264.mp4 | audio: youtube_yPYZpwSpKmA_audio.mp3 // Basic: Video Title (1080p, h264).mp4 | audio: Audio Title - Audio Author.mp3 // Pretty: Video Title (1080p, h264, youtube).mp4 | audio: Audio Title - Audio Author (soundcloud).mp3 // Nerdy: Video Title (1080p, h264, youtube, yPYZpwSpKmA).mp4 | audio: Audio Title - Audio Author (soundcloud, 1242868615).mp3",
		Default:  "pretty",
	})
	typeDownload := cobaltParser.Selector("m", "mode", []string{"auto", "audio", "mute"}, &argparse.Options{
		Required: false,
		Help:     "Mode to download the video. Auto: video with audio | Audio: only audio | Mute: video without audio",
		Default:  "auto",
	})
	proxyDownload := cobaltParser.Flag("x", "proxy", &argparse.Options{
		Required: false,
		Help:     "Tunnel the download through cobalt's servers, bypassing potential restrictions and protecting your identity and privacy",
		Default:  false,
	})
	disableMetadata := cobaltParser.Flag("d", "disable-metadata", &argparse.Options{
		Required: false,
		Help:     "Disable metadata in the downloaded file",
		Default:  false,
	})
	tikTokH265 := cobaltParser.Flag("t", "tiktok-h265", &argparse.Options{
		Required: false,
		Help:     "Use H265 codec for TikTok downloads",
		Default:  false,
	})
	tikTokFullAudio := cobaltParser.Flag("T", "tiktok-full-audio", &argparse.Options{
		Required: false,
		Help:     "Download TikTok videos with the original sound used in a TikTok video",
		Default:  false,
	})
	convertTwitterGif := cobaltParser.Flag("g", "gif", &argparse.Options{
		Required: false,
		Help:     "Convert Twitter videos to GIFs",
		Default:  false,
	})
	saveToDisk := cobaltParser.Flag("s", "save", &argparse.Options{
		Required: false,
		Help:     "Save the downloaded file to disk",
		Default:  true,
	})
	apiUrl := cobaltParser.String("a", "api", &argparse.Options{
		Required: false,
		Help:     "Which API to use. Default is hyperdefined cobalt's API. If you are hosting a custom API, or want to use a different server, you can use it here",
		Default:  gobalt.CobaltApi,
	})
	showCommunityInstances := cobaltParser.Flag("i", "instances", &argparse.Options{
		Required: false,
		Help:     "Show community instances and exit",
		Default:  false,
	})
	debugVerbose := cobaltParser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Enable verbose logging",
		Default:  false,
	})
	apiKey := cobaltParser.String("k", "key", &argparse.Options{
		Required: false,
		Help:     "API key by the instance owner. You may need to provide one to use download. Can be set with COBALT_API_KEY environment variable",
		Default:  gobalt.ApiKey,
	})
	flagBenchmark := cobaltParser.Flag("b", "benchmark", &argparse.Options{
		Required: false,
		Help:     "Run a benchmark to test the download speed and integrity",
		Default:  false,
	})

	err := cobaltParser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	if *debugVerbose {
		log.SetLevel(log.DebugLevel)
	}

	if len(os.Args) < 2 {
		log.Debug("No arguments provided, showing help")
		fmt.Println(cobaltParser.Usage(nil))
		return
	}

	if *showCommunityInstances {
		log.Debug("Flag to show community instances is set, showing instances")
		communityInstances()
		return
	}

	if *apiKey != "" {
		log.Debug("API key was provided via flag, setting it to gobalt")
		gobalt.ApiKey = *apiKey
		log.Debugf("Key from flag: %v | Key from Gobalt: %v | Key from COBALT_API_KEY: %v", *apiKey, gobalt.ApiKey, os.Getenv("COBALT_API_KEY"))
	}

	gobalt.CobaltApi = *apiUrl

	if *flagBenchmark {
		log.Debug("Flag to run benchmark is set, running benchmark")
		result, err := doBenchmark()
		if err != nil {
			log.Fatal(err)
		}
		mapBool := map[bool]string{true: "Yes!", false: "No :("}
		benchmarkTable := table.NewWriter()
		benchmarkTable.SetOutputMirror(os.Stdout)
		benchmarkTable.AppendHeader(table.Row{"Instance", "Time to download", "Download speed (KB/s)", "File size (KB)", "File hash matches?"})
		benchmarkTable.AppendRow(table.Row{result.Name, result.TimeToDownload, result.DownloadSpeed, result.FileSize, mapBool[result.HashMatches]})
		benchmarkTable.SetStyle(table.StyleLight)
		benchmarkTable.Render()
		return
	}

	newDownload := gobalt.CreateDefaultSettings()
	log.Debugf("Creating new cobalt download with default options: %v", newDownload)
	newDownload.Url = *urlToDownload
	switch *youtubeVideoCodec {
	case "av1":
		newDownload.YoutubeVideoFormat = gobalt.AV1
	case "vp9":
		newDownload.YoutubeVideoFormat = gobalt.VP9
	case "h264":
		newDownload.YoutubeVideoFormat = gobalt.H264
	default:
		newDownload.YoutubeVideoFormat = gobalt.H264
	}
	newDownload.VideoQuality, _ = strconv.Atoi(*youtubeVideoQuality)
	switch *audioCodec {
	case "opus":
		newDownload.AudioFormat = gobalt.Opus
	case "ogg":
		newDownload.AudioFormat = gobalt.Ogg
	case "wav":
		newDownload.AudioFormat = gobalt.Wav
	case "mp3":
		newDownload.AudioFormat = gobalt.MP3
	case "best":
		newDownload.AudioFormat = gobalt.Best
	default:
		newDownload.AudioFormat = gobalt.Best
	}
	newDownload.AudioBitrate, _ = strconv.Atoi(*audioQuality)
	switch *fileNamePattern {
	case "basic":
		newDownload.FilenameStyle = gobalt.Basic
	case "pretty":
		newDownload.FilenameStyle = gobalt.Pretty
	case "nerdy":
		newDownload.FilenameStyle = gobalt.Nerdy
	case "classic":
		newDownload.FilenameStyle = gobalt.Classic
	default:
		newDownload.FilenameStyle = gobalt.Pretty
	}
	switch *typeDownload {
	case "auto":
		newDownload.Mode = gobalt.Auto
	case "audio":
		newDownload.Mode = gobalt.Audio
	case "mute":
		newDownload.Mode = gobalt.Mute
	default:
		newDownload.Mode = gobalt.Auto
	}
	newDownload.Proxy = *proxyDownload
	newDownload.DisableMetadata = *disableMetadata
	newDownload.TikTokH265 = *tikTokH265
	newDownload.TikTokFullAudio = *tikTokFullAudio
	newDownload.TwitterConvertGif = *convertTwitterGif
	log.Debugf("Options changed to: %v", newDownload)

	err = fetchContent(newDownload, *saveToDisk)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func fetchContent(options gobalt.Settings, save bool) error {
	log.Debug("Fetching content now, save to disk: ", save)
	log.Info("Sending request to cobalt server...")
	media, err := gobalt.Run(options)
	if err != nil {
		return err
	}
	log.Debug("Cobalt replied our request with the following url: ", media.URL)
	fmt.Println(media.URL)
	if save {
		log.Info("Downloading the file to disk...")

		requestDownload, err := http.NewRequest("GET", media.URL, nil)
		requestDownload.Header.Set("User-Agent", useragent)
		log.Debug("Creating new request to download the file\nUser-Agent: ", useragent)
		if err != nil {
			return err
		}

		responseDownload, err := gobalt.Client.Do(requestDownload)
		log.Debug("Sending request to download the file using gobalt client")
		if err != nil {
			return err
		}
		defer responseDownload.Body.Close()

		log.Debug("Request ok, status code: ", responseDownload.StatusCode)

		isResponseHTML := strings.Contains(responseDownload.Header.Get("Content-Type"), "text/html")
		if responseDownload.StatusCode != http.StatusOK || isResponseHTML {
			if isResponseHTML {
				return fmt.Errorf("we got blocked trying to download the file, contact the instance owner if you think this is a mistake\nDetails: response is HTML (%s)", responseDownload.Header.Get("Content-Type"))
			}
			readBody, _ := io.ReadAll(responseDownload.Body)
			log.Debugf("got status %v while download the file.\nBody:\n%v", responseDownload.Status, string(readBody))
			return fmt.Errorf("error downloading the file: %s", responseDownload.Status)
		}

		f, err := os.OpenFile(media.Filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		bar := progressbar.DefaultBytes(
			responseDownload.ContentLength,
			"downloading "+media.Filename,
		)
		io.Copy(io.MultiWriter(f, bar), responseDownload.Body)
		f.Sync()
		fmt.Println()
		log.Info("File downloaded successfully!")
	}

	return nil
}

func communityInstances() {
	instances, err := gobalt.GetCobaltInstances()
	if err != nil {
		log.Fatal("Error fetching community instances:", err)
		return
	}
	instancesTable := table.NewWriter()
	instancesTable.SetOutputMirror(os.Stdout)
	instancesTable.AppendHeader(table.Row{"API", "Score", "Trust", "Version (commit)", "Turnstile"})
	for _, instance := range instances {
		instancesTable.AppendRow(table.Row{instance.API, fmt.Sprintf("%.0f%%", instance.Score), instance.Trust, fmt.Sprintf("%v (%v)", instance.Version, instance.Commit), instance.Turnstile})
	}
	instancesTable.SetStyle(table.StyleRounded)
	instancesTable.Render()
}

type Benchmark struct {
	Name           string        // Instance name
	TimeToDownload time.Duration // Time to download the file
	DownloadSpeed  int           // Download speed in KB/s
	FileSize       int           // File size in KB
	FileHash       string        // File hash in SHA256
	HashMatches    bool          // If the hash matches the known good hash
}

func doBenchmark() (*Benchmark, error) {
	//Know good hash: a092e6e57ff79077b5b3a6db97739cd925b462662bac82236f9de4227ac84757
	cobaltBench := &Benchmark{
		Name: gobalt.CobaltApi,
	}

	log.Info("Starting benchmark...")
	downloadBenchmark := gobalt.CreateDefaultSettings()
	downloadBenchmark.Url = "https://x.com/lostydust/status/1720929746987425821"
	downloadBenchmark.Proxy = true
	downloadBenchmark.VideoQuality = 1080

	log.Debug("Running benchmark with the following options: ", downloadBenchmark)
	log.Debugf("API: %s | Key: %s", gobalt.CobaltApi, gobalt.ApiKey)
	grabUrl, err := gobalt.Run(downloadBenchmark)
	if err != nil {
		return nil, err
	}
	log.Debug("Ok, got tunnel url: ", grabUrl.URL)
	requestDownload, err := http.NewRequest("GET", grabUrl.URL, nil)
	requestDownload.Header.Set("User-Agent", useragent)
	if err != nil {
		return nil, err
	}
	fileBuffer := bytes.NewBuffer(nil)
	log.Debug("Starting download now...")
	start := time.Now()
	responseDownload, err := gobalt.Client.Do(requestDownload)
	if err != nil {
		return nil, err
	}
	defer responseDownload.Body.Close()
	if responseDownload.StatusCode != http.StatusOK {
		err = fmt.Errorf("got http status %v while benchmarking the file", responseDownload.Status)
		return nil, err
	}
	_, err = io.Copy(fileBuffer, responseDownload.Body)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(start)
	log.Debug("Downloaded file in ", elapsed.Seconds(), " seconds")
	hashfile := sha256.New()
	hashfile.Write(fileBuffer.Bytes())
	cobaltBench.TimeToDownload = elapsed
	cobaltBench.FileSize = fileBuffer.Len() / 1024
	cobaltBench.DownloadSpeed = int(float64(cobaltBench.FileSize) / elapsed.Seconds())
	cobaltBench.FileHash = fmt.Sprintf("%x", hashfile.Sum(nil))
	cobaltBench.HashMatches = cobaltBench.FileHash == "a092e6e57ff79077b5b3a6db97739cd925b462662bac82236f9de4227ac84757"
	log.Debugf("File hash: %s", cobaltBench.FileHash)
	log.Debugf("Hash matches? %v", cobaltBench.FileHash == "a092e6e57ff79077b5b3a6db97739cd925b462662bac82236f9de4227ac84757")
	log.Info("[PASS] Benchmark finished!")
	return cobaltBench, nil
}
