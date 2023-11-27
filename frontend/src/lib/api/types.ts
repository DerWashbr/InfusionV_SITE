export interface OnlineStreamer {
	Username: string;
	Fraktion: string;
	Charname: string;
	Followers: number;
	ProfileImg: string;
	ThumbnailImg: string;
	IsLive: boolean;
	Viewers: number;
    ShortStory: string;
    Title: string;
	LastStreamed: Date;
}

export interface OfflineStreamer {
	Username: string;
	Fraktion: string;
	Charname: string;
	Followers: number;
	ProfileImg: string;
	ShortStory: string;
	LastStreamed: Date;
}

export interface ServerPlayer {
    clients: number;
    sv_maxclients: number;
};