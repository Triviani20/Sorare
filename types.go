package main

type User struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type Salt struct {
	Salt string `json:"salt"`
}

type TCard struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Onsale bool   `json:"onSale"`
	Age    int    `json:"age"`
}

type SCard struct {
	Card TCard `json:"card"`
}

type TUser struct {
	Email               string `json:"email"`
	Password            string `json:"password"`
	ClientMutationId    string `json:"clientMutationId"`
	OtpSessionChallenge string `json:"otpSessionChallenge"`
}
type SignInInput struct {
	Input TUser `json:"input"`
}

type OfferConnection struct {
	Nodes    []TBlockchainId `json:"nodes"`
	PageInfo TPageInfo       `json:"pageInfo"`
}
type TBlockchainId struct {
	BlockchainId string `json:"blockchainId"`
}
type TPageInfo struct {
	EndCursor       string `json:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
}
type TError struct {
	Message string `json:"message"`
}

type SignInPayload struct {
	ClientMutationId string `json:"clientMutationId"`
	CUser            TCUser `json:"currentUser"`
}

///
/// TYPES FOR SIGNIN RESPONSES
///
type TCUser struct {
	Slug                    string          `json:"slug"`
	Email                   string          `json:"email"`
	EthereumAddress         string          `json:"ethereumAddress"`
	AvailableBalance        string          `json:"availableBalance"`
	PendingDirectOffersSent OfferConnection `json:"pendingDirectOffersSent"`
	JwtToken                JwtTK           `json:"jwtToken"`
}
type JwtTK struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expiredAt"`
}
type SignInResp struct {
	ClientMutationId    string   `json:"clientMutationId"`
	CurrentUser         TCUser   `json:"currentUser"`
	Err                 []TError `json:"errors"`
	OtpSessionChallenge string   `json:"otpSessionChallenge"`
}

type SignIn struct {
	SignIn SignInResp `json:"signIn"`
}

///
///
///
type TCancelPayload struct {
	BlockchainId     string `json:"blockchainId"`
	ClientMutationId string `json:"clientMutationId"`
}

type Cancel struct {
	Cancel CancelOfferResp `json:"cancelOffer"`
}

type CancelOfferResp struct {
	Errors           []TError `json:"errors"`
	ClientMutationId string   `json:"clientMutationId"`
}
type SessionCookie struct {
	SorareSessionId string `json:"_sorare_session_id"`
}

//coger variables os.getenv()

type GraphQLRequest struct {
	Query         string `json:"query"`
	OperationName string `json:"OperationName,omitempty"`
	Variables     string `json:"variables"`
}
