package fallbackjokes

// Joke represents a fallback joke structure.
type Joke struct {
	Setup     string
	Punchline string
	Image     string
}

const FUNNY_IMAGE = "/static/funny-go.png"

// JokesList contains some fallback jokes.
var JokesList = []Joke{
	{
		Setup:     "Why was the JavaScript developer sad?",
		Punchline: "Because he didn't Node how to Express himself.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Go developer avoid the club?",
		Punchline: "Too much concurrency.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher stand still for a while?",
		Punchline: "He was garbage collecting his thoughts.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How do Go programmers say goodbye?",
		Punchline: "defer exit()",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher refuse the Python's offer to hang out?",
		Punchline: "He didn't want to get into any kind of deadlock.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why was the Go slice feeling down?",
		Punchline: "It got cut.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "What's a Gopher's favorite music?",
		Punchline: "Go-Go music.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How do Gophers communicate?",
		Punchline: "Through channels.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher always have a backup?",
		Punchline: "Because he didn't want to lose his Go routine.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher keep pausing?",
		Punchline: "He was stuck in a loop.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why was the Gopher so good at multitasking?",
		Punchline: "He was great at handling goroutines.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher bring a map to work?",
		Punchline: "To manage his key-value pairs.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher get a promotion?",
		Punchline: "He had no pointer errors.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher never panic?",
		Punchline: "Because he knew how to recover.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How did the Gopher flirt?",
		Punchline: "Are you a GC? Because my heart stops when I see you.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "What did the Gopher say after a tiring day?",
		Punchline: "Go-routines made me do it.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher always carry a notebook?",
		Punchline: "In case he encountered a panic, he could defer to his notes.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher stay calm in the storm?",
		Punchline: "He was used to handling Go-routines efficiently.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How did the Gopher order his coffee?",
		Punchline: "With no pointers, just a slice of sugar.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher prefer dark mode?",
		Punchline: "Because he loved shadowing variables.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "What's a Gopher's favorite time of the year?",
		Punchline: "Go-tober.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher always have a smile on his face?",
		Punchline: "He knew the value of a good interface.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How did the Gopher make his code dance?",
		Punchline: "With a Go-routine.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why did the Gopher always work in the garden?",
		Punchline: "He was good at planting Go-routines.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How did the Gopher greet his friend from far away?",
		Punchline: "Long time, no C!",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why was the Gopher so security conscious?",
		Punchline: "He knew a thing or two about Go-pher holes.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "Why was the Gopher never stressed at work?",
		Punchline: "He knew how to channel his energy.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "What did the Gopher wear to the party?",
		Punchline: "A Go-Tee.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "How do Gophers stay cool in summer?",
		Punchline: "They always have their Go-fans running.",
		Image:     FUNNY_IMAGE,
	},
	{
		Setup:     "What did the Gopher say to the pessimistic programmer?",
		Punchline: "Just Go with the flow.",
		Image:     FUNNY_IMAGE,
	},
}
