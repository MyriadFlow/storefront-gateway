package products

type Product struct {
	Name			string `json:"product_name,omitempty"`
	ProductID		string `json:"product_name_id"`
	Image			string `json:"product_image,omitempty"`
	Price           string `json:"price,omitempty"`
	SellerName		string `json:"seller_name,omitempty"`
	Likes			string `json:"likes,omitempty"`
}

type MarketPlaceCounts struct {	
	Actions		int `json:"action,omitempty"`
	Artwork			int `json:"artwork,omitempty"`
	Artists			int `json:"artists,omitempty"`
}

var Highlights = []Product{
	{Name:"VR BOY", ProductID: "#070", Image:"url", Price:"$100,000", SellerName:" John Sander", Likes:"78"},
	{Name:"VR BOY", ProductID: "#070", Image:"url", Price:"$100,000", SellerName:" John Sander", Likes:"78"},
	{Name:"VR BOY", ProductID: "#070", Image:"url", Price:"$100,000", SellerName:" John Sander", Likes:"78"},
	{Name:"VR BOY", ProductID: "#070", Image:"url", Price:"$100,000", SellerName:" John Sander", Likes:"78"},
}

var Trending = []Product{
    {Name:"Crypto Raptors", ProductID: "#123", Image:"url", Price:"$23984", SellerName:"Sonu Nigam", Likes:"100"},
    {Name:"The Binary Girl", ProductID: "#124", Image:"url", Price:"$3765", SellerName:" Manashi parajpe", Likes:"728"},
    {Name:"The Binary Girl", ProductID: "#456", Image:"url", Price:"$643528", SellerName:" Prajakta Mali", Likes:"738"},
    {Name:"Sphere Art", ProductID: "#654", Image:"url", Price:"$64252", SellerName:"Monika Binsal", Likes:"378"},
}