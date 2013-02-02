package vulpes

type Game interface {
	Children(Turn bool) []Game   // Determines the children nodes from this one. Must return an empty list
	Heuristic(Turn bool) float64 // Determines the heuristic score of the current node, from the perspective of the first player
	EndState(Turn bool) uint8    // 0 for tie, 1 for first player win, 2 for loss, 3 for not ended
}

// Just returns the score of a given state.
func Search(State Game, Depth uint32, Turn bool, Alpha, Beta, MinScore, MaxScore float64) float64 {
	End := State.EndState()
	if End == 1 {
		return MaxScore
	} else if End == 2 {
		return MinScore
	} else if End == 0 {
		return 0
	}
	if Depth <= 0 {
		return State.Heuristic(Turn)
	}
	if Turn {
		for _, Child := range State.Children(true) {
			TempScore := Search(Child, Depth-1, false, Alpha, Beta, MinScore, MaxScore)
			if TempScore >= Alpha {
				Alpha = TempScore
				if Beta <= Alpha {
					return Alpha
				}
			}
		}
		return Alpha
	}
	for _, Child := range State.Children(false) {
		TempScore := Search(Child, Depth-1, true, Alpha, Beta, MinScore, MaxScore)
		if TempScore <= Beta {
			Beta = TempScore
			if Beta <= Alpha {
				return Beta
			}
		}
	}
	return Beta
}

// Takes a starting node for the game, and returns the best child node and it's score
func SolveGame(State Game, Depth uint32, Turn bool, MinScore, MaxScore float64) (Game, float64) {
	End := State.EndState()
	if End == 1 {
		return State, MaxScore
	} else if End == 2 {
		return State, MinScore
	} else if End == 0 {
		return State, 0
	}
	if Depth <= 0 {
		return State, State.Heuristic(Turn)
	}
	Best := State
	if Turn {
		for _, Child := range State.Children(true) {
			TempScore := Search(Child, Depth-1, false, MinScore, MaxScore, MinScore, MaxScore)
			if TempScore > MinScore {
				Best = Child
				MinScore = TempScore
			}
		}
		return Best, MinScore
	}
	for _, Child := range State.Children(false) {
		TempScore := Search(Child, Depth-1, true, MinScore, MaxScore, MinScore, MaxScore)
		if TempScore < MaxScore {
			Best = Child
			MaxScore = TempScore
		}
	}
	return Best, MaxScore
}