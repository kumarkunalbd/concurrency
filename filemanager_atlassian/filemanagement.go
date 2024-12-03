package filemanager_atlassian

type filemanager interface {
	AddFiles(data []FileData)
	GetTopFilesCollectionsFileSize(numfiles int) []int // collection Ids
	TotalSize() int
}

type FileData struct {
	filename     string
	filesize     int
	collectionId int
}

type FileManagerType1 struct {
	MaxHeap       []FileData
	collectionMap map[int][]FileData
}

func (fm *FileManagerType1) AddFile(data []FileData) {
	for _, file := range data {
		if valarr, present := fm.collectionMap[file.collectionId]; present {
			valarr = append(valarr, file)
		} else {
			newar := make([]FileData, 0)
			newar = append(newar, file)
			fm.collectionMap[file.collectionId] = newar
		}
	}
}

func (fm *FileManagerType1) GetTopFilesCollectionsFileSize(numfiles int) []int {
	mxheap := NewMaxHeap()
	for collectID, filedataArray := range fm.collectionMap {
		sumofFileSize := 0

		for _, filedata := range filedataArray {
			sumofFileSize += filedata.filesize
		}

		mxheap.Add(&Data{
			Content:  collectID,
			Priority: sumofFileSize,
		})
	}
	return getTopContentIdFromMaxHeap(mxheap, numfiles)
}

func getTopContentIdFromMaxHeap(mxheap *MaxHeap, num int) []int {
	topconetentid := make([]int, 0)
	for len(mxheap.DataList) > 0 && num > 0 {
		removearr := mxheap.Remove()
		curCollectionId := removearr.Content.(int)
		topconetentid = append(topconetentid, curCollectionId)
		num--
	}
	return topconetentid
}

// 1 := 5
// 2 := 3
// 3 :=  7

type MaxHeap struct {
	DataList    []*Data
	DataIndexes map[Data]int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		DataList:    make([]*Data, 0),
		DataIndexes: make(map[Data]int),
	}
}

type Data struct {
	Content  interface{}
	Priority int
}

func (mxh *MaxHeap) Add(aData *Data) {
	mxh.DataList = append(mxh.DataList, aData)
	curI := len(mxh.DataList) - 1
	mxh.DataIndexes[*aData] = curI

	for curI > 0 {
		curdata := mxh.DataList[curI]
		pI := (curI - 1) / 2
		pData := mxh.DataList[pI]

		if pData.Priority < curdata.Priority {
			mxh.DataList[curI], mxh.DataList[pI] = mxh.DataList[pI], mxh.DataList[curI]
			mxh.DataIndexes[*curdata] = pI
			mxh.DataIndexes[*pData] = curI
			curI = pI
			continue
		}
		break
	}
}

func (mxh *MaxHeap) RemoveAtIndex(index int) *Data {
	item := mxh.DataList[index]
	n := len(mxh.DataList)

	ndata := mxh.DataList[n-1]
	mxh.DataList[n-1], mxh.DataList[index] = mxh.DataList[index], mxh.DataList[n-1]
	mxh.DataIndexes[*ndata] = index
	delete(mxh.DataIndexes, *item)
	mxh.DataList = mxh.DataList[:(n - 1)]
	if index == n-1 {
		return item
	}
	curI := index

	for {
		maxI := curI
		curData := mxh.DataList[curI]
		maxIData := curData

		lc := 2*curI + 1
		if lc < len(mxh.DataList) && mxh.DataList[lc].Priority > maxIData.Priority {
			maxI = lc
			maxIData = mxh.DataList[lc]
		}
		rc := 2*curI + 2
		if rc < len(mxh.DataList) && mxh.DataList[rc].Priority > maxIData.Priority {
			maxI = rc
			maxIData = mxh.DataList[rc]
		}
		if maxI == curI {
			break
		}
		mxh.DataList[curI], mxh.DataList[maxI] = mxh.DataList[maxI], mxh.DataList[curI]
		mxh.DataIndexes[*maxIData] = curI
		mxh.DataIndexes[*curData] = maxI
		curI = maxI
	}

	return item
}

func (mxh *MaxHeap) Remove() *Data {
	return mxh.RemoveAtIndex(0)
}
