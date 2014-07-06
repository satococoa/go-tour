package main
import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (reader rot13Reader) Read(p []byte) (n int, err error) {
  n, err = reader.r.Read(p)
  for i := 0; i < len(p); i++ {
    if ('A' <= p[i] && p[i] < 'Z' - 12) || ('a' <= p[i] && p[i] < 'z' - 12) {
			p[i] += 13
		} else {
			p[i] -= 13
		}
  }
  return
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
