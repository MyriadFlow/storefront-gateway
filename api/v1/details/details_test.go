package details

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/api/types"
	"github.com/TheLazarusNetwork/marketplace-engine/config"
	"github.com/TheLazarusNetwork/marketplace-engine/models/Org"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_GetDetails(t *testing.T) {
	config.Init("../../../.env")
	logwrapper.Init("../../../logs")
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
		var org Org.Org
		testingcommon.ExtractPayload(&response, &org)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		assert.Equal(t, envutil.MustGetEnv("ORG_NAME"), org.Name)
		assert.Equal(t, envutil.MustGetEnv("HOME_TITLE"), org.HomeTitle)
		assert.Equal(t, envutil.MustGetEnv("HOME_DESCRIPTION"), org.HomeDescription)
		assert.Equal(t, envutil.MustGetEnv("GRAPH_URL"), org.GraphUrl)
		assert.Equal(t, envutil.MustGetEnv("CREATIFY_ADDRESS"), org.CreatifyAddress)
		assert.Equal(t, envutil.MustGetEnv("MARKETPLACE_ADDRESS"), org.MarketPlaceAddress)
		assert.Equal(t, envutil.MustGetEnv("FOOTER"), org.Footer)

		//TODO
		// assert.Equal(t, envutil.MustGetEnv("TOP_HIGHLIGHTS"), )
		// pq.Array(org.TopHighlights)
		// assert.Equal(t, envutil.MustGetEnv("TRENDINGS"), org.Trendings)
		// assert.Equal(t, envutil.MustGetEnv("TOP_BIDS"), strings.Join(org.TopBids))

		logrus.Debug(org)
	})

}
