package auth

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
)

func supabaseinit() {
	supabaseUrl := os.Getenv("SUPABASE_API_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	supabaseClient = supabase.CreateClient(supabaseUrl, supabaseKey)
}

var supabaseClient *supabase.Client

// marshal and unmarshal
func FromJWTSupabaseTokenGetData(c *gin.Context, user_token string) (string, string,error) {
	ctx := context.Background()
	supabaseUser, err := supabaseClient.Auth.User(ctx, user_token)
	if err != nil {
		log.Println("Unable to Authenticate Token", err.Error())
		c.JSON(http.StatusInternalServerError, "unable to authenticate ")
		return "", "",err
	}
	return supabaseUser.ID, supabaseUser.Email,nil
}

func SupabaseAthentication(email, password string) (bool, error) {
	// supabase := supa.CreateClient(os.Getenv("SupabaseUrl"), os.Getenv("SupabaseKey"))

	_, errSigin := supabaseClient.Auth.SignIn(context.Background(), supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if errSigin != nil {
		log.Println(errSigin.Error())
		_, errSigup := supabaseClient.Auth.SignUp(context.Background(), supabase.UserCredentials{
			Email:    email,
			Password: "password",
		})
		if errSigup != nil {
			log.Println(errSigup.Error())
			return false, errSigup
		}

	}
	if err := supabaseClient.Auth.SendMagicLink(context.Background(), email); err != nil {
		log.Println("sorry! we could not send the magic link")
	}
	return true, nil
}

func SupabaseSendMagicLink(email string) bool {
	// supabase := supa.CreateClient(os.Getenv("SupabaseUrl"), os.Getenv("SupabaseKey"))

	if err := supabaseClient.Auth.SendMagicLink(context.Background(), email); err != nil {
		log.Println("sorry! we could not send the magic link")
		return false
	}
	log.Println("Magic link send successfully to your Email id")
	return true
}
