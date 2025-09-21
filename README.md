# Present

### MVP/TODO

- [x] Takes a directory of markdown files to render
    - [ ] Files will be ordered with the following convention 
    (slide#_title_subtitle.md) with the subtitle being optional, so some 
    examples would be:
        - 001_Introduction_Functional-Programming.md
        - 002_Persistent-Data-Structures.md

    Note: Using underscores to divide the sections of slide number, title and
    subtitle let us use dashes to signify a space. Will also accept file names
    that already have spaces it is just personal that I don't like to put spaces
    on file names.

- [x] Renders the md files into an alternate terminal
- [x] Hitting ctrl-n will advance the slide (go to the next md file)
- [x] Hitting ctrl-p will go back a slide
- [ ] Status bar detailing the dir name and/or file name
- [ ] There should be a cursor with visual highlighting
- [ ] Hitting ctrl-r will run the code snippet on a slide, if there is one and show the
  output

## Usage

I am thinking of something similar to this flow:

1. use command `go-present /path/to/dir`

2. while presenting the md will be formatted in a way that it shows the slides
   content and the speaker notes

3. using next, prev, or run commands will modify what is shown

Note: should think about what editing a file should look like (maybe ctrl-e
should open the file in the default editor)
