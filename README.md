contractor (for golang)
==========
Contractor basically allows you to dinamically manipulate struct. I have made this package because I could have different versions of the same struct (Just like of API Versions) . An newer version of the struct could have more or sell fields.



Getting started
==========
First of all install the package with go get  `"go get github.com/donseba/contractor"`
Add it to your imports where you want to use it. 

Structured Struct Layout 
==========
Basically we are going to store an reference to the struct in an `map[string]map[string]interface{}` ..say wut?

And it would look like something like this : 
```go
var TestContracts = map[string]map[string]interface{}{
	"01": map[string]interface{}{
		"Struct1": v1.Struct1{},
		"Struct2": v1.Struct2{},
	},
	"02": map[string]interface{}{
		"Struct1": v2.Struct1{},
		"Struct2": v2.Struct2{},
	},
}
```
As you can see, the first level contains the version of the struct. 

The second level is the most tricky parts. 

**keys** You always have to make sure all versions contain the same keys. Well it is not mandatory, but since it is so dinamic we have to set some groundrules.

**values** The values contain the struct reference.

Usage
======
```go
// Get the set of contracts which hold a map of structs assigned to version 01
contractSet := contractor.NewContract(models.TestContracts["01"])
```

```go
// Read the struct and asign it to an value in case we want to send.
Struct1, err := contractSet.Read("Struct1")
if err != nil {
  // Do your tipical error handling here
}

// Create some dummy data like Json.Unmarshal()  
dummyData := make(map[string]interface{})
dummyData["Field1"] = "valField1"
dummyData["Field2"] = "valField2"

// Assing the dummy data to the Struct.
Struct1.Set(dummyData)
```

The above example crearly only works if we have an valid `v01.struct1{}` Having 2 fields `Field1` & `Field2`
```go
package v01

type Struct1 struct {
	Field1        string
	Field2        string
}
```

Output
======
To get the output you can use the following : 
```go
//Read from the Struct (for now 2 ways to do this..)
fmt.Printf( "%+v\n", Struct1.Case ) // This will become unavailable over time.
fmt.Printf( "%+v\n", Struct1.Get() )
```
resulting in :
```console 
&{Field1:valField1 Field2:valField2}
```


Or you could convert it to json : 
```go
json_msg, _ := json.Marshal(Struct1.Get())
fmt.Fprintf(w, "%s", json_msg)
```
resulting in :
```json
{"Field1":"valField1","Field2":"valField2"}
```


