package example

import "github.com/codecov/example-go/awesome"

var result string

func Setup() {

    // Comment

    result = awesome.Smile()

}

func GetResult() string {

    /*
    Comment
    xxx
    */

    return result

}
