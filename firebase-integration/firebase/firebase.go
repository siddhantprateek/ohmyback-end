package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	// firebaseAuth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func firebaseApp() {
	// context: package provides a way to carry around request-scoped values
	// and cancelation signals across API boundaries.
	// Used by main function, initialization and tests and as the top-level
	// Context for incoming request.
	ctx := context.Background() // returns non-nil, empty Context.

	// path to service account configuration
	opt := option.WithCredentialsFile("firebase-integration/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// fmt.Printf(firebaseAuth)

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
