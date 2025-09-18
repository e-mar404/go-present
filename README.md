# Present

### MVP

- Takes a directory of markdown files to render
    - Files will be ordered with the following convention 
      (slide#_title_subtitle.md) with the subtitle being optional, so some 
      examples would be:
        - 001_Introduction_Functional-Programming.md
        - 002_Persistent-Data-Structures.md

    Note: Using underscores to divide the sections of slide number, title and
    subtitle let us use dashes to signify a space. Will also accept file names
    that already have spaces it is just personal that I don't like to put spaces
    on file names.

- Renders the md files into an alternate terminal

- Hitting ctrl-n will advance the slide (go to the next md file)

- Hitting ctrl-p will go back a slide

- Hitting ctrl-r will run the code snippet on a slide, if there is one and show the
  output

- There should be a cursor

## Usage

I am thinking of something similar to this flow:

1. use command `go-present /path/to/dir or /path/to/md/file`

2. while presenting the md will be formatted in a way that it shows the slides
   content and the speaker notes

3. using nst, prev, or run commands will modify what is shown

Note: should think about what editing a file should look like (maybe ctrl-e
should open the file in the default editor)

### Why not nvim plugin?

Started doing a neovim plugin and html server but that was pushing neovim too
much to what it isn't and to keep the unix philosophy "do one thing and do it
well".

Therefore a specific cli for presenting md files and not using another tool that
requires to write an abstraction over it.

## Todo's

✓ Get md files from a given path + tests
- create sliderenderer struct + tests
- create presentation struct + tests
