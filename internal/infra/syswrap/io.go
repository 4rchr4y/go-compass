package syswrap

import "io"

type IoClient struct{}

func (IoClient) Copy(dst io.Writer, src io.Reader) (int64, error) {
	return io.Copy(dst, src)
}
