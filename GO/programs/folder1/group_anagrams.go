// 49. Group Anagrams 
//
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		sSlice := strings.Split(s, "")
		sort.Strings(sSlice)
		key := strings.Join(sSlice,"")

		groups[key] = append(groups[key], s)

	}
	result := make([][]string,0,len(groups)) 
	for _, val := range groups {
		result = append(result, val)

	}
	return result 
}
