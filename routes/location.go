package routes

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/kataras/iris/v12"
)

func handleAPIError(ctx iris.Context, err error) {
	ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
		Title("🛖🕯️-> Internal Server Error").
		Detail(err.Error()))
}

func makeAPIRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Autocomplete(ctx iris.Context) {
	limit := "10"
	location := ctx.URLParam("location")
	limitQuery := ctx.URLParam("limit")
	if limitQuery != "" {
		limit = limitQuery
	}

	apiKey := os.Getenv("LOCATION_TOKEN")
	if apiKey == "" {
		handleAPIError(ctx, errors.New("LOCATION_TOKEN environment variable is not set"))
		return
	}

	url := "https://api.locationiq.com/v1/autocomplete.php?key=" + apiKey + "&q=" + location + "&limit=" + limit

	body, err := makeAPIRequest(url)
	if err != nil {
		handleAPIError(ctx, err)
		return
	}

	var objMap []map[string]interface{}
	if err := json.Unmarshal(body, &objMap); err != nil {
		handleAPIError(ctx, err)
		return
	}

	ctx.JSON(objMap)
}

func Search(ctx iris.Context) {
	location := ctx.URLParam("location")

	apiKey := os.Getenv("LOCATION_TOKEN")
	if apiKey == "" {
		handleAPIError(ctx, errors.New("LOCATION_TOKEN environment variable is not set"))
		return
	}

	url := "https://api.locationiq.com/v1/search.php?key=" + apiKey + "&q=" + location + "&format=json&dedupe=1&addressdetails=1&matchquality=1&normalizeaddress=1&normalizecity=1"

	body, err := makeAPIRequest(url)
	if err != nil {
		handleAPIError(ctx, err)
		return
	}

	var objMap []map[string]interface{}
	if err := json.Unmarshal(body, &objMap); err != nil {
		handleAPIError(ctx, err)
		return
	}

	ctx.JSON(objMap)
}
