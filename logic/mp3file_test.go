package logic

import "testing"

func TestParseMp3File(t *testing.T) {
	s, err := ParseMP3File("/Users/jimmy.jiang/doc/基督/灵命日粮/202509/mw250910.mp3")
	if err != nil {
		t.Errorf("ParseMP3File() error = %v", err)
	}
	t.Logf("ParseMP3File() = %v", s)
}
