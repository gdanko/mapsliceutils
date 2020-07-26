package mapsliceutils

import (
	"github.com/thoas/go-funk"
	"gopkg.in/yaml.v2"
)

func mapSliceItems(slice *yaml.MapSlice) map[string]yaml.MapItem {
	items := make(map[string]yaml.MapItem)
	for _, item := range *slice {
		strKey := item.Key.(string)
		items[strKey] = item
	}
	return items
}

func mapSliceDiff(sliceOne *yaml.MapSlice, sliceTwo *yaml.MapSlice) map[string]yaml.MapItem {
	var sliceOneKeys, sliceTwoKeys, intersection []string
	sliceDiff := make(map[string]yaml.MapItem)
	for _, item := range *sliceOne {
		sliceOneKeys = append(sliceOneKeys, item.Key.(string))
	}
	for _, item := range *sliceTwo {
		sliceTwoKeys = append(sliceTwoKeys, item.Key.(string))
	}
	for _, x := range sliceOneKeys {
		if !funk.Contains(sliceTwoKeys, x) {
			intersection = append(intersection, x)
		}
	}
	for _, item := range *sliceOne {
		strKey := item.Key.(string)
		if funk.Contains(intersection, item.Key.(string)) {
			sliceDiff[strKey] = item
		}
	}
	return sliceDiff
}

func mergeMapSlices(destSlice *yaml.MapSlice, sourceSlice *yaml.MapSlice, mergedSlice *yaml.MapSlice) {
	sourceItems := mapSliceItems(sourceSlice)
	sourceOnlyItems := mapSliceDiff(sourceSlice, destSlice)

	for _, destItem := range *destSlice {
		destKey := destItem.Key.(string)
		destValue := destItem.Value

		if funk.Contains(sourceItems, destKey) {
			sourceItem := sourceItems[destKey]
			sourceValue := sourceItem.Value
			nestedDestSlice, destIsNested := destValue.(yaml.MapSlice)
			nestedSourceSlice, sourceIsNested := sourceValue.(yaml.MapSlice)

			if sourceIsNested && destIsNested {
				nestedMap := &yaml.MapSlice{}
				nested := yaml.MapItem{Key: destKey, Value: nestedMap}
				(*mergedSlice) = append((*mergedSlice), nested)
				mergeMapSlices(&nestedDestSlice, &nestedSourceSlice, nestedMap)
			} else {
				*(mergedSlice) = append((*mergedSlice), sourceItem)
			}
		} else {
			*(mergedSlice) = append((*mergedSlice), destItem)
		}
	}
	for _, sourceOnlyItem := range sourceOnlyItems {
		*(mergedSlice) = append((*mergedSlice), sourceOnlyItem)
	}
}

func DeepMerge(destSlice *yaml.MapSlice, sourceSlice *yaml.MapSlice) (*yaml.MapSlice) {
	mergedSlice := &yaml.MapSlice{}
	mergeMapSlices(destSlice, sourceSlice, mergedSlice)
	return mergedSlice
}
