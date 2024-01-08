#!/bin/sh

for DIR in build cardano submit sync watch
do
	cd $DIR && buf generate proto
	cd ..
done
