package bj

//RoomQueue takes care of getting people in the game when the round ends
type RoomQueue struct {
	PendingPlayers []string //ids
}

//Enqueue adds players to the queue
func (r *RoomQueue) Enqueue(playerID string) {

	r.PendingPlayers = append(r.PendingPlayers, playerID)

}

//Dequeue returns if queue empty and the playerID if it isn't
func (r *RoomQueue) Dequeue() (bool, string) {

	if len(r.PendingPlayers) == 0 {
		return true, ""
	}

	playerID := r.PendingPlayers[0]

	//remmoves first player from the list
	r.PendingPlayers = r.PendingPlayers[1:]

	return false, playerID
}

//SpaceTracker keeps track of free spaces at the table
type SpaceTracker struct {
	FreeSpots []int //index that are free to join
}

//PeekSpot returns a free spot if available, -1 if table full
func (s *SpaceTracker) PeekSpot() int {

	if len(s.FreeSpots) == 0 {
		return -1
	}

	return s.FreeSpots[0]
}

//PopSpot removes last free spot from list
func (s *SpaceTracker) PopSpot() {
	s.FreeSpots = s.FreeSpots[1:]
}
