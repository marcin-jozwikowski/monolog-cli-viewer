build_and_run:
	go build -ldflags "-s -w"  mcv.go 
	touch var/log.log
	cat var/log.log | ./mcv

run_against_data:
	go run mcv.go -test

test:
	go test mcv_test.go

log:
	echo '{"message":"Example log of DEBUG level.","context":{"firewall_name":"main"},"level":100,"level_name":"DEBUG","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' > var/log.log
	echo '{"message":"Example log of INFO level.","context":{"firewall_name":"main"},"level":200,"level_name":"INFO","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of NOTICE level.","context":{"firewall_name":"main"},"level":300,"level_name":"NOTICE","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of WARNING level.","context":{"firewall_name":"main"},"level":400,"level_name":"WARNING","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of ERROR level.","context":{"firewall_name":"main"},"level":500,"level_name":"ERROR","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of CRITICAL level.","context":{"firewall_name":"main"},"level":600,"level_name":"CRITICAL","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of ALERT level.","context":{"firewall_name":"main"},"level":700,"level_name":"ALERT","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log
	echo '{"message":"Example log of EMERGENCY level.","context":{"firewall_name":"main"},"level":800,"level_name":"EMERGENCY","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}' >> var/log.log

record_gif:
	# https://www.terminalizer.com/install
	cd docs/
	npm install --local
	./node_modules/terminalizer/bin/app.js record demo.yml

render_gif:
	./node_modules/terminalizer/bin/app.js render demo.yml -o demo.gif