# nerlinks

Named entity recognition combining Tika's content extraction capabilities, with Stanford's NLP server, and Written in #Golang. 

###Design

All communication is done to server side tools via socket connections, rather than embedding libraries and other complex API
bits and pieces in the code. 

This frees us up to focus on development of the capability to combine different results from different services. 

Using Tika we get a large number of files handled for free so we don't have to worry too much about what files are sent to the server. Send them all! - We handle the exceptions as best as possible.

####Threading

The tool uses a default of eight goroutines/threads to communicate with Tika. If you increase your memory you may be able to up the number of threads. Configured in the nerlinks.go as 'throttle'.

Tika has its own bottleneck on memory used to extract plain text content and so this is the limiting factor for now. 

###TODO:

- Issues will be created as an when new features are recognised
- For example, creation of a synonym handler to combine results that we know (but the extrator) doesn't know to be identical. 
- Installation instructions
- Export functionality
- Database functionality (everything is done at runtime at present so there is no persistence across sessions)

##Examples

Some nice examples for users to take a look at and share. 

##Vogue:

[![asciicast](https://asciinema.org/a/90984.png)](https://asciinema.org/a/90984)

##Archival Collections: 

[![asciicast](https://asciinema.org/a/91272.png)](https://asciinema.org/a/91272)

##Instructions

Some basic instructions to begin. Will be fleshed out in the coming weeks. Video link to come. 

1) Start both services (start.bat, start-tools.sh), download links below. 
2) Run (cross platform)
    Linux: $./nerlinks -file <foldername>

    Windows: >nerlinks -file <foldername> 

3) Wait for the results to arrive.

###Start Tika

Download from: [TIKA Server](https://tika.apache.org/download.html)

   #Run TIKA and specify a port to operate on
   start java -jar "tools/tika-server-1.13.jar" --port=9998

###Start Stanford NER

Download from: [Stanford NER](http://nlp.stanford.edu/software/stanford-corenlp-full-2015-12-09.zip)

    # Run the server using all jars in the current directory (e.g., the CoreNLP home directory)
    java -mx1000m -cp "*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer [port] [timeout]
