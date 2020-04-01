log4g
=====

 Simple, Beautiful, and Time rotateable logging for Go. 


usage: 

import "github.com/samuelduann/log4g"

logger := log4g.NewLogger("test_log/test_log", log4g.FilenameSuffixInHour)

logger.Noticef("test %d", 5)
