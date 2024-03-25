The simplest flashcards version.

It stores map of words with translations and ask you to fill translations into the terminal. To stop enter "done".

No tests yet.

No refactoring yet. (bug when I put done from the start)

Other versions TBD:

- Parse the input file (json.Marshal / Unmarshal)
- Fix tests and add new tests for files
- Errors handing for file input
- Input file or data from browser
- Add DB instead of map
- Errors handing for DB
...
Think about other versions:

- Create common lib/functions for a few flashcards programs versions: dataValidation (can be JSON, text-file, manual input), Dictation with random func (?)
- Text-file input
- Provide an output file for added words
- Create DB backup
- Performance testing