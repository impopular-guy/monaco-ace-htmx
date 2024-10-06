# Monaco/Ace + HTMX

- Uses
    - [Go 1.23.0](https://go.dev/)
    - [htmx 2.0.2](https://htmx.org/)
    - [monaco-editor 0.52.0](https://microsoft.github.io/monaco-editor/)
    - [monaco-vim 0.4.1](https://github.com/brijeshb42/monaco-vim)
    - [ace 1.32.2](https://ace.c9.io/)
- To run:
    > `git clone ...` -> `cd repo` -> `go run main.go`

### Functionalities
- toggle vim mode
- toggle relative line numbers
- change language and update default code for that language
- preserve edited code when changing languages

### Problems
- Refreshing page will wipe out all edited code, this doesn't save anything to localStorage, etc.
- This version of `monaco-vim` suffers from https://github.com/brijeshb42/monaco-vim/issues/108 . However it doesn't effect functionality as long as you don't load monaco-vim during initial page load.
- The vim statusbar for ace editor is not that good, you might need to make your own status bar logic.

