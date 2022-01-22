package engine

// point list that will be able to tie animation to movement
// a -> b -> c with a, b, c being points and -> being transitions

//

type Waypoint struct {
	Position       Vec2
	TransitionTime float64
	TransitionMode int
}

type WaypointList struct {
	Waypoints []Waypoint
}

type WaypointSystem struct {
	Lists []*WaypointList
}

func CreateWaypointList() *WaypointList {
	return &WaypointList{}
}

func (w *WaypointList) AddWaypoint(position Vec2, transitionTime float64, transitionMode int) {
	w.Waypoints = append(w.Waypoints, Waypoint{position, transitionTime, transitionMode})
}

func (w *WaypointSystem) Update(deltaTime float64) {

}
