package datastructres

type MinHeap struct {
	DataList    []*Data
	DataIndexes map[Data]int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		DataList:    make([]*Data, 0),
		DataIndexes: make(map[Data]int),
	}
}

func (mxh *MinHeap) Add(aData *Data) {
	mxh.DataList = append(mxh.DataList, aData)
	curI := len(mxh.DataList) - 1
	mxh.DataIndexes[*aData] = curI

	for curI > 0 {
		curdata := mxh.DataList[curI]
		pI := (curI - 1) / 2
		pData := mxh.DataList[pI]

		if pData.Priority > curdata.Priority {
			mxh.DataList[curI], mxh.DataList[pI] = mxh.DataList[pI], mxh.DataList[curI]
			mxh.DataIndexes[*curdata] = pI
			mxh.DataIndexes[*pData] = curI
			curI = pI
			continue
		}
		break
	}
}

func (mxh *MinHeap) RemoveAtIndex(index int) *Data {
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
		minI := curI
		curData := mxh.DataList[curI]
		minIData := curData

		lc := 2*curI + 1
		if lc < len(mxh.DataList) && mxh.DataList[lc].Priority < minIData.Priority {
			minI = lc
			minIData = mxh.DataList[lc]
		}
		rc := 2*curI + 2
		if rc < len(mxh.DataList) && mxh.DataList[rc].Priority < minIData.Priority {
			minI = rc
			minIData = mxh.DataList[rc]
		}
		if minI == curI {
			break
		}
		mxh.DataList[curI], mxh.DataList[minI] = mxh.DataList[minI], mxh.DataList[curI]
		mxh.DataIndexes[*minIData] = curI
		mxh.DataIndexes[*curData] = minI
		curI = minI
	}

	return item
}

func (mxh *MinHeap) Remove() *Data {
	return mxh.RemoveAtIndex(0)
}
