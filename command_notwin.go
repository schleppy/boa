//go:build !windows
// +build !windows

package boa

var preExecHookFn func(*Command)
