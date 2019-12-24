package npcs

type Npc interface {
	Interaction()
}

type Behaviour struct {
	Seen    SeenRules
	Patrol  PatrolRules
	Idle    IdleRules
	Special SpecialRules
}

// NPC Active moves (Has seen the player)
type SeenRules struct {
	Move bool
	
}

// After an NPC has seen and then subsequently lost the player
type PatrolRules struct {
}

// NPC Idle moves (Has not seen the player)
type IdleRules struct {
}

// Stub
type SpecialRules struct {
}
