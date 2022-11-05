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

	// update all the pieces of the view
	c.updateWord()

	if c.m.lastMistake == ' ' {
		// was the right letter
		c.a.kb.enableAll()
	} else if c.m.lastMistake == '!' {
		// all the word is correct
		// TODO
		// pick the next
		// enable all
	} else {
		// was the wrong letter
		c.a.kb.disable(c.m.lastMistake)

	}

	// obvs should not be done here
	// call c.a.kb.disable(letter)
	// c.a.kb.keys[letter].Disable()
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
