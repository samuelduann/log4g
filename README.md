log4g
=====

 Simple, Beautiful, and Time rotateable logging for Go. 


usage: 

logger := log4g.NewLogger("test_log/test_log", log4g.FilenameSuffixInHour)

logger.Noticef("test %d", 5)
