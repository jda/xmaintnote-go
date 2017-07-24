# xmaintnote for go
Go library for handling network [Maintenance Notification BCOP](https://github.com/maint-notification/maint-notification/) icals

[![GoDoc](https://godoc.org/github.com/jda/xmaintnote-go?status.svg)](https://godoc.org/github.com/jda/xmaintnote-go) 
[![Build Status](https://travis-ci.org/jda/xmaintnote-go.svg?branch=master)](https://travis-ci.org/jda/xmaintnote-go) 
[![Report Card](https://goreportcard.com/badge/github.com/jda/xmaintnote-go)](https://goreportcard.com/report/github.com/jda/xmaintnote-go) 

## TODO
1. Can iCalendar files contain more than one event?
   This has implications for parsing. E.g. Do we need separate
   MaintNote and MaintEvent? What if reader stream has multiple events?
