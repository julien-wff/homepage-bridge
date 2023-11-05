package collectors

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type PromResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Value []any `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type MinioData struct {
	ClusterUsageBytes  string `json:"cluster_usage_bytes"`
	ClusterBucketCount string `json:"cluster_bucket_count"`
	ClusterObjectCount string `json:"cluster_object_count"`
}

type MinioConfig struct {
	endpoint string
	username string
	password string
}

func MinioCollector() (MinioData, error) {
	// Get config from env
	var config MinioConfig
	config.endpoint = os.Getenv("MINIO_ENDPOINT")
	config.username = os.Getenv("MINIO_USERNAME")
	config.password = os.Getenv("MINIO_PASSWORD")

	// Check config
	if config.endpoint == "" {
		panic("MINIO_ENDPOINT is not set")
	}

	// Get metrics
	var minioData MinioData
	wg := sync.WaitGroup{}
	metrics := map[string]*string{
		"minio_cluster_usage_total_bytes":  &minioData.ClusterUsageBytes,
		"minio_cluster_bucket_total":       &minioData.ClusterBucketCount,
		"minio_cluster_usage_object_total": &minioData.ClusterObjectCount,
	}

	var err error
	for metric, result := range metrics {
		wg.Add(1)
		go getMetric(&wg, &config, metric, result, &err)
	}
	wg.Wait()

	if err != nil {
		return MinioData{}, err
	}

	return minioData, nil
}

func getMetric(wg *sync.WaitGroup, config *MinioConfig, query string, result *string, err *error) {
	defer wg.Done()

	// Build URL
	u, err1 := url.Parse(config.endpoint)
	if err1 != nil {
		*err = err1
		return
	}
	u.Path = "/api/v1/query"
	q := u.Query()
	q.Set("query", query)
	u.RawQuery = q.Encode()

	// Build request
	req, err2 := http.NewRequest("GET", u.String(), nil)
	if err2 != nil {
		*err = err2
		return
	}
	if config.username != "" && config.password != "" {
		req.SetBasicAuth(config.username, config.password)
	}

	// Send request
	res, err3 := http.DefaultClient.Do(req)
	if err3 != nil {
		*err = err3
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	// Read body into PromResponse struct
	var promResponse PromResponse
	err4 := json.NewDecoder(res.Body).Decode(&promResponse)
	if err4 != nil {
		*err = err4
		return
	}

	// Return result (check if the key exists)
	if len(promResponse.Data.Result) > 0 {
		*result = (promResponse.Data.Result[0].Value[1]).(string)
	} else {
		*result = "0"
	}
}
