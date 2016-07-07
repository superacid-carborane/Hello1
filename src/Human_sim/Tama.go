package Human_sim

//I don't think I designed this very well
type Character interface {
	Act()
}

type Feelings struct {
	//A bunch of bools sorted by good, bad, neutral, and other
}

type Bad_feels struct {
	//Feelings{all except bad feels false}
}

type Good_feels struct {
	//Feelings{all except good feels false}
}

type Neutral_feels struct {
	//Feelings{all except neutral feels false}
}

type Other_feels struct {
	//Feelings{all except other feels false}
}

/*type Disposition []Feelings (maybe slice of good etc feels?)
//Disposition is more constant, weights different emotions depending on what it is
*/

type Mood []Feelings

//changes from day to day as a result of actions

type Personality struct {
	intelligence int
	charisma     int
	energy       int
	laziness     int
	//Disposition
}

//once I figure this out I'm going to write what Act() does
