package nlp

import (
	"github.com/go-ego/gse"
)

func (s *Segmenter) Cut(str string, hmm ...bool) []string {
	if len(hmm) > 0 {
		return s.seg.Cut(str, hmm[0])
	}

	return s.seg.Slice(str)
}

func (s *Segmenter) CutSearch(str string, hmm ...bool) []string {
	return s.seg.CutSearch(str, hmm...)
}

func (s *Segmenter) CutAll(str string) []string {
	return s.seg.CutAll(str)
}

func (s *Segmenter) Pos(str string, hmm ...bool) (pos []gse.SegPos) {
	if len(hmm) > 0 {
		return s.pos.Cut(str, hmm[0])
	}

	return s.seg.Pos(str, false)
}

func (s *Segmenter) Trim(str []string) []string {
	return s.seg.Trim(str)
}

func (s *Segmenter) CutTrim(str string) []string {
	return s.Trim(s.Cut(str))
}

func (s *Segmenter) CutStr(str []string) string {
	return s.seg.CutStr(str)
}

func (s *Segmenter) Cuts(str string) string {
	return s.CutStr(s.Cut(str))
}

func (s *Segmenter) CutTrimStr(str string) string {
	return s.CutStr(s.CutTrim(str))
}

func (s *Segmenter) CutSearchTrimStr(str string) string {
	arr := s.Trim(s.CutSearch(str))
	return s.CutStr(arr)
}

func (s *Segmenter) CutSearchStr(str string) string {
	return s.CutStr(s.CutSearch(str))
}

func (s *Segmenter) CutHtml(str string) string {
	return s.seg.CutTrimHtmls(str)
}

func (s *Segmenter) CutUrl(str string, num ...bool) string {
	return s.seg.CutUrls(str, num...)
}
