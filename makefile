all:
	git add .
	git commit -m "update code `date +%dd-mm-YYYY`"
	git pull --rebase
	git push
test:
	go test .