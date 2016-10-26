package main

//TIKA
var tika_path_detect string = "http://127.0.0.1:9998/detect/stream"
var tika_path_meta string = "http://127.0.0.1:9998/meta"
var tika_path_meta_form string = "http://127.0.0.1:9998/meta/form"
var tika_path_meta_recursive string = "http://127.0.0.1:9998/rmeta/form/text"		//other options; xml/text/html

//STANFORD NER
var ner_extract string = "http://localhost:9000/?properties='annotators':'tokenize,ssplit,pos,ner','outputFormat':'json'"
