package repository

import (
	"bytes"
	"context"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/renanberto/air-voice/internal"
	"io/ioutil"
	"time"

	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
)

type SpeechRepository struct{}

func NewSpeechRepository() internal.SpeechRepository {
	return &SpeechRepository{}
}

func (m *SpeechRepository) TTS(content string) ([]byte, error) {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: content},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "pt-PT",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}

func (m *SpeechRepository) PlaySound(sound []byte) error {
	s, format, err := wav.Decode(ioutil.NopCloser(bytes.NewReader(sound)))
	if err != nil {
		return err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	playing := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(playing)
	})))
	<-playing
	return nil
}
