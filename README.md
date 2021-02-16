# getcongif

I found this method of using .env variable more suitable , fast and protective.
unlike before accessing enviroment with `os` package at every file and location, but here with power of using multiple ENVIROMENT file with hand in dependencies injection is much more convinent and redable.

We can use different-different `.env` fie at the time of initializing object and then smoothly access the varable.

Directly use config.go and look into main.go file for how to work with this.

# Makefile
use command make run to run file and see the value SECRET_KEY which is "secret" from .env file.