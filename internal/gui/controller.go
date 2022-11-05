package accenter

import (
	"math/rand"
	"time"
)

type guiController struct {
	a *guiApp
	m *guiModel
}

// Create a new controller, linked to the view and the model.
func NewController() *guiController {

	// initialize the random
	rand.Seed(time.Now().UnixNano())

	c := &guiController{}

	// create the view
	c.a = newApp(c)

	// initialize the model
	c.m = newModel()

	// create the UI, using placeholders everywhere
	c.a.buildUI()

	// update all the moving parts to match the current state:
	// the model has reasonable default values,
	// the view has only placeholders
	c.initAll()

	return c
}

// Run the app.
func (c *guiController) Run() {
	// run the app (will block)
	c.a.runApp()
}

func (c *guiController) initAll() {
	c.updateWord()
	c.updateGlossesInfo()
}

// --------------------------------------------------------------------------------
//  Reacts to events from UI (the view calls these funcs from the callbacks):
//  change the state of the model, then update the view.
// --------------------------------------------------------------------------------

// A keyboard button was clicked.
func (c *guiController) clicked(letter rune) {
	// fmt.Printf("C: Clicked '%c'\n", letter)
	// update the model
	c.m.clicked(letter)

	// check what the click did
	switch c.m.lastMistake {

	// was the right letter
	case ' ':
		c.a.kb.enableAll()
		c.updateWord()

	// all the word is correct
	case '!':
		// fmt.Printf("C: You won!\n")
		// disable all keyboard buttons, user will press next
		c.a.kb.disableAll()
		c.updateWord()

	// was the wrong letter
	default:
		c.a.kb.disable(c.m.lastMistake)

	}

}

// Clicked the button requesting a hint.
func (c *guiController) clickedHint() {
	c.m.clickedHint()
	c.updateWord()
}

// Clicked the button requesting the next word.
func (c *guiController) clickedNext() {
	// pick the next
	c.m.pickNewSecretWord()
	// enable all the keys
	c.a.kb.enableAll()
	// update the view elements
	c.updateWord()
	c.updateGlossesInfo()
}

// Clicked the button to mark a word as useless.
func (c *guiController) clickedUseless() {
	c.m.clickedUseless()
	c.clickedNext()
}

// --------------------------------------------------------------------------------
//  The model has changed:
//  the controller knows which view elements must be updated.
// --------------------------------------------------------------------------------

// The word to show has changed.
func (c *guiController) updateWord() {
	// get the word to show from the model
	// place it in the view
	c.a.updateWord(c.m.showWord)
}

// The word info to show has changed.
func (c *guiController) updateGlossesInfo() {
	// get the word to show from the model
	// place it in the view
	c.a.updateGlossesInfo(c.m.glossesInfo)
}
