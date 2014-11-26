package log4g

import (
	"os"
	"testing"
	"time"
)

func Test_Noticef(t *testing.T) {
	logger := NewLogger("test_log/test_log", FilenameSuffixInSecond)
	for i := 0; i < 5; i++ {
		logger.Notice("test")
		time.Sleep(100 * time.Millisecond)
	}
}

func Test_CheckAndMkdir(t *testing.T) {
	logger := NewLogger("test_log/test_log", FilenameSuffixInHour)
	logger.checkAndMkdir("./test_log/test_log")
}

func TestMain(m *testing.T) {
	os.RemoveAll("test_log")
}
