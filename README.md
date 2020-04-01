log4g
=====

 Simple, Beautiful, and Time rotateable logging for Go. 


usage: 

logger := NewLogger("test_log/test_log", FilenameSuffixInSecond)

logger.Noticef("test %d", 5)
