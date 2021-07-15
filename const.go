package main

const BASE_URL = "https://api.sorare.com"
const SIGNIN_MUTATION = `mutation SignInMutation($input: signInInput!) {
    signIn(input: $input) {
      clientMutationId
    	currentUser {
        	slug
        	email
			ethereumAddress
			availableBalance
			pendingDirectOffersSent(first: 6){
        		nodes{
           			blockchainId
          		}
          		pageInfo{
					endCursor
            		hasNextPage
            		hasPreviousPage
            		startCursor
          		}
          		totalCount
          	}
      	}	
      	errors {
        	message
      	}  
  	}
}`

const VARSIGNIN = `{
	"input":{
	  "email":"appletontim0@gmail.com",
	  "password":"$2a$11$.gLGko0c.YqseblUfoIMGe/zLus1om5bDaKpHwxbC1vj4NX3nIzlS",
	  "clientMutationId": "Trivi"
	  
	}`

const CANCELOFFERPENDING_MUTATION = `mutation CancelOfferPendingMutation ($input2: cancelOfferInput!) {
	cancelOffer(input: $input2) {
		clientMutationId
		errors {
		  message
	  	}
		
	}
}`

const VARCANCEL = `{ 
	"input2": {
		"blockchainId": "374528354332339117392269559142722165925",
		"clientMutationId": "Trivi"
	  }}
}`

const P = `{
	card(slug: "cristiano-ronaldo-dos-santos-aveiro-2020-rare-10") {
			id
			name
			position
			price
			slug
			owner {
			  ownerable {
				... on User {
				 nickname
			   }
			 }
			}
			owners {
			  ownerable {
				 ... on User {
				  nickname
				}
			  }
			  from
			  price
			}
			club {
			  name
			}
			... on WithPictureInterface {
			  pictureUrl
			}
		  }
  }
  
`
const CR = `query ($key: String!){
	card(slug: $key) {
			id
			name
			position
			price
			slug
			owner {
			  ownerable {
				... on User {
				 nickname
			   }
			 }
			}
			owners {
			  ownerable {
				 ... on User {
				  nickname
				}
			  }
			  from
			  price
			}
			club {
			  name
			}
			... on WithPictureInterface {
			  pictureUrl
			}
		  }
  }`
const VCR = `{"key": "cristiano-ronaldo-dos-santos-aveiro-2020-rare-10"}`

const SORARE = `mutation SignInMutation($input: signInInput!) {
    signIn(input: $input) {
      currentUser {
        slug
      }
      errors {
        message
      }
    }
  }`

const VS = `{
	"input": {"email": "appletontim0@gmail.com", "password": "$2a$11$.gLGko0c.YqseblUfoIMGe/zLus1om5bDaKpHwxbC1vj4NX3nIzlS"}
  }`

const SORARE2 = `mutation SignInMutation($input: signInInput!) {
	signIn(input: $input) {
	  currentUser {
		slug
		__typename
	  }
	  otpSessionChallenge
	  errors {
		code
		path
		message
		__typename
	  }
	  __typename
	}
  }`

const OP_SIGNIN_MUTATION = "SignInMutation"

const LOGIN_JWT = `mutation SignInMutation($input: signInInput!) {
    signIn(input: $input) {
      currentUser {
        	slug
        	jwtToken(aud: "TRIVI") {
          	token
          	expiredAt
        	}
        	email
			ethereumAddress
			availableBalance
			pendingDirectOffersSent(first: 20){
        		nodes{
           			blockchainId
          		}
          		pageInfo{
					endCursor
            		hasNextPage
            		hasPreviousPage
            		startCursor
          		}
          		totalCount
        	}
		}	
    }
}`

const AUD = "TRIVI"
