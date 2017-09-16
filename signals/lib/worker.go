package lib

type WG struct{
	main chan func()
	allDone chan bool
}

func New(int n) {
	res := WG{
		main := make(chan func()),
		allDone := make(chan bool),
	}

	procDone := make(chan bool)

	for i := 0; i < n; i++{
		go func(){
			go func(){
//L'operatore <- sta a significare che f "ascolta" il canale res.main
				for{
					f := <- res.main
//Se il programma è stato eseguito
					if f == nil{
						procDone <- true
						return
					}
					f()
				}
			}
		}()
	}

	go func () {
		for i := 0; i < n; i++{
			_ = <- procDone
		}
		res.allDone <- true
	}()
}

func (wg WG) Add(f func()){
	wg.main <- f
}

func (wg WG) Wait(f func()){
	wait(wg.main)
	_ = <- wg.allDone
}