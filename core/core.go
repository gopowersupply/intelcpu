package core

// List - Array of cores, some actions can be applied to all cores at the same time
type List []*Core

// Core - Core presentation
type Core struct {
	Path string
	Num  uint16
}