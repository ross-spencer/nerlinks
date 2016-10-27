# stanfordlinks

###Start Tika

Download from: [TIKA Server](https://tika.apache.org/download.html)

   #Run TIKA and specify a port to operate on
   start java -jar "tools/tika-server-1.13.jar" --port=9998

###Start Stanford NER

Download from: [Stanford NER](http://nlp.stanford.edu/software/stanford-corenlp-full-2015-12-09.zip)

    # Run the server using all jars in the current directory (e.g., the CoreNLP home directory)
    java -mx1000m -cp "*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer [port] [timeout]