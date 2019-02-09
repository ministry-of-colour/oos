all: 

setup: 
	npm install -g vue-cli
	vue init webpack-simple public
	vue init webpack-simple admin

run-public: 
	$(MAKE) -C public run

build:
	$(MAKE) -C api build