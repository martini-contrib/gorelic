gorelic
=============

NewRelic middleware for martini framework.

# Installation
 - Run "go get github.com/martini-contrib/gorelic"
 - Add following lines to your main() function:
  gorelic.InitNewrelicAgent("7bceac019c7dcafae1ef95be3e3a3ff8866de266", "test martini app", true)
  m.Use(gorelic.Handler) //m - is martini application instance

Please, dont forget to change NewRelic license to your license
