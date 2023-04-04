package main

func main() {

}

type storageSlice struct {
	data []int
}

func (s *storageSlice) find(id int) (int, bool) {
	for _, el := range s.data {
		if id == el {
			return el, true
		}
	}

	return 0, false
}

type storageMap struct {
	data map[int]struct{}
}

func (s *storageMap) find(id int) (int, bool) {
	if _, ok := s.data[id]; ok {
		return id, true
	}

	return 0, false
}
