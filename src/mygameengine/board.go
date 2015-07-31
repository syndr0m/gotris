package mygameengine

type Board struct {
	frame     uint64
	onStart   func()
	onStop    func()
	onKeyDown func(key int)
	onKeyUp   func(key int)
	onRepaint func()
}

func (board *Board) OnStart(f func())          { board.onStart = f }
func (board *Board) OnStop(f func())           { board.onStop = f }
func (board *Board) OnKeyDown(f func(key int)) { board.onKeyDown = f }
func (board *Board) OnRepaint(f func())        { board.onRepaint = f }

func (board *Board) GetFrame() uint64 { return board.frame }

func (board *Board) Repaint() {
	board.frame++
	if board.onRepaint != nil {
		board.onRepaint()
	}
}
func (board *Board) Start() {
	board.frame = 0
	if board.onStart != nil {
		board.onStart()
	}
}
func (board *Board) Stop() {
	if board.onStop != nil {
		board.onStop()
	}
}
func (board *Board) KeyDown(key int) {
	if board.onKeyDown != nil {
		board.onKeyDown(key)
	}
}

func NewBoard() *Board {
	board := new(Board)
	return board
}
