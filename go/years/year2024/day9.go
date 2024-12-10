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

type StorageSpace struct {
	id    int
	index int
	size  int
}

func (d Day9) Part1() string {
	diskMap := d.Input[0]
	checksum := consolidateStorage(diskMap, true)
	return strconv.Itoa(checksum)
}

func (d Day9) Part2() string {
	diskMap := d.Input[0]
	checksum := consolidateStorage(diskMap, false)
	return strconv.Itoa(checksum)
}

func consolidateStorage(diskMap string, allowFragmentation bool) int {
	freeSpace := make([]StorageSpace, 0)
	files := make([]StorageSpace, 0)
	storage := make([]int, 0)

	currentStorageIndex := 0
	for i, ch := range diskMap {
		numBlocks, err := strconv.Atoi(string(ch))
		utils.PanicErr(err)

		storageSpace := StorageSpace{-1, currentStorageIndex, numBlocks}
		if i%2 == 0 {
			storageSpace.id = i / 2
			files = append(files, storageSpace)
		} else {
			freeSpace = append(freeSpace, storageSpace)
		}

		storage = fillRange(storage, storageSpace.id, currentStorageIndex, currentStorageIndex+numBlocks)
		currentStorageIndex += numBlocks
	}

	// if allowing fragmentation, we only continue after the file once the entire file has been re-allocated.
	// if not allowing fragmentation, we would continue if we couldn't find any available space.
	for i := len(files) - 1; i >= 0; {
		freeIndex := -1
		for j, space := range freeSpace {
			if space.size <= 0 || (!allowFragmentation && space.size < files[i].size) {
				continue
			}
			freeIndex = j
			break
		}

		if freeIndex == -1 || freeSpace[freeIndex].index > files[i].index {
			i--
			continue
		}

		count := 0
		for j := range freeSpace[freeIndex].size {
			if j >= files[i].size {
				break
			}
			storage[freeSpace[freeIndex].index+j] = files[i].id
			storage[files[i].index+files[i].size-1-j] = -1
			count++
		}
		files[i].size -= count
		freeSpace[freeIndex].size -= count
		freeSpace[freeIndex].index += count

		if files[i].size <= 0 {
			i--
		}
	}

	checksum := 0
	for i, num := range storage {
		if num == -1 {
			continue
		}
		checksum += i * num
	}
	return checksum
}

func fillRange[T any](slice []T, value T, start int, end int) []T {
	length := len(slice)
	for i := start; i < end; i++ {
		if i >= length {
			slice = append(slice, value)
			length++
			continue
		}
		slice[i] = value
	}
	return slice
}
