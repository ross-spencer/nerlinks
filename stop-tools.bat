@echo off

stop java -jar "tools/tika-server-1.13.jar"
stop java -mx1000m -cp "tools/stanford-corenlp-full-2015-12-09/*" edu.stanford.nlp.pipeline.StanfordCoreNLPServer

exit