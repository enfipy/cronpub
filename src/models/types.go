package models

type FileType uint8

const (
	FileType_IMAGE FileType = iota
	FileType_VIDEO
	FileType_GIF
)
