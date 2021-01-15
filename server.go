package main

import (
	"blackjack/logic"
	pb "blackjack/proto"

	"google.golang.org/grpc"
)

//Server Grpc game server
type Server struct {
	pb.UnimplementedGameServer
	Game *logic.Game
}

func main() {

	print("hello")

	//init empty grpc server
	s := grpc.NewServer()

	//register it as a game server from proto file
	pb.RegisterGameServer(s, &Server{
		//init game logic
		Game: &logic.Game{
			Cards: make([]int, 14),
			Seats: make([]*logic.Seat, 5),
			PendingPlayers: &logic.RoomQueue{
				PendingPlayers: make([]string, 0),
			},
			Decks:        8,
			TurnTimeout:  10,
			State:        logic.Idle,
			PlayerCount:  0,
			HouseVisible: false,
			HouseScore:   0,
			HouseCards:   make([]int, 0),
		},
	})

}
