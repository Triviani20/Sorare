package main

//"net/http"

//"context"
//"encoding/json"
//"fmt"
//"log"
//"os"

//"github.com/jameskeane/bcrypt"
//"github.com/machinebox/graphql"

func main() {

	//pendingOfferList := getListPendingOffer(resp.SignIn.CurrentUser)

	//fmt.Printf("%+v\n", pendingOfferList)

	//var cancelPayload TCancelPayload

	//resp2 := deletePendingOffer(CANCELOFFERPENDING_MUTATION, csrf, cookie, cancelPayload, pendingOfferList)

	//fmt.Printf("%+v\n", resp2)
	//respuesta := LogRequest()
	//fmt.Println(respuesta)

	/*jsonData := map[string]string{
		"query": `
			{
				people {
					firstname,
					lastname,
					website
				}
			}
		`,
	}
	jsonValue, _ := json.Marshal(jsonData)
	fmt.Println(jsonData)
	request, _ := http.NewRequest("POST", "https://<GRAPHQL_API_HERE>", bytes.NewBuffer(jsonValue))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))*/
}
