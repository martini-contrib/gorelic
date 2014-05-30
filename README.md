GoRelic
=======

NewRelic middleware for martini framework.

## Usage

~~~ go
import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
)

func main(){
	m := martini.Classic()

	gorelic.InitNewrelicAgent("YOUR_NEWRELIC_LICENSE_KEY", "YOUR_APPLICATION_NAME", true)
	m.Use(gorelic.Handler)

	m.Run()
}
~~~

## Authors

* [Yuriy Vasiyarov](http://github.com/yvasiyarov)