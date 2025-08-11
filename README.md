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

- Renders the md into html and displays it

- Hitting enter or "l" will advance the slide (go to the next md file)

- Hitting shift + enter or "h" will go back a slide

- Hitting "r" will run the code snippet on a slide, if there is one and show the
  output

## Usage

I am thinking of something similar to this flow:

1. run the command `present` and pass the path of the directory that has the
   slides.
2. this should convert all the md slides into html files and then start a http
   fileserver 
3. while presenting all the keymaps above should work as intended
