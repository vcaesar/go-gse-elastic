package nlp

func (s *Segmenter) InitCong(conf Conf) {
	s.conf = conf
	s.InitLoad()
}

func (s *Segmenter) InitLoad() (err error) {
	if s.conf.Dict != "" {
		err = s.Load(s.conf.Dict)
	} else {
		err = s.Load()
	}

	if err != nil {
		return
	}

	if s.conf.Stop != "" {
		err = s.LoadStop(s.conf.Stop)
	} else {
		err = s.LoadStop()
	}

	return
}
