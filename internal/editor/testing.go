package editor

import "io"

var _ Editor = (*ReverseEditor)(nil)

type ReverseEditor struct {
}

func (e *ReverseEditor) Edit(prefix, suffix string, contents io.Reader) ([]byte, error) {
	b, err := io.ReadAll(contents)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}

	return b, nil
}
