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

- Hitting ctrl-n will advance the slide (go to the next md file)

- Hitting ctrl-p will go back a slide

- Hitting leader-rc will run the code snippet on a slide, if there is one and show the
  output

## Usage

I am thinking of something similar to this flow:

1. open a directory on vim and start presenting with a command / toggle

2. the plugin will start a go server that takes the md files that the lua client
   sends and will convert them to html

3. the go server will serve those files, listen for updates of the lua client,
   and open the page (prob some localhost:42069 like address)

4. while presenting the neovim client should have speaker notes and should be
   able to run code that will and show the output on both neovim and on the
   website

5. pressing next and prev on neovim should change the slide on both neovim and
   the website (this means that what ever page the neovim instance is showing
   should be mirrored on the website)

Note: editing text on the md file in the neovim instance should be fairly quick to
   change on the website as well
