package settings

type group struct {
	Config     config
	Database   database
	Log        log
	TokenMaker tokenMaker
}

var Group = new(group)

func Inits() {
	Group.Config.Init()
	Group.Database.Init()
	Group.Log.Init()
	Group.TokenMaker.Init()
}
