#!/usr/bin/env bash

gnome-terminal -e 'java -mx1000m -jar tools/tika-server-1.13.jar --port=9998'  #> /dev/null 2>&1 &
gnome-terminal -e 'java -mx1000m -cp "tools/stanford-corenlp-full-2015-12-09/*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer' #> /dev/null 2>&1 &
 
