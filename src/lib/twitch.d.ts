declare namespace Twitch {
	interface PlayerOptions {
		channel: string;
		width: string | number;
		height: string | number;
		autoplay?: boolean;
		layout?: string;
		muted?: boolean;
		parent: string | string[];
		playsinline?: boolean;
	}

	class Player {
		constructor(id: string, options: PlayerOptions);
	}
}
