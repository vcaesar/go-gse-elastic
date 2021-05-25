package nlp

import (
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

type Conf struct {
	Dict string `toml:"dict"`
	Stop string `toml:"stop"`
}

type Segmenter struct {
	seg gse.Segmenter
	pos pos.Segmenter

	conf Conf
}

func (s *Segmenter) Load(files ...string) error {
	return s.seg.LoadDict(files...)
}

func (s *Segmenter) Dict() *gse.Dictionary {
	return s.seg.Dict
}

func (s *Segmenter) WithGse() {
	s.pos.WithGse(s.seg)
}

func (s *Segmenter) AddToken(text string, freq float64, pos ...string) {
	s.seg.AddToken(text, freq, pos...)
}

func (s *Segmenter) RemoveToken(text string) {
	s.seg.RemoveToken(text)
}

func (s *Segmenter) LoadStop(files ...string) error {
	return s.seg.LoadStop(files...)
}

func (s *Segmenter) AddStop(text string) {
	s.seg.AddStop(text)
}

func (s *Segmenter) IsStop(text string) bool {
	return s.seg.IsStop(text)
}
