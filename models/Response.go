package models

type ResponseData struct {
	BriefInfo      string              `json:"brief_info"`
	Currency       string              `json:"currency"`
	CurrencySymbol string              `json:"currency_symbol"`
	Depositable    bool                `json:"depositable"`
	GoodsInfo      map[string]GoodInfo `json:"goods_info"`
	InventoryPrice string              `json:"inventory_price"`
	Items          []Item              `json:"items"`
	PageNum        int                 `json:"page_num"`
	PageSize       int                 `json:"page_size"`
	TotalAmount    string              `json:"total_amount"`
	TotalAmountUsd string              `json:"total_amount_usd"`
	TotalCount     int                 `json:"total_count"`
	TotalPage      int                 `json:"total_page"`
}

type GoodInfo struct {
	AppID           int         `json:"appid"`
	Can3DInspect    bool        `json:"can_3d_inspect"`
	CanInspect      bool        `json:"can_inspect"`
	Description     interface{} `json:"description"`
	Game            string      `json:"game"`
	GoodsID         int         `json:"goods_id"`
	IconURL         string      `json:"icon_url"`
	ItemID          interface{} `json:"item_id"`
	MarketHashName  string      `json:"market_hash_name"`
	MarketMinPrice  string      `json:"market_min_price"`
	Name            string      `json:"name"`
	OriginalIconURL string      `json:"original_icon_url"`
	ShortName       string      `json:"short_name"`
	SteamPrice      string      `json:"steam_price"`
	SteamPriceCNY   string      `json:"steam_price_cny"`
}

type Item struct {
	ActionLink           string      `json:"action_link"`
	AllowBundleInventory bool        `json:"allow_bundle_inventory"`
	Amount               int         `json:"amount"`
	Appid                int         `json:"appid"`
	AssetInfo            AssetInfo   `json:"asset_info"`
	AssetId              string      `json:"assetid"`
	BackgroundImageURL   string      `json:"background_image_url"`
	CanUseInspectTrnURL  bool        `json:"can_use_inspect_trn_url"`
	ClassId              string      `json:"classid"`
	ContextId            int         `json:"contextid"`
	DepositIndex         interface{} `json:"deposit_index"`
	Equipped             bool        `json:"equipped"`
	FraudWarnings        interface{} `json:"fraudwarnings"`
	Game                 string      `json:"game"`
	GoodsID              int         `json:"goods_id"`
	IconURL              string      `json:"icon_url"`
	ImgSrc               string      `json:"img_src"`
	InstanceId           string      `json:"instanceid"`
	ItemID               interface{} `json:"item_id"`
	MarketHashName       string      `json:"market_hash_name"`
	MarketMinPrice       string      `json:"market_min_price"`
	Name                 string      `json:"name"`
	OriginalIconURL      string      `json:"original_icon_url"`
	Progress             interface{} `json:"progress"`
	ProgressText         interface{} `json:"progress_text"`
	Properties           interface{} `json:"properties"`
	PunishEndTime        interface{} `json:"punish_end_time"`
	SellMinPrice         string      `json:"sell_min_price"`
	SellOrderID          interface{} `json:"sell_order_id"`
	SellOrderIncome      string      `json:"sell_order_income"`
	SellOrderMode        interface{} `json:"sell_order_mode"`
	SellOrderPrice       string      `json:"sell_order_price"`
	ShortName            string      `json:"short_name"`
	State                int         `json:"state"`
	StateText            string      `json:"state_text"`
	StateToast           string      `json:"state_toast"`
	SteamPrice           string      `json:"steam_price"`
	Tradable             bool        `json:"tradable"`
	TradableText         string      `json:"tradable_text"`
	TradableTime         int         `json:"tradable_time"`
}

type AssetInfo struct {
	ActionLink string           `json:"action_link"`
	Appid      int              `json:"appid"`
	AssetId    string           `json:"assetid"`
	ClassId    string           `json:"classid"`
	ContextId  int              `json:"contextid"`
	GoodsID    int              `json:"goods_id"`
	ID         string           `json:"id"`
	Info       AssetInfoDetails `json:"info"`
	InstanceId string           `json:"instanceid"`
	PaintWear  string           `json:"paintwear"`
}

type AssetInfoDetails struct {
	FraudWarnings   interface{}   `json:"fraudwarnings"`
	IconURL         string        `json:"icon_url"`
	OriginalIconURL string        `json:"original_icon_url"`
	PaintIndex      int           `json:"paintindex"`
	PaintSeed       int           `json:"paintseed"`
	Stickers        []interface{} `json:"stickers"`
	TournamentTags  []interface{} `json:"tournament_tags"`
}

type Response struct {
	Code string       `json:"code"`
	Data ResponseData `json:"data"`
	Msg  interface{}  `json:"msg"`
}
