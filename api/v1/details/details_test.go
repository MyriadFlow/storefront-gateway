package details

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/api/types"
	"github.com/TheLazarusNetwork/marketplace-engine/config"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig/dbinit"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type apiResponse struct {
	Name               string   `json:"name"`
	HomeTitle          string   `json:"homeTitle"`
	HomeDescription    string   `json:"homeDescription"`
	GraphUrl           string   `json:"graphurl"`
	CreatifyAddress    string   `json:"creatifyAddress"`
	MarketPlaceAddress string   `json:"marketPlaceAddress"`
	Footer             string   `json:"footer"`
	TopHighlights      []string `json:"topHighlights"`
	Trendings          []string `json:"trendings"`
	TopBids            []string `json:"topBids"`
}

func Test_GetDetails(t *testing.T) {
	config.Init("../../../.env")
	logwrapper.Init("../../../logs")
	dbinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	t.Run("Should be able to get org details", func(t *testing.T) {
		_ = "/api/v1.0/details"
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		getDetails(c)
		var response types.ApiResponse
		body := rr.Body
		json.NewDecoder(body).Decode(&response)
		var org apiResponse
		testingcommon.ExtractPayload(&response, &org)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		assert.Equal(t, envutil.MustGetEnv("ORG_NAME"), org.Name)
		assert.Equal(t, envutil.MustGetEnv("HOME_TITLE"), org.HomeTitle)
		assert.Equal(t, envutil.MustGetEnv("HOME_DESCRIPTION"), org.HomeDescription)
		assert.Equal(t, envutil.MustGetEnv("GRAPH_URL"), org.GraphUrl)
		assert.Equal(t, envutil.MustGetEnv("CREATIFY_ADDRESS"), org.CreatifyAddress)
		assert.Equal(t, envutil.MustGetEnv("MARKETPLACE_ADDRESS"), org.MarketPlaceAddress)
		assert.Equal(t, envutil.MustGetEnv("FOOTER"), org.Footer)

		assert.Equal(t,
			envutil.MustGetEnv("TOP_HIGHLIGHTS"), strings.Join(org.TopHighlights, ","),
		)
		assert.Equal(t, envutil.MustGetEnv("TRENDINGS"), strings.Join(org.Trendings, ","))
		assert.Equal(t, envutil.MustGetEnv("TOP_BIDS"), strings.Join(org.TopBids, ","))

		logrus.Debug(org)
	})

}