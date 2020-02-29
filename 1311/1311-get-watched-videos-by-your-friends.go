package main

// 1. build a list of people that are {level}-th friends with {id}
// 2. get a map of videos and each of its frequency
// 3. sort map by frequency
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	// 1. build a list of people that are {level}-th friends with {id}
	// i-th friends
	ithFriends := append([]int{}, friends[id]...)

	// key: ID, value: length of shortest path from {id}
	shortestPaths := map[int]int{}

	for _, f := range ithFriends {
		shortestPaths[f] = 1
	}

	for i := 2; i <= level; i++ {
		tmp := []int{}

		for _, f := range ithFriends {
			tmp = append(tmp, friends[f]...)
		}

		ithFriends = tmp
		tmp = []int{}

		// add to shortest path info if not already there
		for _, f := range ithFriends {
			if shortestPaths[f] == 0 {
				shortestPaths[f] = i
			}
		}
	}
	// shortest path to self is 0, so remove
	delete(shortestPaths, id)

	//fmt.Println("shortestPaths:", shortestPaths)

	levelthFriends := make(map[int]struct{})
	for id, path := range shortestPaths {
		if path == level {
			levelthFriends[id] = struct{}{}
		}
	}

	//fmt.Println("levelthFriends:", levelthFriends)

	// 2. get a map of videos and each of its frequency
	videoCount := map[string]int{}

	for id := range levelthFriends {
		for _, v := range watchedVideos[id] {
			videoCount[v] += 1
		}
	}

	//fmt.Println("videoCount:", videoCount)

	// 3. sort map by frequency
	res := getSortedVideos(videoCount)

	//fmt.Println("res:", res)

	return res
}

// getSortedVideos receives a map of video name and its count, and
// returns the name of video sorted by:
// 1. frequency
// 2. if same frequency, alphabetically
func getSortedVideos(m map[string]int) []string {
	frq := 0
	res := []string{}

	// 1. sort by frequency
	for len(res) < len(m) {
		list := []string{}

		for k, v := range m {
			if v == frq {
				list = append(list, k)
			}
		}

		// 2. if same count, sort alphabetically
		sort.Strings(list)
		res = append(res, list...)
		frq++
	}

	return res
}
