// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package log has the interface for loggers in the sdk
package log

//Log interface to be implmented to be considered as a logger for web application
type Log interface {
	//Info logs the info
	Info(l ...interface{})
	//Debug logs debug specific information which will be skipped in production
	Debug(l ...interface{})
	//Warn logs warning info
	Warn(l ...interface{})
	//Error logs error info
	Error(l ...interface{})
	//Fatal logs the error and exits
	Fatal(l ...interface{})
}
