package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/akamensky/argparse"
	iso6391 "github.com/emvi/iso-639-1"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lostdusty/gobalt"
	"github.com/mergestat/timediff"
)

func main() {
	flagParser := argparse.NewParser("cobalt", "save what you love directly from command-line, no bullshit involved.")
	doDownload := flagParser.NewCommand("download", "download something using cobalt")
	getCobaltInstances := flagParser.NewCommand("instances", "get the list of cobalt instances")

	//jsonCobaltInstances := getCobaltInstances.Flag("J", "json-output", &argparse.Options{Required: false, Help: "return output in JSON format"})
	doDownload.ExitOnHelp(true)

	URL := doDownload.String("u", "url", &argparse.Options{
		Required: false,
		Help:     "The url to download using cobalt",
	})

	optionVideoCodec := doDownload.Selector("c", "video-codec", []string{"av1", "vp9", "h264"}, &argparse.Options{
		Required: false,
		Help:     "Video codec to be used. Applies only to youtube downloads. AV1: 8K/HDR, lower support | VP9: 4K/HDR, best quality | H264: 1080p, works everywhere",
		Default:  "h264",
	})

	optionVideoQuality := doDownload.Selector("q", "video-quality", []string{"144", "240", "360", "480", "720", "1080", "1440", "2160"}, &argparse.Options{
		Required: false,
		Help:     "Quality of the video",
		Default:  "1080",
	})

	optionAudioFormat := doDownload.Selector("f", "audio-format", []string{"opus", "ogg", "wav", "mp3", "best"}, &argparse.Options{
		Required: false,
		Help:     "Audio format/codec to be used. Using the default the audio won't be re-encoded",
		Default:  "best",
	})

	optionFilenamePattern := doDownload.Selector("p", "filename-pattern", []string{"basic", "pretty", "nerdy", "classic"}, &argparse.Options{
		Required: false,
		Help:     "File name pattern. Classic: youtube_yPYZpwSpKmA_1920x1080_h264.mp4 | audio: youtube_yPYZpwSpKmA_audio.mp3 // Basic: Video Title (1080p, h264).mp4 | audio: Audio Title - Audio Author.mp3 // Pretty: Video Title (1080p, h264, youtube).mp4 | audio: Audio Title - Audio Author (soundcloud).mp3 // Nerdy: Video Title (1080p, h264, youtube, yPYZpwSpKmA).mp4 | audio: Audio Title - Audio Author (soundcloud, 1242868615).mp3",
		Default:  "pretty",
	})

	optionAudioOnly := doDownload.Flag("a", "no-video", &argparse.Options{
		Required: false,
		Help:     "Extract audio only",
		Default:  false,
	})

	optionVimeoDash := doDownload.Flag("V", "vimeo-dash", &argparse.Options{
		Required: false,
		Help:     "Downloads Vimeo videos using dash instead of progressive",
		Default:  false,
	})

	optionFullTikTokAudio := doDownload.Flag("t", "full-tiktok-audio", &argparse.Options{
		Required: false,
		Help:     "Enables download of original sound used in a tiktok video",
		Default:  false,
	})

	optionVideoOnly := doDownload.Flag("v", "no-audio", &argparse.Options{
		Required: false,
		Help:     "Downloads only the video, without audio, when possible",
		Default:  false,
	})

	optionDubAudio := doDownload.Flag("d", "dubbed-audio", &argparse.Options{
		Required: false,
		Help:     "Downloads youtube audio dubbed, if present. Change the language using -l <ISO 639-1 format>",
		Default:  false,
	})

	optionDisableMetadata := doDownload.Flag("m", "metadata", &argparse.Options{
		Required: false,
		Help:     "Disables file metadata",
		Default:  false,
	})

	optionConvertTwitterGif := doDownload.Flag("g", "gif", &argparse.Options{
		Required: false,
		Help:     "Disables conversion of twitter gifs to a .gif file",
		Default:  true,
	})

	var outputJson, jsonOutput = doDownload.Flag("j", "json", &argparse.Options{Required: false, Help: "Output to stdout as json"}), getCobaltInstances.Flag("j", "json", &argparse.Options{Required: false, Help: "Output to stdout as json"})

	commandStatus := doDownload.Flag("s", "status", &argparse.Options{
		Required: false,
		Help:     "Will only check status of the select cobalt server, print and exit. All other options will be ignored, except -j",
		Default:  false,
	})

	customCobaltApi := doDownload.String("i", "api", &argparse.Options{
		Required: false,
		Help:     "Change the cobalt api endpoint to be used. See others instances in https://instances.hyper.lol",
		Default:  gobalt.CobaltApi,
	})

	customLanguage := doDownload.String("l", "language", &argparse.Options{
		Required: false,
		Help:     "Downloads dubbed youtube audio according to the language set following the ISO 639-1 format. Only takes effect if -d was passed as an argument",
		Default:  gobalt.UserLanguage,
	})

	openInBrowser := doDownload.Flag("b", "browser", &argparse.Options{
		Required: false,
		Help:     "Opens the response link in default browser, if successful",
		Default:  false,
	})

	err := flagParser.Parse(os.Args)
	if err != nil && errors.Is(err, flagParser.Parse(os.Args)) {
		panic(fmt.Errorf("expected sub-command '%v' or '%v'", doDownload.GetName(), getCobaltInstances.GetName()))
	}

	if getCobaltInstances.Happened() {
		err := getInstances(*jsonOutput)
		if err != nil {
			if *jsonOutput {
				panic(errorJson(err))
			}
			panic(err)
		}
		return
	}

	if *commandStatus {
		err := checkStatus(*customCobaltApi, *outputJson)
		if err != nil {
			if *jsonOutput {
				panic(errorJson(err))
			}
			panic(err)
		}
		return
	}

	if *URL == "" {
		fmt.Println(flagParser.Help(doDownload))
		return
	}

	validateLanguage := iso6391.ValidCode(strings.ToLower(*customLanguage))
	if !validateLanguage {
		if *outputJson {
			panic(fmt.Errorf("invalid language code, check if the language code is following ISO 639-1 format"))
		}
		panic("Invalid language code: " + *customLanguage)
	}

	newSettings := gobalt.CreateDefaultSettings()
	if *customCobaltApi != gobalt.CobaltApi {
		gobalt.CobaltApi = *customCobaltApi
	}

	switch *optionAudioFormat {
	case "ogg":
		newSettings.AudioCodec = gobalt.Ogg
	case "wav":
		newSettings.AudioCodec = gobalt.Wav
	case "mp3":
		newSettings.AudioCodec = gobalt.MP3
	case "best":
		newSettings.AudioCodec = gobalt.Best
	case "opus":
		newSettings.AudioCodec = gobalt.Opus
	default:
		newSettings.AudioCodec = gobalt.Best
	}
	switch *optionVideoCodec {
	case "av1":
		newSettings.VideoCodec = gobalt.AV1
	case "h264":
		newSettings.VideoCodec = gobalt.H264
	case "vp9":
		newSettings.VideoCodec = gobalt.VP9
	default:
		newSettings.VideoCodec = gobalt.H264
	}
	switch *optionFilenamePattern {
	case "classic":
		newSettings.FilenamePattern = gobalt.Classic
	case "basic":
		newSettings.FilenamePattern = gobalt.Basic
	case "pretty":
		newSettings.FilenamePattern = gobalt.Pretty
	case "nerdy":
		newSettings.FilenamePattern = gobalt.Nerdy
	}
	newSettings.AudioOnly = *optionAudioOnly
	newSettings.ConvertTwitterGifs = *optionConvertTwitterGif
	newSettings.DisableVideoMetadata = *optionDisableMetadata
	newSettings.DubbedYoutubeAudio = *optionDubAudio
	newSettings.FullTikTokAudio = *optionFullTikTokAudio
	newSettings.UseVimeoDash = *optionVimeoDash
	newSettings.Url = *URL
	newSettings.VideoOnly = *optionVideoOnly
	quality, err := strconv.Atoi(*optionVideoQuality)
	if err != nil {
		if *outputJson {
			panic(fmt.Errorf("expected int on flag -q, got something else: %s", *optionVideoQuality))
		}
		panic(fmt.Errorf("expected int on flag -q, got something else: %s\nError details: %e", *optionVideoQuality, err))
	}
	newSettings.VideoQuality = quality

	cobaltRequest, err := gobalt.Run(newSettings)
	if err != nil {
		if *outputJson {
			panic(errorJson(err))
		}
		panic(err)
	}

	if *outputJson {
		safeUrlOutput := make([]string, 0)
		for _, u := range cobaltRequest.URLs {
			safe := url.QueryEscape(u)
			safeUrlOutput = append(safeUrlOutput, safe)
		}

		if cobaltRequest.Status == "picker" {
			unmarshalOutput := map[string]interface{}{"error": false, "message": cobaltRequest.Text, "urls": safeUrlOutput}
			output, _ := json.Marshal(unmarshalOutput)
			fmt.Println(string(output))
			return
		}

		unmarshalOutput := map[string]interface{}{"error": false, "message": cobaltRequest.Text, "urls": safeUrlOutput}
		output, _ := json.Marshal(unmarshalOutput)
		fmt.Println(string(output))
		return
	}

	if cobaltRequest.Status == "picker" {
		fmt.Println(cobaltRequest.URLs)
		return
	}

	if *openInBrowser {
		for _, urls := range cobaltRequest.URLs {
			err := openInDefaultBrowser(urls)
			if err != nil {
				fmt.Println("Failed to open URL on default browser:", err)
			}
		}
	}

	fmt.Println(cobaltRequest.URL)
}

func checkStatus(api string, returnJson bool) error {
	check, err := gobalt.CobaltServerInfo(api)
	if err != nil {
		if returnJson {
			return err
		}
		return fmt.Errorf("failed to contact cobalt server at %s due of the following error %e", api, err)
	}

	if returnJson {
		respJson := map[string]interface{}{"error": false,
			"message":   "contact was successful",
			"branch":    check.Branch,
			"commit":    check.Commit,
			"name":      check.Name,
			"startTime": check.StartTime,
			"url":       check.URL,
			"Version":   check.Version,
			"Cors":      fmt.Sprint(check.Cors),
		}
		outputJson, _ := json.Marshal(respJson)
		fmt.Println(string(outputJson))
		return nil
	}

	fmt.Printf("%s Status:\nBranch: %v\nCommit: %v\nName: %v\nStart time: %v (%v)\nURL: %v\nVersion: %v\nCors: %v", api, check.Branch, check.Commit, check.Name, time.UnixMilli(check.StartTime).Format(time.RFC1123), utilHumanTime(check.StartTime), check.URL, check.Version, check.Cors)
	return nil
}

func errorJson(err error) string {
	marshalThis := map[string]interface{}{"error": true,
		"message": fmt.Sprintf("%s", err),
		//"urls":    make([]string, 0),
	}
	errorInJson, _ := json.Marshal(marshalThis)
	return string(errorInJson)
}

func openInDefaultBrowser(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("start", url).Start()
	default:
		return exec.Command("xdg-open", url).Start()
	}
}

func getInstances(args bool) error {
	instances, err := gobalt.GetCobaltInstances()
	if err != nil {
		return err
	}

	if args {
		appendHeaders := map[string]interface{}{
			"error":   false,
			"message": "success!",
		}
		merged := make([]any, 0)
		merged = append(merged, appendHeaders)
		merged = append(merged, instances)
		outJson, _ := json.Marshal(merged)
		fmt.Println(string(outJson))
		return nil
	}

	instancesTable := table.NewWriter()
	instancesTable.SetOutputMirror(os.Stdout)
	instancesTable.AppendHeader(table.Row{"name", "api url (frontend)", "api online (frontend)", "cors", "since", "version (branch, commit)"})
	for _, v := range instances {
		instancesTable.AppendRow(table.Row{v.Name, fmt.Sprintf("%v (front: %v)", v.URL, v.FrontendUrl), fmt.Sprintf("%v (front: %v)", v.ApiOnline, v.FrontEndOnline), v.Cors, utilHumanTime(v.StartTime), fmt.Sprintf("%v (%v, %v)", v.Version, v.Branch, v.Commit)})
	}
	instancesTable.SetStyle(table.StyleRounded)
	instancesTable.Render()

	return nil
}

func utilHumanTime(unixTime int64) string {
	since := time.UnixMilli(unixTime)
	return timediff.TimeDiff(since)
}
