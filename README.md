
# gortsplib

[![Build Status](https://travis-ci.org/aler9/gortsplib.svg?branch=master)](https://travis-ci.org/aler9/gortsplib)
[![Go Report Card](https://goreportcard.com/badge/github.com/mivtdole/gortsplib)](https://goreportcard.com/report/github.com/mivtdole/gortsplib)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/mivtdole/gortsplib?tab=doc)

RTSP 1.0 library for the Go programming language, written for [rtsp-simple-server](https://github.com/aler9/rtsp-simple-server).

Features:
* Provides primitives, a class for building clients (`ConnClient`) and a class for building servers (`ConnServer`)
* Supports TCP and UDP streaming protocols

## Examples

* [client-tcp](examples/client-tcp.go)
* [client-udp](examples/client-udp.go)

## Documentation

https://pkg.go.dev/github.com/mivtdole/gortsplib

## Links

Related projects
* https://github.com/aler9/rtsp-simple-server
* https://github.com/pion/sdp (SDP library used internally)
* https://github.com/pion/rtcp (RTCP library used internally)

IETF Standards
* RTSP 1.0 https://tools.ietf.org/html/rfc2326
* RTSP 2.0 https://tools.ietf.org/html/rfc7826
* HTTP 1.1 https://tools.ietf.org/html/rfc2616
