package events

func GetBalance(id int64)(balance string){

	return "1.2345"
}

func GetHelpMg(id int64)(help string) {

	return "This FAQ explores the how, what, where, and why of coffee. "+
		"It will explain the elements involved in making a great cup of brewed coffee " +
		"(espresso is a vast enough subject to deserve its own FAQ). The FAQ will be particularly " +
		"helpful for those who have little or no knowledge about coffee, but even more experienced " +
		"people should be able to glean new information."
}

func GetGreetingMsg() (greet string){

	return "Welcome press /start to continue.."
}
func GetWelcomeMsg() (welcome string) {
	return "You are welcome"
}

// value is a string path to the file, FileReader, or FileBytes
func GetPhoto(id int64) (value interface{}){

	return "./coffee.jpg"
}