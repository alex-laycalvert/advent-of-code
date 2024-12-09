package year2024

import (
	"strconv"

	"github.com/alex-laycalvert/advent-of-code/utils"
)

// ANSWERS:
//
// Part 1: 6346871685398
//
// Part 2: 6373055193464
type Day9 struct {
	Input []string
}

func (d Day9) Part1() string {
	answer := 0
	diskMap := d.Input[0]

	freeSpace := make([]int, 0)
	files := make([]int, 0)
	storage := make([]int, 0)
	currentStorageIndex := 0
	for i, ch := range diskMap {
		numBlocks, err := strconv.Atoi(string(ch))
		utils.PanicErr(err)
		// file
		if i%2 == 0 {
			for range numBlocks {
				storage = append(storage, i/2)
				files = append(files, currentStorageIndex)
				currentStorageIndex++
			}
			continue
		}

		// free space
		for range numBlocks {
			storage = append(storage, -1)
			freeSpace = append(freeSpace, currentStorageIndex)
			currentStorageIndex++
		}
	}

	freeSpaceIndex := 0
	filesIndex := len(files) - 1
	for freeSpaceIndex < len(freeSpace) && filesIndex >= 0 && freeSpace[freeSpaceIndex] < files[filesIndex] {
		storage[freeSpace[freeSpaceIndex]] = storage[files[filesIndex]]
		freeSpaceIndex++
		filesIndex--
	}

	for i := 0; i < len(files); i++ {
		answer += i * storage[i]
	}

	return strconv.Itoa(answer)
}

func (d Day9) Part2() string {
	answer := 0
	diskMap := d.Input[0]

	type StorageSpace struct {
		id    int
		index int
		size  int
	}

	freeSpace := make([]StorageSpace, 0)
	files := make([]StorageSpace, 0)
	storage := make([]int, 0)

	currentStorageIndex := 0
	for i, ch := range diskMap {
		numBlocks, err := strconv.Atoi(string(ch))
		utils.PanicErr(err)

		storageSpace := StorageSpace{index: currentStorageIndex, size: numBlocks}
		if i%2 == 0 {
			storageSpace.id = i / 2
			files = append(files, storageSpace)
		} else {
			storageSpace.id = -1
			freeSpace = append(freeSpace, storageSpace)
		}

		for j := currentStorageIndex; j < currentStorageIndex+numBlocks; j++ {
			storage = append(storage, storageSpace.id)
		}
		currentStorageIndex += numBlocks
	}

	for i := len(files) - 1; i >= 0; i-- {
		fileSpace := files[i]
		freeIndex := -1
		for j, space := range freeSpace {
			if space.size < fileSpace.size {
				continue
			}
			freeIndex = j
			break
		}

		if freeIndex == -1 {
			continue
		}

		if freeSpace[freeIndex].index > fileSpace.index {
			continue
		}

		for j := range fileSpace.size {
			storage[freeSpace[freeIndex].index+j] = fileSpace.id
			storage[fileSpace.index+j] = -1
		}
		freeSpace[freeIndex].size -= fileSpace.size
		freeSpace[freeIndex].index += fileSpace.size
	}

	for i, num := range storage {
		if num == -1 {
			continue
		}
		answer += i * num
	}

	return strconv.Itoa(answer)
}
