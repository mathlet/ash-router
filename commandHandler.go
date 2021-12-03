package ashrouter

// A CommandHandler is just a function that handles command invokes
type CommandHandler func(*Context) error
