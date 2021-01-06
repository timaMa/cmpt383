#!/usr/bin/python3

from PyPDF2 import PdfFileMerger, PdfFileReader
import os

rst_file = 0
for file in os.listdir("./input"):
    if file.endswith(".rst"):
        rst_file = "./input/" + file

print("Read .rst file: " + rst_file)

os.system("rst2pdf " + rst_file + " ./output_pdfs/rst.pdf")

pdfs = ['./output_pdfs/md.pdf', './output_pdfs/rst.pdf']

merger = PdfFileMerger()

for pdf in pdfs:
	merger.append(PdfFileReader(open(pdf, 'rb')), import_bookmarks=False)

merger.write("merge.pdf")
print("Output the merge.pdf")
merger.close()