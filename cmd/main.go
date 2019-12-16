package cmd

import "../pkg"

func main(){

	gameManager, err := pkg.NewGame()
	if err != nil {
		return
	}

	gameManager.Start()


}