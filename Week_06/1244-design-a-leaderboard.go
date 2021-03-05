package Week_06

type Leaderboard struct {
	idToScore map[int]int
	scoreNumber map[int]int
}

func Constructor() Leaderboard {
	return Leaderboard{
		idToScore: map[int]int{},
		scoreNumber: map[int]int{},
	}
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
	if oldScore, has := this.idToScore[playerId]; has {
		this.idToScore[playerId] += score
		this.scoreNumber[oldScore]--
		this.scoreNumber[oldScore + score]++
	} else {
		this.idToScore[playerId] = score
		this.scoreNumber[score]++
	}
}


func (this *Leaderboard) Top(K int) int {
	scoreList := make([]int, 0, len(this.scoreNumber))
	for s, _ := range this.scoreNumber {
		scoreList = append(scoreList, s)
	}

	sort.Slice(scoreList, func (i, j int) bool {
		return scoreList[j] < scoreList[i]
	})

	count := 0
	sum := 0

	for _, maxScore := range scoreList {
		numbers := this.scoreNumber[maxScore]
		for ; numbers > 0 && count < K; {
			sum += maxScore

			numbers--
			count++
		}

		if count >= K {
			return sum
		}
	}

	return sum
}


func (this *Leaderboard) Reset(playerId int)  {
	s := this.idToScore[playerId]
	this.idToScore[playerId] = 0
	this.scoreNumber[s]--
	this.scoreNumber[0]++
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
