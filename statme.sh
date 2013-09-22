#! /bin/bash
if [ $# -eq 0 ] then
	echo "ERROR: No arguments supplied"
else
	#PORT=2003	# Graphite
	#SERVER=graphite.domain.com
	
	PORT=8125	# Statsd
	SERVER=statsd.domain.com
	
	#echo "example.playing $1 `date +%s`"  | nc -u ${SERVER} ${PORT}; # to Graphite
	echo -n "example.does_it_aggregate:$1|c"  | nc -u ${SERVER} ${PORT}; # to StatsD
fi
