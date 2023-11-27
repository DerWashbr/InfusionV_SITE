import axios from 'axios';
import type { OnlineStreamer, OfflineStreamer } from './types';

export let onlineStreamers: OnlineStreamer[] = [];
export let offlineStreamers: OfflineStreamer[] = [];
export let playerCount = 0;
export let maxClients = 0;
export let totalViewers = 0;
export let totalStreamers = 0;
export let isApiWorking = true;
export const isDebug = false;

export const fetchData = async () => {
	try {
		const { data: streamers } = await axios.get<OnlineStreamer[]>(
			'http://localhost:7080/api/streamers/all'
		);

		onlineStreamers = [];
		offlineStreamers = [];

		streamers.forEach((s) => {
			if (s.IsLive) {
				onlineStreamers.push({
					...s,
					ThumbnailImg: s.ThumbnailImg.replace('{width}', '200').replace('{height}', '200')
				});
			} else {
				offlineStreamers.push({ ...s });
			}
		});

		totalViewers = onlineStreamers.reduce((sum, s) => sum + s.Viewers, 0);
		totalStreamers = onlineStreamers.length;
	} catch (error) {
		console.error(error);
		isApiWorking = false;
		playerCount = 0; // Set playerCount to 0
		maxClients = 0; // Set maxClients to 0
	}

	if (isDebug) {
		console.log('Debug mode');
	}
};

export function formatElapsedTime(startTime: Date | undefined): string {
	if (!startTime) {
		return 'Keine Daten vorhanden';
	}

	const startDate = new Date(startTime);
	const currentDate = new Date();
	const elapsedTime = (currentDate.getTime() - startDate.getTime()) / (60 * 1000); // Difference in minutes
	const elapsedDays = Math.floor(elapsedTime / 1440);
	const elapsedHours = Math.floor((elapsedTime % 1440) / 60);
	const elapsedMinutes = Math.floor(elapsedTime % 60);

	if (elapsedDays > 0) {
		return `${elapsedDays}d ${elapsedHours}h ${elapsedMinutes}m`;
	} else if (elapsedHours > 0) {
		return `${elapsedHours}h ${elapsedMinutes}m`;
	} else {
		return `${elapsedMinutes}m`;
	}
}
