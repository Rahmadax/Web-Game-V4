package gameInstance

type GameInstance interface {
	Start()
	Stop()
}

type gameInstance struct {

}

func NewGame() (GameInstance, error){
	var game GameInstance

	return game, nil
}

func AddPlayer(gameInstance, playerId string) (GameInstance, error){
	var game GameInstance

	return game, nil
}

func (gameInstance *gameInstance) Start(){

}

func (gameInstance *gameInstance) Stop(){

}