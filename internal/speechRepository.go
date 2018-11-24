package internal

type SpeechRepository interface {
	TTS(string) ([]byte,error)
	PlaySound([]byte) error
}