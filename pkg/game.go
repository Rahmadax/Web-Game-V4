package pkg

type GameManager interface {
	Start()
	Stop()
}

type gameManager struct {
	engine interface{}
}

func NewGame() (GameManager, error){
	var game GameManager

	return game, nil
}

func (gameManager *gameManager) Start(){

}