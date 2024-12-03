package mostpopular

import datastructres "ConcurrencyP/datastructures"

type MostPopular interface {
	IncreasePopularity(contentId int) int
	MostPopular() int
	DecreasePopularity(contentId int) int
}

type ContentPopularityManager struct {
	mxHeap   *datastructres.MaxHeap
	contentP map[int]int
}

func NewContentPopularityManager() *ContentPopularityManager {
	return &ContentPopularityManager{
		mxHeap:   datastructres.NewMaxHeap(),
		contentP: make(map[int]int),
	}
}

func (cm *ContentPopularityManager) IncreasePopularity(contentId int) int {
	if val, present := cm.contentP[contentId]; present {
		indexOfContent := cm.mxHeap.DataIndexes[datastructres.Data{
			Content:  contentId,
			Priority: val,
		}]
		cm.mxHeap.RemoveAtIndex(indexOfContent)
		cm.mxHeap.Add(&datastructres.Data{
			Content:  contentId,
			Priority: val + 1,
		})
		cm.contentP[contentId] = val + 1
	} else {
		cm.mxHeap.Add(&datastructres.Data{
			Content:  contentId,
			Priority: 1,
		})
		cm.contentP[contentId] = val + 1
	}
	return cm.contentP[contentId]
}

func (cm *ContentPopularityManager) MostPopular() int {
	if len(cm.mxHeap.DataList) == 0 {
		return -1
	}
	return cm.mxHeap.DataList[0].Content.(int)
}
func (cm *ContentPopularityManager) DecreasePopularity(contentId int) int {
	if val, present := cm.contentP[contentId]; present {
		indexOfContent := cm.mxHeap.DataIndexes[datastructres.Data{
			Content:  contentId,
			Priority: val,
		}]
		cm.mxHeap.RemoveAtIndex(indexOfContent)
		if val > 1 {
			cm.mxHeap.Add(&datastructres.Data{
				Content:  contentId,
				Priority: val - 1,
			})
			cm.contentP[contentId] = val - 1
		} else {
			delete(cm.contentP, contentId)
		}
		return cm.contentP[contentId]
	}
	return -1
}
