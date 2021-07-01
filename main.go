package main

import "fmt"

//"context"
//"encoding/json"
//"fmt"
//"log"
//"os"

//"github.com/jameskeane/bcrypt"
//"github.com/machinebox/graphql"

func main() {

	/*salt, csrf, cookie := getSalt() //salt in bytes
	s := readSalt(salt)             //salt in struct

	pass := os.Getenv("BACKEND_PASSWORD")
	//pass := "addasdasasd"
	// hash and verify a password with a static salt
	hash, _ := bcrypt.Hash(pass, s.Salt)

	if bcrypt.Match(pass, hash) {
		fmt.Println("They match")
	}

	client := graphql.NewClient("https://api.sorare.com/graphql")

	// make a request
	/*req := graphql.NewRequest(`{
	    	card (slug:"cristiano-ronaldo-dos-santos-aveiro-2020-rare-10" ) {
	        	 	id
	            	name
	            	onSale
					age
	        	}
	    	}
		`)*/

	/*reqCR := graphql.NewRequest(`
		query ($value: String!) {
			card (slug:$value ) {
				id
			    name
			    onSale
			    age
		    }
		}

	`)

	reqCR.Var("value", "cristiano-ronaldo-dos-santos-aveiro-2020-rare-10")
	//cogemos las variables de entorno
	ctxCR := context.Background()

	var respDataCR sCard
	if err := client.Run(ctxCR, reqCR, &respDataCR); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", respDataCR)*/

	/****************************************************************************************/

	/*mail := os.Getenv("BACKEND_USER")
		//pass := os.Getenv("BACKEND_PASSWORD")

		req := graphql.NewRequest(`
			mutation SignInMutation($input: signInInput!) {
				signIn(input: $input) {
			  		currentUser {
						slug
						coinBalance
			  		}
			  		errors {
						message
			  		}
				}
		  	}
		`)
		// set any variables
		//req.Var("value", "cristiano-ronaldo-dos-santos-aveiro-2020-rare-10")
		//valor := "{" + "email" + ":" + mail + "," + " " + "password" + ": " + hash + "}"
		//valor := {email: rufusduff6517@gmail.com, password: $2a$11$pti0jyxjdJPKaC0t720OSOA1qIJ8zDdBmkTwDrx9202Oz/xnOplNO}
		//fmt.Println(valor)

		var payload TUser
		payload.Email = mail
		payload.Password = hash

		sign, _ := json.Marshal(payload)

		//fmt.Printf("%+v\n", payload)
		fmt.Println(string(sign))
		req.Var("input", payload) //metemos email and password hashed
		// set header fields

		//fmt.Println("CSRF:", csrf)
		fmt.Println("COOKIE:", cookie)
		req.Header.Set("x-csrf-token", csrf)

		// define a Context for the request
		ctx := context.Background()

		//var graphqlResponse interface{}
		//var respData sCard
		var respData SignIn
		if err := client.Run(ctx, req, &respData); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", respData)

	}*/

	payload := setPayload()

	resp := loadSignInQuery(`mutation SignInMutation($input: signInInput!) {
    signIn(input: $input) {
      currentUser {
        slug
		email
		ethereumAddress
		availableBalance
		nickname
      }
      errors {
        message
      }
    }
  }`, payload)

	fmt.Printf("%+v\n", resp)

}
