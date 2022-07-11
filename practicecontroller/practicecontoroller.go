package practicecontroller

type Tuna struct {
	Name  string 'json : "name"'
    Price int    'json:"price"'
    Stock int    'json:"stock"'
}

// For Practice.

func get_hello() {

}

//THis function get "Hello World" json object.
func RequestHello(w http.ResponseWriter, r *http.Request) {
	hello_json := get_hello()
	return hello_json
}