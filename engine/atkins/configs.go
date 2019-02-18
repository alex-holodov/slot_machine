package atkins

const (
	TotalLines       = 32
	LinesAmount      = 3
	RelesCount       = 5
	PayLines         = 20
	BonusGames       = 10
	MinScalesToBonus = 3
	MiddleLine       = 1
)

type Symbol int

const (
	Atkins Symbol = iota
	Steak
	Ham
	Buffalo
	Sausage
	Eggs
	Butter
	Cheese
	Bacon
	Mayonnaise
	Scale
)

var payTable = map[Symbol][RelesCount]int64{
	Atkins:     {5000, 500, 50, 5},
	Steak:      {1000, 200, 40, 3},
	Ham:        {500, 150, 30, 2},
	Buffalo:    {300, 100, 25, 2},
	Sausage:    {200, 75, 20, 0},
	Eggs:       {200, 75, 20, 0},
	Butter:     {100, 50, 15, 0},
	Cheese:     {100, 50, 15, 0},
	Bacon:      {50, 25, 10, 0},
	Mayonnaise: {50, 25, 10, 0},
}

var stripsTable = [TotalLines][RelesCount]Symbol{
	{Scale, Mayonnaise, Ham, Ham, Bacon},
	{Mayonnaise, Buffalo, Butter, Cheese, Scale},
	{Ham, Steak, Eggs, Atkins, Steak},
	{Sausage, Sausage, Scale, Scale, Ham},
	{Bacon, Cheese, Cheese, Butter, Cheese},
	{Eggs, Mayonnaise, Mayonnaise, Bacon, Sausage},
	{Cheese, Ham, Butter, Cheese, Butter},
	{Mayonnaise, Butter, Ham, Sausage, Bacon},
	{Sausage, Bacon, Sausage, Steak, Buffalo},
	{Butter, Steak, Bacon, Eggs, Cheese},
	{Buffalo, Sausage, Steak, Bacon, Sausage},
	{Bacon, Mayonnaise, Buffalo, Mayonnaise, Ham},
	{Eggs, Ham, Butter, Sausage, Butter},
	{Mayonnaise, Atkins, Mayonnaise, Cheese, Steak},
	{Steak, Butter, Cheese, Butter, Mayonnaise},
	{Buffalo, Eggs, Sausage, Ham, Eggs},
	{Butter, Cheese, Eggs, Mayonnaise, Sausage},
	{Cheese, Bacon, Bacon, Bacon, Ham},
	{Eggs, Sausage, Mayonnaise, Buffalo, Atkins},
	{Atkins, Buffalo, Buffalo, Sausage, Butter},
	{Bacon, Scale, Ham, Cheese, Buffalo},
	{Mayonnaise, Mayonnaise, Sausage, Eggs, Mayonnaise},
	{Ham, Butter, Bacon, Butter, Eggs},
	{Cheese, Cheese, Cheese, Buffalo, Ham},
	{Eggs, Bacon, Eggs, Bacon, Bacon},
	{Scale, Eggs, Atkins, Mayonnaise, Butter},
	{Butter, Buffalo, Buffalo, Eggs, Steak},
	{Bacon, Mayonnaise, Bacon, Ham, Mayonnaise},
	{Sausage, Steak, Butter, Sausage, Sausage},
	{Buffalo, Ham, Cheese, Steak, Eggs},
	{Steak, Cheese, Mayonnaise, Mayonnaise, Cheese},
	{Butter, Bacon, Steak, Bacon, Buffalo},
}

var lineTable = [PayLines][RelesCount]int{
	{1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0},
	{2, 2, 2, 2, 2},
	{0, 1, 2, 1, 0},
	{2, 1, 0, 1, 2},
	{1, 0, 0, 0, 1},
	{1, 2, 2, 2, 1},
	{0, 0, 1, 2, 2},
	{2, 2, 1, 0, 0},
	{1, 0, 1, 2, 1},
	{1, 2, 1, 0, 1},
	{0, 1, 1, 1, 0},
	{2, 1, 1, 1, 2},
	{0, 1, 0, 1, 0},
	{2, 1, 2, 1, 2},
	{1, 1, 0, 1, 1},
	{1, 1, 2, 1, 1},
	{0, 0, 2, 0, 0},
	{2, 2, 0, 2, 2},
	{0, 2, 2, 2, 0},
}

var scalePayTable = map[int]int64{
	5: 100,
	4: 25,
	3: 5,
}
