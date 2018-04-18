package tweetsearch

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)
var 
	(
	 consumerKey string   // = "yVUaVdBYnsAzIqkz3Yt7skHIY"
	 consumerSecret string  = "ELGvyaVaxiamrMLgtilB0tqCcw1oyxAvBX6a75F7lWXP8X7G0j"
	  accessToken string     = "744602175540592640-nKlNJ2CBh5J0knOZTd90OzKDDeV1n69"
	 accessSecret string    = "6yctXw0WBUiaDKlu1hspuYybkldLBqCoyoEkJMlOmdoq2"
	)
var log = logger.GetLogger("activity-tweetsearch")
// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	

fmt.Println(" Enter your consumerKey")
	
	//reader := bufio.NewReader(os.Stdin)
	//consumerKey, _ = reader.ReadString('\n')
	fmt.Scanf("%s", &consumerKey)
	consumerKey := context.GetInput("consumerKey").(string)

   
	config := oauth1.NewConfig(consumerKey,consumerSecret)
	token := oauth1.NewToken(accessToken,accessSecret)

	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Name:%+v\n", user.ScreenName)

	//Sending a Tweet
	_,_,_ = client.Statuses.Update("Learning twitter API using GOlang #GOlang", nil)
	 

	// Getting 20 tweets from wall
	timeline,_,_ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
    Count: 20,
	})
	fmt.Println(timeline)
	
	//status

	tweet, _, _ := client.Statuses.Show(585613041028431872, nil)
	fmt.Println("******* STATUS IS ******* \n", tweet)



	return true, nil
}
