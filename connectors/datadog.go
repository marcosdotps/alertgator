package connectors

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func DatadogEvents(tags []string) {
	// create a context with 3 seconds timeout
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	//initialize start/end time range for the query (now and one week ago)
	today := time.Now()
	weekAgo := today.Add(-7 * 24 * time.Hour)

	// initialize datadog api client
	ctx := datadog.NewDefaultContext(timeoutCtx)
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewEventsApi(apiClient)

	//compose tags query based on inputs
	var tagsQuery string
	for _, tag := range tags {
		tagsQuery += fmt.Sprintf("%s,", tag)
	}
	//remove trailing comma
	tagsQuery = strings.TrimRight(tagsQuery, ",")
	listParams := *datadogV1.NewListEventsOptionalParameters().WithTags(tagsQuery)
	resp, r, err := api.ListEvents(ctx, weekAgo.Unix(), today.Unix(), listParams)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`:\n%s\n", responseContent)
}
