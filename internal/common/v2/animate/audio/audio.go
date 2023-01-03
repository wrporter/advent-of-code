package audio

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"io"
)

var (
	//go:embed neon-gaming.mp3
	NeonGaming_mp3 []byte
)

type Player struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

type audioStream interface {
	io.ReadSeeker
	Length() int64
}

func NewPlayer() (*Player, error) {
	audioContext := audio.NewContext(32000)
	var s audioStream

	var err error
	s, err = mp3.DecodeWithoutResampling(bytes.NewReader(NeonGaming_mp3))
	if err != nil {
		return nil, err
	}

	p, err := audioContext.NewPlayer(s)
	if err != nil {
		return nil, err
	}
	player := &Player{
		audioContext: audioContext,
		audioPlayer:  p,
	}

	player.audioPlayer.Play()

	return player, nil
}
