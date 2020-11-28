#!/usr/bin/env bash

gnome-terminal -e 'java -mx1000m -jar tools/tika-app-1.24.1.jar --port 9998'  #> /dev/null 2>&1 &
gnome-terminal -e 'java -mx2000m -cp "tools/stanford-corenlp-4.1.0/*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer' #> /dev/null 2>&1 &
 
