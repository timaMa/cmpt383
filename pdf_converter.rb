require 'ffi'
require 'commonmarker'

module Pdf_convert
  extend FFI::Library
  ffi_lib './pdf_converter.so'

  attach_function :convert, [:string], :void
end


file = Dir['./input/*.md'].first
puts "Read the .md file: " + file
input_file = File.open(file)
html = CommonMarker.render_html(input_file.read, :DEFAULT)
puts "Call golang"
Pdf_convert.convert(html)