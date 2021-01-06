# CMPT383 Project: PDFmerger between different file formats

This project can merge `.md` and `.rst` files into one `.pdf`. And it is very easy to use.

## Project idea
This project is a text processor with following workflow:
1. Read in a .md file inside `Ruby` and transfer to html.
2. Call `Go` to convert the html to a `.pdf` file via .html through the Cross-Language Communication method namely Foergin Function Langauge. 
3. Inside `Go`, call `Python` as an executable aiming to convert a `.rst` file to a .pdf file and then merge it with the `.pdf` file just generated in `Go`. 
The idea of my project is to provide a system to handle multi-format files. I apply `Go` as the language which handles the file processing mainly to accelerate the procedure.

## Languages
#### Programming language
`Go`: `Go` is used to pervide an efficient performance to do the task of text processing of format converting from `html` to `pdf` and call `Python` through `exec()`.

#### Script language
`Ruby`: `Ruby` is used to read a .md file which is only implemented in it and then call high-performance `Go` on that.

`Python`: Called by `Go` to merge two resulted .pdf files.  
`Javascript`: `Javascript` is used to implement the web for calculator and to perform simple calculation such as "ADD".


## Communication tools
1. *Run the other language's code as an executable*

This method is used to communicate between `Python` and `Go`. The `Python` part merge two result .pdf files.

2. *Foreign function interface*

FFI is used to communicate between `Ruby` and `Go`. This project uses `ctype` to call `Go` functions from `Ruby` to perform text processing and accelerate the process.


## Steps
```
vagrant up

cd project

ruby pdf_converter.rb

```

## Features
1. Users do not need to explicitly type the file names as the command arguments. The program can check the corresponding files in the `input` directory.

2. It is easy to use by putting `.md` and `.rst` in the `input` directory and run `ruby pdf_converter.rb`.

