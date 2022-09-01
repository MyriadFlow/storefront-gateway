package details

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/api/types"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig/dbinit"
	"github.com/TheLazarusNetwork/marketplace-engine/config/envconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/models/Org"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Details(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
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
		var org Org.Org
		testingcommon.ExtractPayload(&response, &org)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		assert.Equal(t, envconfig.EnvVars.ORG_NAME, org.Name)
		assert.Equal(t, envconfig.EnvVars.HOME_TITLE, org.HomeTitle)
		assert.Equal(t, envconfig.EnvVars.HOME_DESCRIPTION, org.HomeDescription)
		assert.Equal(t, envconfig.EnvVars.GRAPH_URL, org.GraphUrl)
		assert.Equal(t, envconfig.EnvVars.CREATIFY_CONTRACT_ADDRESS, org.CreatifyAddress)
		assert.Equal(t, envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS, org.MarketPlaceAddress)
		assert.Equal(t, envconfig.EnvVars.FOOTER, org.Footer)
		assert.ElementsMatch(t, envconfig.EnvVars.TOP_HIGHLIGHTS, org.TopHighlights)
		assert.ElementsMatch(t, envconfig.EnvVars.TRENDINGS, org.Trendings)
		assert.ElementsMatch(t, envconfig.EnvVars.TOP_BIDS, org.TopBids)
		logrus.Debug(org)
	})

	t.Run("Should be able to update org details", func(t *testing.T) {
		url := "/api/v1.0/details"
		rr := httptest.NewRecorder()
		requestBody := Org.Org{
			HomeTitle:     "Max",
			TopHighlights: []string{"43"},
			Trendings:     []string{"47"},
			TopBids:       []string{"42"},
		}
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		postDetails(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

		t.Cleanup(func() {
			err = Org.UpdateOrg(
				Org.Org{
					Name:               envconfig.EnvVars.ORG_NAME,
					HomeTitle:          envconfig.EnvVars.HOME_TITLE,
					HomeDescription:    envconfig.EnvVars.HOME_DESCRIPTION,
					GraphUrl:           envconfig.EnvVars.GRAPH_URL,
					CreatifyAddress:    envconfig.EnvVars.CREATIFY_CONTRACT_ADDRESS,
					MarketPlaceAddress: envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS,
					Footer:             envconfig.EnvVars.FOOTER,
					TopHighlights:      envconfig.EnvVars.TOP_HIGHLIGHTS,
					Trendings:          envconfig.EnvVars.TRENDINGS,
					TopBids:            envconfig.EnvVars.TOP_BIDS,
				},
			)
		})
	})

}
