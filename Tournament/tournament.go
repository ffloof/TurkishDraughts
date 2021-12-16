package tournament

func Run(){
	"TurkishDraughts/Board"
}

type AI interface {
	Play(board.BoardState) board.BoardState
	GetName() string
}

type mimmaxAI struct {
	name string
	table board.TransposTable
	//4 main settings in minmax.go
	ply int32
	advanced float32
	maxhash int32
	inaccuracy int32

}

type montecarloAI struct {
	name string
	sims int
}

func (mctsai montecarloAI) Play(currentBoard board.BoardState){
	return MCTS(currentBoard, mctsai.sims)
}

func (mmai minmaxAI) Play(currentBoard board.BoardState){
	plays := currentBoard.ValidPlays()

	board.MaxDepth = mmai.ply - 1
	board.MaximumHashDepth =  mmai.MaximumHashDepth
	board.TableDepthAllowedInaccuracy = mmai.inaccuracy
	board.AdvanceWeight = mmai.advanced

	var alpha float32 = 999.0 //TODO: set these according to team playing
	var beta float32 = -999.0
	var bestValue float32
	var bestOption board.BoardState = plays[0]

	_, next := currentBoard.MinMax(0, 999.0, -999.0, board.NewTable()) //TODO: recycle table and reduce depth by 1
}

func Run(){
	var AI1 AI 
	var AI2 AI

	b := board.CreateStartingBoard()
	for {
		options := b.ValidPlays()
		if len(options) == 1 {
			b = options[0]
		} else {
			if b.Turn == board.White {
				//Ai 1 plays
			} else {
				//Ai 2 plays
			}
		}
	}
}