package models

import "time"

type NFTResponse struct {
	ETHTest     NFTSingleResponse `json:"eth_test"`
	PolygonTest NFTSingleResponse `json:"polygon_test"`
	BscTest     NFTSingleResponse `json:"bsc_test"`
}

type NFTSingleResponse struct {
	OwnedNfts []struct {
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`
		Id struct {
			TokenId       string `json:"tokenId"`
			TokenMetadata struct {
				TokenType string `json:"tokenType"`
			} `json:"tokenMetadata"`
		} `json:"id"`
		Balance     string `json:"balance"`
		Title       string `json:"title"`
		Description string `json:"description"`
		TokenUri    struct {
			Raw     string `json:"raw"`
			Gateway string `json:"gateway"`
		} `json:"tokenUri"`
		Media []struct {
			Raw       string `json:"raw"`
			Gateway   string `json:"gateway"`
			Thumbnail string `json:"thumbnail,omitempty"`
			Format    string `json:"format,omitempty"`
			Bytes     int    `json:"bytes,omitempty"`
		} `json:"media"`
		Metadata struct {
			BackgroundImage string `json:"background_image,omitempty"`
			Image           string `json:"image"`
			IsNormalized    bool   `json:"is_normalized,omitempty"`
			SegmentLength   int    `json:"segment_length,omitempty"`
			ImageUrl        string `json:"image_url,omitempty"`
			Name            string `json:"name"`
			Description     string `json:"description"`
			Attributes      []struct {
				DisplayType string      `json:"display_type"`
				Value       interface{} `json:"value"`
				TraitType   string      `json:"trait_type"`
			} `json:"attributes,omitempty"`
			NameLength int    `json:"name_length,omitempty"`
			Version    int    `json:"version,omitempty"`
			Url        string `json:"url,omitempty"`
		} `json:"metadata"`
		TimeLastUpdated  time.Time `json:"timeLastUpdated"`
		ContractMetadata struct {
			TokenType           string `json:"tokenType"`
			ContractDeployer    string `json:"contractDeployer,omitempty"`
			DeployedBlockNumber int    `json:"deployedBlockNumber,omitempty"`
			OpenSea             struct {
				FloorPrice            float64   `json:"floorPrice"`
				CollectionName        string    `json:"collectionName"`
				SafelistRequestStatus string    `json:"safelistRequestStatus"`
				ImageUrl              string    `json:"imageUrl"`
				Description           string    `json:"description"`
				ExternalUrl           string    `json:"externalUrl"`
				TwitterUsername       string    `json:"twitterUsername"`
				LastIngestedAt        time.Time `json:"lastIngestedAt"`
				DiscordUrl            string    `json:"discordUrl,omitempty"`
			} `json:"openSea"`
			Name        string `json:"name,omitempty"`
			Symbol      string `json:"symbol,omitempty"`
			TotalSupply string `json:"totalSupply,omitempty"`
		} `json:"contractMetadata"`
	} `json:"ownedNfts"`
	TotalCount int    `json:"totalCount"`
	BlockHash  string `json:"blockHash"`
}
