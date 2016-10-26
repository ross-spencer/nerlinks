@echo off

start java -jar "tools/tika-server-1.13.jar" --port=9998
start java -mx1000m -cp "tools/stanford-corenlp-full-2015-12-09/*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer

