package model

import (
	"encoding/json"
	"sort"

	"github.com/hashicorp/go-version"
)

type SortedVersions []string

func (sv *SortedVersions) UnmarshalJSON(data []byte) error {
	var v []string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*sv = SortVersions(v)
	return nil
}

func SortVersions(stringVersions []string) []string {
	semvers := []*version.Version{}

	for _, raw := range stringVersions {
		v, _ := version.NewVersion(raw)
		semvers = append(semvers, v)
	}

	sort.Sort(sort.Reverse(version.Collection(semvers)))

	sortedStrings := []string{}
	for _, s := range semvers {
		sortedStrings = append(sortedStrings, s.String())
	}

	return sortedStrings
}
