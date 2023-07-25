package main

import (
	"fmt"
	"strconv"
)

type FileInfo struct {
	Name string
	Size int
}

type CloudStorage struct {
	files map[string]FileInfo
}

func main() {
	queries := [][]string{
		{"ADD_FILE", "/file-a.txt", "4"},
		{"ADD_FILE", "/file-a.txt", "8"},
		{"ADD_FILE", "/dir-a/dir-c/file-b.txt", "11"},
		{"ADD_FILE", "/dir-a/dir-c/file-c.txt", "1"},
		{"ADD_FILE", "/dir-b/file-f.txt", "3"},
		{"GET_FILE_SIZE", "/file-a.txt"},
		{"GET_FILE_SIZE", "/file-c.txt"},
		{"MOVE_FILE", "/dir-b/file-f.txt", "/dir-b/file-e.txt"},
		{"MOVE_FILE", "/dir-b/file-a.txt", "/dir-b/file"},
		{"MOVE_FILE", "/file-a.txt", "/dir-b/file-e.txt"},
	}
	result := CloudStorageSimulator(queries)
	fmt.Println(result)
}

func (c *CloudStorage) addFile(name string, size int) string {
	if _, ok := c.files[name]; ok {
		c.files[name] = FileInfo{Name: name, Size: size}
		return "overwritten"
	} else {
		c.files[name] = FileInfo{Name: name, Size: size}
		return "created"
	}
}

func (c *CloudStorage) getFileSize(name string) string {
	if fileInfo, ok := c.files[name]; ok {
		return strconv.Itoa(fileInfo.Size)
	} else {
		return ""
	}
}

func (c *CloudStorage) moveFile(nameFrom string, nameTo string) string {
	if fileInfo, ok := c.files[nameFrom]; ok {
		if _, ok := c.files[nameTo]; ok {
			return "false"
		} else {
			c.files[nameTo] = fileInfo
			delete(c.files, nameFrom)
			return "true"
		}
	} else {
		return "false"
	}
}

func CloudStorageSimulator(queries [][]string) []string {
	storage := &CloudStorage{files: make(map[string]FileInfo)}
	result := []string{}
	for _, query := range queries {
		operation := query[0]
		switch operation {
		case "ADD_FILE":
			name := query[1]
			size, _ := strconv.Atoi(query[2])
			result = append(result, storage.addFile(name, size))
		case "GET_FILE_SIZE":
			name := query[1]
			result = append(result, storage.getFileSize(name))
		case "MOVE_FILE":
			nameFrom := query[1]
			nameTo := query[2]
			result = append(result, storage.moveFile(nameFrom, nameTo))
		}
	}
	return result
}
