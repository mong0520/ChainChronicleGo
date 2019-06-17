package ccfsm

const (
	READY              = "ready"
	SELECT_ACTION      = "select_action"
	GACHA_SELECT_POOL  = "gacha_select_pool"
	GACHA_SELECT_COUNT = "gacha_select_count"
	GACHA_ACTION       = "gacha_action"
	TOWER_SELECT_ID    = "tower_select_id"
	TOWER_SELECT_MAX   = "tower_select_max"

	STATUS_READY = "status_ready"
)

// type State struct {
// 	PoolID int
// 	Count  int
// 	FSM    *fsm.FSM
// }

// func NewGacha() *State {
// 	state := &State{}

// 	state.FSM = fsm.NewFSM(
// 		READY,
// 		fsm.Events{
// 			{Name: SELECT_ACTION, Src: []string{READY}, Dst: GACHA_READY},
// 			{Name: GACHA_SELECT_POOL, Src: []string{GACHA_READY}, Dst: GACHA_SELECT_POOL},
// 			{Name: GACHA_SELECT_COUNT, Src: []string{GACHA_SELECT_POOL}, Dst: GACHA_SELECT_COUNT},
// 			{Name: GACHA_ACTION, Src: []string{GACHA_SELECT_COUNT}, Dst: GACHA_READY},
// 		},
// 		fsm.Callbacks{
// 			"enter_state": func(e *fsm.Event) { state.enterState(e) },
// 		},
// 	)

// 	return state
// }

// func (g *State) enterState(e *fsm.Event) {
// 	fmt.Printf("From %s enter state %s", e.Src, e.Dst)
// }
