package filemanager_atlassian

import (
	"fmt"
	"testing"
)

func TestFileManager(t *testing.T) {
	inputfiledatat := []FileData{
		{filename: "ABC1",
			filesize:     5,
			collectionId: 2},
		{filename: "ABC2",
			filesize:     7,
			collectionId: 2},
		{filename: "ABC1",
			filesize:     10,
			collectionId: 1},
		{filename: "ABC1",
			filesize:     15,
			collectionId: 1},
		{filename: "ABC1",
			filesize:     5,
			collectionId: 2},
	}

	filemanagerty := FileManagerType1{
		MaxHeap:       make([]FileData, 0),
		collectionMap: make(map[int][]FileData),
	}
	filemanagerty.AddFile(inputfiledatat)
	fmt.Printf("Array=%v map=%v\n", filemanagerty.MaxHeap, filemanagerty.collectionMap)
	topcol := filemanagerty.GetTopFilesCollectionsFileSize(2)
	fmt.Printf("topCollection=%v\n", topcol)
}
